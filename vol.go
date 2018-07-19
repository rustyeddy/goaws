package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
)

var (
	volumes map[string]*Volume
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

// Volumes returns the Volumemap
func Volumes(region string) map[string]*Volume {
	return FetchVolumes(region)
}

// GetVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func FetchVolumes(region string) (vmap map[string]*Volume) {
	var e *ec2.EC2
	if e = Client(region); e == nil {
		log.Errorf("  failed to get aws client for %s ", region)
		return nil
	}

	req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	result, err := req.Send()
	if err != nil {
		log.Warning("  # failed response to request %v", err)
		return nil
	}
	if len(result.Volumes) <= 0 {
		log.Info("  no volumes found ")
		return nil
	}
	if vmap = vmapFromAWS(result, region); vmap == nil {
		log.Errorln("  # failed to extract volumes from result")
		return nil
	}
	return vmap
}

// parse the response from AWS
func vmapFromAWS(result *ec2.DescribeVolumesOutput, region string) (vmap map[string]*Volume) {
	vmap = make(map[string]*Volume)
	for _, awsvol := range result.Volumes {
		vol := &Volume{
			raw:        &awsvol,
			VolumeId:   *awsvol.VolumeId,
			SnapshotId: *awsvol.SnapshotId,
			Size:       *awsvol.Size,
			State:      awsvol.State,
			AvailZone:  *awsvol.AvailabilityZone,
			Region:     region,
		}

		for _, att := range awsvol.Attachments {
			vol.AttachState = att.State
			vol.InstanceId = *att.InstanceId
		}
		vmap[vol.VolumeId] = vol
	}
	return vmap
}

// DeleteVolume will send a request to AWS to delete the given volume
func DeleteVolume(region string, volid string) error {
	var (
		vol *Volume
		ex  bool
	)

	volumes := Volumes(region)
	if vol, ex = volumes[volid]; !ex {
		return store.ErrNotFound.Append(string(volid))
	}

	if vol.State != "available" {
		fmt.Printf("  volume state expected (available) got (%s) ", vol.State)
	}
	/*
		switch vol.State {
		case "creating", "in-use":
			if err := DetachVolume(region, vol.VolumeId); err != nil {
				return fmt.Errorf("detach failed %s / %s :", region, vol.VolumeId)
			}
		case "deleting", "deleted", "error":
			log.Debugln("  skipping  volume status ", vol.State)
			return nil

		case "available":
			// OK, prepare the request to detach volume
			e := Client(region)
			fmt.Printf("  DetachVolume request ")
			req := e.DetachVolumeRequest(&ec2.DetachVolumeInput{
				VolumeId: aws.String(vol.VolumeId),
			})
			// Send the request and get a result
			result, err := req.Send()
			if err != nil {
				return fmt.Errorf("volid %s: %v", vol.State, err)
			}
			fmt.Printf("  result from Detach %v\n", result)

		default:
			log.Errorf("  whoa do not know about state, continue ", vol.State)
		}
	*/
	e := Client(region)
	req := e.DeleteVolumeRequest(&ec2.DeleteVolumeInput{
		VolumeId: aws.String(volid),
	})

	// Send the request, get the results and dump
	result, err := req.Send()
	if err != nil {
		return fmt.Errorf("  # failed response to request %v \n", err)
	}
	log.Infof("  DELETE result %+v \n", result)
	return nil
}

// DetachVolume will remove the volume from the instance
func DetachVolume(region string, volid string) error {

	log.Debugln("DetatchVolume %s %s ", region, volid)
	defer log.Debugln("  returning from deleteVolume ")

	var svc *ec2.EC2
	if svc = Client(region); svc == nil {
		return fmt.Errorf("detach vol %s %s: ", region, volid)
	}

	log.Debugf("  sending delete for region %s vol %s\n ", region, volid)
	req := svc.DetachVolumeRequest(&ec2.DetachVolumeInput{
		VolumeId: aws.String(volid),
	})
	result, err := req.Send()
	if err != nil {
		return fmt.Errorf("detach volume %v: ", err)
	}
	log.Infoln("  detached volume %s result %+v \n", volid, result)
	return nil
}
