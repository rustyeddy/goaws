package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	regions []string
	region  string
)

// SetRegion to the current region
func SetRegion(r string) {
	region = r
}

// Region returns the region we are currently working in
func Region() string {
	return region
}

// Names returns the names of all regions
func Regions() (names []string) {
	if regions == nil {
		if regions = fetchRegions(); regions == nil {
			log.Errorf("expected (regions) got ()")
			return nil
		}
	}
	return regions
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
