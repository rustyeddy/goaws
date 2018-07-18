package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	regions []string
)

// Names returns the names of all regions
func Regions() (names []string, err error) {
	if regions == nil {
		if regions = fetchRegions(); regions == nil {
			return nil, fmt.Errorf("expected (regions) got ()")
		}
	}
	return regions, nil
}

// ClearRegions reset
func ClearRegions() {
	regions = nil
}

// fetchRegionNames from AWS
func fetchRegions() []string {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Errorf("failed to load config, %v", err)
		return nil
	}
	svc := ec2.New(cfg)
	req := svc.DescribeRegionsRequest(&ec2.DescribeRegionsInput{})
	awsRegions, err := req.Send()
	if err != nil {
		log.Errorf("request send failed %v", err)
		return nil
	}

	names := make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		names = append(names, *region.RegionName)
	}
	return names
}

// Client get the EC2 Client for this region
func Client(region string) (ec *ec2.EC2) {
	if cfg, err := external.LoadDefaultAWSConfig(); err != nil {
		log.Fatalf("failed to get ec2 client for region %s ", region)
	} else {
		cfg.Region = region
		return ec2.New(cfg)
	}
	return nil
}
