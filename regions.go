package goaws

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	regions []string
)

func init() {
}

// Regions returns a list of region names, this is just a list of
// strings. This function first checks an in-memory copy, then we'll
// check the store, finaly make the client request if and when
// necessary.
func Regions() []string {
	log.Debug("~~> Getting AWS Regions ")
	defer log.Debugf("<~~ Returning AWS Regions %d", len(regions))

	if regions == nil {
		log.Debug("  regions are nil, must go looking ... ")
		if err := S.FetchObject("regions", &regions); err != nil {
			// If the store cache fails, don't return but move on to AWS
			log.Error("Fetch object from store failed", err)
		}

		// go to the source for regions
		log.Debug("  could not be found in local store, fetching from AWS ...")
		if regions = fetchRegions(); regions == nil {
			log.Error("failed to get regions from AWS, host is lost ...")
			return nil
		}

		// we have some regions, we'll store them
		log.Debug("  AWS retrieve success got %d regions", len(regions))
		log.Debug("  Store AWS regions in local store ... ")
		if _, err := S.StoreObject("regions", regions); err != nil {
			log.Error("failed to StoreObject regions ", err)
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

	regions = nil // reset and fetch if not already nil
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

	regions = make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}
	return regions
}

// saveRegions
// =======================================================
func SaveRegions(names []string) {
	_, err := S.StoreObject("regions", names)
	if err != nil {
		log.Error("failed saving %v -> %v", names, err)
	}
}

// readRegions
// =======================================================
func ReadRegions() []string {
	var regs []string
	err := S.FetchObject("regions", &regs)
	if err != nil {
		log.Error("wanted (list of regions) got (nothing) ", err)
		return nil
	}
	return regs
}

func oldSaveRegions(path string, names []string) {
}

// readRegions
// =======================================================
func oldReadRegions(path string) (names []string) {

	log.Debugf(" ~~~ readRegions path %s ~~~ ", path)
	defer log.Debugf("   leaving with names %+v", names)

	var buf []byte
	var err error
	if buf, err = ioutil.ReadFile(path); err == nil {
		if err = json.Unmarshal(buf, names); err == nil {
			return names
		}
	}
	if err != nil {
		log.Errorf("failed to read %s err %v", path, err)
	}
	return nil
}
