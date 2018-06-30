package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

var (
	volumes []*ec2.DescribeVolumesOutput
	volmap  map[string]*ec2.DescribeVolumesOutput
)

func init() {
	volmap = make(map[string]*ec2.DescribeVolumesOutput)
}

// FetchVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func FetchVolumes(region string) (volmap map[string]*ec2.DescribeVolumesOutput) {
	log.Println("GetVolumes from ", region)
	defer log.Printf("  return GetVolumes region %s size %d ", region, len(volmap))

	// First, check local cache ...
	if cache == nil {
		log.Debug("  did not find volumes in cache ")
	}

	idxname := region + "-volumes"
	var volumes *ec2.DescribeVolumesOutput

	log.Debugf("  looking in local cache for %s ", idxname)

	// Check our local cache first
	err := cache.FetchObject(idxname, volumes)
	if err == nil && volumes != nil {
		log.Debugf("  Found cached version of %s .. ", idxname)
		volmap[region] = volumes
		return volmap
	}

	log.Debugf("  Fetch Volumes from AWS for %s", region)
	var e *ec2.EC2
	if e = GetEC2(region); e == nil {
		log.Errorf("  failed to get aws client for %s ", region)
		return nil
	}

	log.Debugf("  sending request for volumes region ", region)
	req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	if result, err := req.Send(); err != nil {
		log.Errorf("  # failed response to request ")
		return nil
	} else {
		log.Debugf("  got result from region %s ", region)
		if volmap == nil {
			volmap = make(map[string]*ec2.DescribeVolumesOutput)
		}
		volmap[region] = result
	}

	// Save the results in the storage cache
	obj, err := cache.StoreObject(idxname, &volmap)
	if err != nil || obj == nil {
		log.Errorf("  failed to cache object %s", idxname)
		return nil
	}
	log.Debugf("  retrieved %s -> %v", volmap)
	return volmap
}
