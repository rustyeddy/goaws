package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
)

type Volume struct {
	VolumeId    string
	InstanceId  string
	SnapshotId  string
	AvailZone   string
	Region      string
	Size        int64
	State       ec2.VolumeState
	AttachState ec2.VolumeAttachmentState
	raw         *ec2.CreateVolumeOutput
}

func (v *Volume) String() string {
	return fmt.Sprintf("%s %s %s %s %dGb", v.VolumeId, v.InstanceId, v.AvailZone, v.State, v.Size)
}

// GetVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func GetVolumes(region string) (vmap Volmap) {
	log.Println("GetVolumes from ", region)
	defer log.Printf("  return GetVolumes region %s ", region)

	idxname := region + "-volumes"
	log.Debugf("  looking in local cache for %s ", idxname)

	err := cache.FetchObject(idxname, &vmap)
	if err == nil && vmap != nil {
		log.Printf("  Found cached version of %s .. ", idxname)
		return vmap
	}

	log.Debugf("  Fetch Volumes from AWS for %s", region)
	var e *ec2.EC2
	if e = getEC2(region); e == nil {
		log.Errorf("  failed to get aws client for %s ", region)
		return nil
	}

	// Get the Volumes from AWS
	log.Debugf("  sending request for volumes region ", region)
	req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  # failed response to request %v", err)
		return nil
	}

	log.Debugf("  got result %v from region %s ", result, region)

	// Unparse the AWS format for ours
	if vmap = vdisksFromAWS(result, region); vmap == nil {
		log.Error("failed to get vdisks from aws ")
		return nil
	}

	// Store the object in the cache
	go func() {
		// Save the results in the storage cache
		obj, err := cache.StoreObject(idxname, vmap)
		log.Debugf("  cache store object idx %s -> %v ", idxname, obj)
		if err != nil || obj == nil {
			log.Errorf("  failed to cache object idx %s -> %v ", idxname, obj)
			return
		}
	}()

	// VDisk maps
	return vmap
}

// parse the response from AWS
func vdisksFromAWS(result *ec2.DescribeVolumesOutput, region string) (vmap Volmap) {
	for _, vol := range result.Volumes {
		for _, att := range vol.Attachments {
			vol := &Volume{
				raw:         &vol,
				VolumeId:    *vol.VolumeId,
				SnapshotId:  *vol.SnapshotId,
				Size:        *vol.Size,
				State:       vol.State,
				InstanceId:  *att.InstanceId,
				AttachState: att.State,
				AvailZone:   *vol.AvailabilityZone,
				Region:      region,
			}
			vmap[vol.VolumeId] = vol
			allVolumes[vol.VolumeId] = vol
		}
	}
	return vmap
}

// GetVolume will update the instance information for vol
func GetVolume(volid string) *Volume {

	log.Printf("GetVolume %s ", volid)
	defer log.Printf("  return GetVolume %s ", volid)

	if vol, ex := allVolumes[volid]; ex {
		return vol
	}
	return nil
}

// DeleteVolume will send a request to AWS to delete the given volume
func DeleteVolume(volid string) error {

	log.Debugln("DeleteVolume %s", volid)
	defer log.Debugln("  returning from deleteVolume ")

	vol, ex := allVolumes[volid]
	if !ex {
		return store.ErrNotFound.Append(string(volid))
	}
	region := vol.Region

	var svc *ec2.EC2
	if svc = getEC2(region); svc == nil {
		return fmt.Errorf("delvol reg %s vol %s: ", region, vol.VolumeId)
	}

	log.Debugf("  volume state %s", vol.State)
	switch vol.State {
	case "creating", "in-use":
		if err := DetachVolume(region, vol); err != nil {
			return fmt.Errorf("detach failed %s / %s :", region, vol.VolumeId)
		}
	case "deleting", "deleted", "error":
		log.Debugln("  skipping  volume status ", vol.State)
		return nil

	case "available":
		// OK, prepare the request to detach volume
		req := svc.DetachVolumeRequest(&ec2.DetachVolumeInput{
			VolumeId: aws.String(vol.VolumeId),
		})
		// Send the request and get a result
		if result, err := req.Send(); err != nil {
			return fmt.Errorf("volid %s: %v", vol.State, err)
		} else {
			log.Debugln("  sent request to detach volume %v ", result)
		}
	default:
		log.Errorf("  whoa do not know about state, continue ", vol.State)
	}

	log.Debugf("  sending delete for region %s vol %s\n ", region, vol.VolumeId)
	req := svc.DeleteVolumeRequest(&ec2.DeleteVolumeInput{
		VolumeId: aws.String(vol.VolumeId),
	})

	// Send the request, get the results and dump
	result, err := req.Send()
	if err != nil {
		return fmt.Errorf("  # failed response to request %v \n", err)
	}
	log.Fatalf("  result %+v \n", result)
	return nil
}

// DetachVolume will remove the volume from the instance
func DetachVolume(region string, vol *Volume) error {

	log.Debugln("DetatchVolume %s %s ", region, vol.VolumeId)
	defer log.Debugln("  returning from deleteVolume ")

	var svc *ec2.EC2
	if svc = getEC2(region); svc == nil {
		return fmt.Errorf("detach vol %s %s: ", region, vol.VolumeId)
	}

	log.Debugf("  sending delete for region %s vol %s\n ", region, vol.VolumeId)
	req := svc.DetachVolumeRequest(&ec2.DetachVolumeInput{
		VolumeId: aws.String(vol.VolumeId),
	})
	result, err := req.Send()
	if err != nil {
		return fmt.Errorf("detach volume %v: ", err)
	}
	log.Infoln("  detached volume %s result %+v \n", vol.VolumeId, result)
	return nil
}
