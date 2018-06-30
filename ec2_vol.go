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
func FetchVolumes(region string) ([]*ec2.DescribeVolumesOutput, map[string]*ec2.DescribeVolumesOutput) {
	log.Println("GetVolumes from ", region)
	defer log.Println("  return GetVolumes ", region)

	// First, check local cache ...
	if cache == nil {
		log.Debug("  did not find volumes in cache ")
	}

	// Get the volumes
	if e := GetEC2(region); e != nil {
		req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
		if result, err := req.Send(); err == nil {
			volmap[region] = result
			volumes = append(volumes, result)
		} else {
			log.Errorf("  # failed response to request ")
			return nil, nil
		}
	}
	return volumes, volmap
}
