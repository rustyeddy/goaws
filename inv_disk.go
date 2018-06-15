package aws

import "github.com/aws/aws-sdk-go-v2/service/ec2"

// Disk represents an AWS EBS Volume
type Disk struct {
	Region string
	*ec2.CreateVolumeOutput
}

// DiskFromVolume returns a *Disk from ec2 volume input
func DiskFromVolume(vol *ec2.CreateVolumeOutput) Disk {
	return Disk{CreateVolumeOutput: vol}
}
