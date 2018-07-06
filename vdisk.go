package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type VDisk struct {
	VolumeId    string
	InstanceId  string
	SnapshotId  string
	AvailZone   string
	Size        int64
	State       ec2.VolumeState
	AttachState ec2.VolumeAttachmentState
	raw         *ec2.CreateVolumeOutput
}

func (v *VDisk) String() string {
	return fmt.Sprintf("%s %s %s %dGb", v.VolumeId, v.InstanceId, v.AvailZone, v.Size)
}
