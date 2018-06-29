package goaws

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	regions []string
	region  string
)

// SetCurrentRegion(region string)

// Regions returns a list of region names, this is just a list of
// strings. This function first checks an in-memory copy, then we'll
// check the store, finaly make the client request if and when
// necessary.
func Regions() []string {
	log.Debug("~~> AWS Regions ")
	defer log.Debugf("<~~ Returning AWS Regions %d", len(regions))

	if regions != nil {
		log.Debug("  regions in-memory cache hit ")
		return regions
	}

	log.Debug("  No copy of Regions in memory: checking the cache... ")
	// Check for a local cache of regions
	if !cache.Exists("regions") {
		log.Infoln("  -- cache entry was not found ... ")
	} else {
		log.Infoln("  ~~> cache object found! Fetching it ...")
		err := cache.FetchObject("regions", &regions)
		if err != nil {
			log.Debugf("  ## error fetching regions %v ..", err)
		}
	}

	if regions != nil {
		log.Debugf("  We have regions! %d of em", len(regions))
		return regions
	}

	// go to the source for regions
	log.Debugln("  ~> Nothing local, fetch from AWS ...")
	if regions = fetchRegions(); regions == nil {
		log.Error("  ## failed to get regions from AWS, host is lost ...")
		return nil
	}

	// we have some regions, we'll store them
	log.Debugf("  <~ got %d regions", len(regions))
	log.Debugln("  Store regions list locally ... ")
	if _, err := cache.StoreObject("regions", regions); err != nil {
		log.WithField("error", err).Error("  !! failed to StoreObject regions ", err)
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
	if regions == nil {
		log.Debugf(" No regions saved ")
	} else {
		log.Debugf(" returning with %d regions ", len(regions))
		return regions
	}
	SaveRegions(regions)
	return regions
}

// SaveRegions will store the regions list in our local cache
func SaveRegions(names []string) {
	_, err := cache.StoreObject("regions", names)
	if err != nil {
		log.Errorf("failed saving %v -> %v", names, err)
	}
}

// ReadRegions will read the regions list from the stored object.
func ReadRegions() []string {
	var regs []string
	err := cache.FetchObject("regions", &regs)
	if err != nil {
		log.Error("wanted (list of regions) got (nothing) ", err)
		return nil
	}
	return regs
}
