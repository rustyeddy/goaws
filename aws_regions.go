package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// Regions returns the list of regions, it will first check an
// in memory copy, it will then check for cached file copy, finally
// we'll head to AWS got this info.
func Regions() []string {
	if regions == nil {
		if regions = readRegions("regions.json"); regions == nil {
			regions = fetchRegions()
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

	regions := make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}
	saveRegions("regions.json", regions)
	return regions
}

// saveRegions
func saveRegions(path string, names []string) {
	if jb, err := json.Marshal(names); err == nil {
		if err = ioutil.WriteFile(path, jb, 0644); err == nil {
			log.Debug("saved regions.json ")
		}
		log.Error("failed write ", path, err)
	}
	log.Error("failed to marshal names ")
}

// readRegions
func readRegions(path string) (names []string) {
	var buf []byte
	var err error

	if buf, err = ioutil.ReadFile(path); err == nil {
		if err = json.Unmarshal(buf, names); err == nil {
			return names
		}
	}
	log.Error("failed to read ", path, err)
	return nil
}
