package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// getEC2 returns an ec2 service ready for use
func getEC2(name string) (ec2Svc *ec2.EC2) {
	log.Debugln("Get EC2 for region ", region)
	defer log.Debugln(" leaving EC2 %v ", ec2Svc)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("  Failed to Load Default AWS Config %q -> %v ", name, err)
		return nil
	}

	log.Debugf("  loaded Default config, create ec2 client ")

	//cfg.Region = endpoints.UsWest2RegionID
	ec2Svc = ec2.New(cfg)
	if ec2Svc == nil {
		log.Fatalln("failed to get an EC2 client ", name, err)
	}
	return ec2Svc
}
