package goaws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Client get the EC2 Client for this region
func ec2svc(region string) (ec *ec2.EC2) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to get ec2 client for region %s ", region)
	}
	cfg.Region = region
	return ec2.New(cfg)
}
