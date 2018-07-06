package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func GetVolumes(region string) []VDisk {
	log.Println("GetVolumes from ", region)
	defer log.Printf("  return GetVolumes region %s ", region)

	idxname := region + "-volumes"
	log.Debugf("  looking in local cache for %s ", idxname)

	var volumes []VDisk
	err := cache.FetchObject(idxname, &volumes)
	if err == nil && volumes != nil {
		log.Debugf("  Found cached version of %s .. ", idxname)
		return volumes
	}

	log.Debugf("  Fetch Volumes from AWS for %s", region)
	var e *ec2.EC2
	if e = getEC2(region); e == nil {
		log.Errorf("  failed to get aws client for %s ", region)
		return nil
	}

	log.Debugf("  sending request for volumes region ", region)
	req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  # failed response to request %v", err)
		return nil
	}
	log.Debugf("  got result %v from region %s ", result, region)

	vdisks := vdisksFromAWS(result)
	if vdisks == nil {
		log.Errorf("failed to get vdisks from aws ")
		return nil
	}

	go func() {
		// Save the results in the storage cache
		obj, err := cache.StoreObject(idxname, volumes)
		if err != nil || obj == nil {
			log.Errorf("  failed to cache object %s", idxname)
			return
		}
	}()
	return vdisks
}

func vdisksFromAWS(result *ec2.DescribeVolumesOutput) (vd []VDisk) {
	//volumes = ec2.DescribeVolumesOutput

	log.Fatal(" %+v ", result)

	/*
		for _, vol := range results {
			vinfo = map[string]string{
				"VolumeId":         *vol.VolumeId,
				"SnapshotId":       *vol.SnapshotId,
				"AvailabilityZone": *vol.AvailabilityZone,
				"State":            string(vol.State),
			}
			for _, att := range vol.Attachments {
				vinfo["AttachVolumeId"] = *att.VolumeId
				vinfo["InstanceId"] = *att.InstanceId
				//vinfo["AttachState"] = aws.String(*att.State)
				mdisk[vol.VolumeId] = vol
			}
		}
	*/
	return vd
}
