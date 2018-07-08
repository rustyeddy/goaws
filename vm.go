package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type VM struct {
	InstanceId string
	VolumeId   string
	State      ec2.InstanceState
	KeyName    string
	AvaillZone string
}

func (v *VM) String() string {
	return fmt.Sprintf("%s %s %s %s", v.InstanceId, v.VolumeId, v.State, v.KeyName)
}
