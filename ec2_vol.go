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

// Volume is attached to a CPU
type Volume struct {
	/*
		VolumeId    string
		InstanceId  string
		SnapshotId  string
		AvailZone   string
		Region      string
		Size        int64
		State       ec2.VolumeState
		AttachState ec2.VolumeAttachmentState
	*/
	region string
	ec2.CreateVolumeOutput
}

// String will print out a representation of the volume
func (v Volume) String() string {
	return fmt.Sprintf("%s %s %s %s %dGb", v.VolumeId, v.InstanceId, v.AvailZone, v.State, v.Size)
}

//func (v *Volume) VolumeId() string {
//	return v.VolumeId
//}

//func (v *Volume) SnapshotId() string {
//	return v.SnapshotId
//}

//func (v *Volume) Size() int {
//	return v.Size
//}

func (v *Volume) AvailZone() string {
	return *v.AvailabilityZone
}

func (v *Volume) Region() string {
	return v.region
}

//func (v *Volume) State() string {
//	return v.State
//}

func (v *Volume) AttachmentState() string {
	log.Fatalf("  %+v ", v.Attachments)
	return "foo" //v.Attachments[0].State
}

func (v *Volume) InstanceId() string {
	log.Fatalf("  %+v ", v.Attachments)
	return "bar" // v.Attachments[0].InstanceId
}

// Volumes returns the "Fetched" Volumes
func Volumes(region string) map[string]*Volume {
	var reg *Region
	if reg := RegionMap.Get(region); reg == nil {
		return nil
	}
	reg.Volumes = FetchVolumes(reg.Name)
	return reg.Volumes
}

// GetVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func FetchVolumes(region string) (vmap map[string]*Volume) {
	var e *ec2.EC2
	if e = ec2svc(region); e == nil {
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
		vol := &Volume{}
		vol.CreateVolumeOutput = result.Volumes[0]
		vmap[*awsvol.VolumeId] = vol
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
	e := ec2svc(region)
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
	if svc = ec2svc(region); svc == nil {
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
