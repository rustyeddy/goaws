package goaws

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	regions       []string
	currentRegion string
)

// CurrentRegion is a global convenience to keep track of the
// Region we are currently working on.
func CurrentRegion() string {
	currentRegion = "us-west-2"
	return currentRegion
}

// SetCurrentRegion(region string)

// Regions returns a list of region names, this is just a list of
// strings. This function first checks an in-memory copy, then we'll
// check the store, finaly make the client request if and when
// necessary.
func Regions() []string {
	log.Debug("~~> Getting AWS Regions ")
	defer log.Debugf("<~~ Returning AWS Regions %d", len(regions))

	if regions == nil {

		log.Debug("  regions are nil, checking local store ... ")

		// Check for a local cache of regions
		if !cache.Exists("regions") {
			log.Infoln("  ## regions list is NOT cached, going AWS ")
		} else {
			if err := cache.FetchObject("regions", &regions); err != nil {
				log.Debugf("  ## error fetching object ")
				return nil
			}
		}

		// go to the source for regions
		log.Debugln("  ~> Nothing local, fetch from AWS ...")
		if regions = fetchRegions(); regions == nil {
			log.Error("  ## failed to get regions from AWS, host is lost ...")
			return nil
		}

		// we have some regions, we'll store them
		log.Debugf("  <~ got %d regions", len(regions))
		log.Fatalf("%+v", regions)

		log.Debugln("  Store regions list locally ... ")
		if _, err := cache.StoreObject("regions", regions); err != nil {
			log.WithField("error", err).Error("  !! failed to StoreObject regions ", err)
		}
	}
	return regions
}

// String will print regions
func String() string {
	rs := Regions()
	return strings.Join(rs, "\n")
}

// fetchRegionNames from AWS
func fetchRegions() []string {

	log.Debug("~~> fetchRegions entered ")
	defer log.Debugf("<~~ fetchRegions exiting ")
	regions = nil // reset and fetch if not already nil

	log.Debug("  -- loading Default AWS Config ")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Errorf("failed to load config, %v", err)
		return nil
	}
	svc := ec2.New(cfg)
	req := svc.DescribeRegionsRequest(&ec2.DescribeRegionsInput{})

	log.Debug("  -- sending request for Regions to AWS ")
	awsRegions, err := req.Send()
	if err != nil {
		log.Errorf("request send failed %v", err)
		return nil
	}

	log.Debug("  -- process response from AWS  ")
	regions = make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}
	log.Debugf("  ")
	return regions
}

// saveRegions
// =======================================================
func SaveRegions(names []string) {
	_, err := cache.StoreObject("regions", names)
	if err != nil {
		log.Error("failed saving %v -> %v", names, err)
	}
}

// readRegions
// =======================================================
func ReadRegions() []string {
	var regs []string
	err := cache.FetchObject("regions", &regs)
	if err != nil {
		log.Error("wanted (list of regions) got (nothing) ", err)
		return nil
	}
	return regs
}
