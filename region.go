package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// Region will capture all infomation about this region
type Region struct {
	Name      string
	Instances map[string]*Instance
	Volumes   map[string]*Volume
	Buckets   map[string]*Bucket
}

// regmap is ready to rool
type regmap map[string]*Region

var (
	RegionMap     regmap = make(regmap, 22) // TODO number of regions?
	RegionNames   []string
	currentRegion string
)

// Get region
func (r regmap) Get(region string) *Region {
	if r, e := r[region]; e {
		return r
	}
	return nil
}

// Exists region
func (r regmap) Exists(region string) bool {
	_, e := r[region]
	return e
}

// SetRegion to the current region
func SetRegion(r string) {
	currentRegion = r
}

// Region returns the region we are currently working in
func CurrentRegion() string {
	return currentRegion
}

// Names returns the names of all regions
func Regions() (names []string) {
	if RegionNames == nil {
		if RegionNames = fetchRegions(); RegionNames == nil {
			log.Errorf("expected (regions) got ()")
			return nil
		}
		for _, name := range RegionNames {
			RegionMap[name] = &Region{Name: name}
		}
	}
	return RegionNames
}

// ClearRegions reset
func ClearRegions() {
	RegionNames = nil
	RegionMap = nil
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
