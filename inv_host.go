package main

import "github.com/aws/aws-sdk-go-v2/service/ec2"

// Host represents an Instance on AWS
type Host struct {
	Region string
	*ec2.Instance
}

// HostFromInstance creates our Host from an AWS Instance
func HostFromInstance(ec2Inst *ec2.Instance) Host {
	return Host{Region: "unknown", Instance: ec2Inst}
}
