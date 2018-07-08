package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// getEC2 returns an ec2 service ready for use
func getEC2(region string) (ec2Svc *ec2.EC2) {
	log.Debugln("Get EC2 for region ", region)
	defer log.Debugln(" leaving EC2 %v ", ec2Svc)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("  Failed to Load Default AWS Config %q -> %v ", region, err)
		return nil
	}

	fmt.Println("  loaded Default config, region ", region)

	cfg.Region = region
	ec2Svc = ec2.New(cfg)
	if ec2Svc == nil {
		log.Fatalln("failed to get an EC2 client ", region, err)
	}
	return ec2Svc
}
