package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// Snapshot is created from a Volume
type Snapshot struct {
	SnapshotId string
	VolumeId   string
	State      string
	Size       int
	StartTime  string
	OwnerId    string
}

// GetSnapshots will retrieve the snapshots from the cache or
// directly from AWS if the cache is missing
func GetSnapshots(region string) (sn Snapmap) {
	log.Printf("Get snapshots from AWS %s", region)
	defer log.Printf("  return from snapshots")

	idxname := region + "-snaps"
	log.Debugf("  looking in local cache for %s ", idxname)
	err := cache.FetchObject(idxname, &sn)
	if err != nil {
		log.Info("  cache miss %v ", err)
		return nil
	}
	log.Debugf("  Fetch Volumes from AWS for %s", region)
	var e *ec2.EC2
	if e = getEC2(region); e == nil {
		log.Errorf("  failed to get aws client for %s ", region)
		return nil
	}

	log.Debugf("  sending request for volumes region ", region)
	req := e.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  # failed response to request %v", err)
		return nil
	}
	log.Debugf("  got result %v from region %s ", result, region)

	smap := snapsFromAWS(result, region)
	if smap == nil {
		log.Errorf("failed to get vdisks from aws ")
		return nil
	}
	go func() {
		// Save the results in the storage cache
		obj, err := cache.StoreObject(idxname, smap)
		log.Debugf("  cache store object idx %s -> %v ", idxname, obj)
		if err != nil || obj == nil {
			log.Errorf("  failed to cache object idx %s -> %v ", idxname, obj)
			return
		}
	}()
	return sn
}

func snapsFromAWS(result *ec2.DescribeSnapshotsOutput, region string) (smap Snapmap) {
	/*
				for _, vol := range result.Volumes {
					for _, att := range vol.Attachments {
						vol := &Volume{
							raw:         &vol,
							VolumeId:    *vol.VolumeId,
							SnapshotId:  *vol.SnapshotId,
							Size:        *vol.Size,
							State:       vol.State,
							InstanceId:  *att.InstanceId,
							AttachState: att.State,
							AvailZone:   *vol.AvailabilityZone,
		Region: region
						}
						vmap[vol.VolumeId] = vol
						allVolumes[vol.VolumeId] = vol
					}
				}
	*/
	return smap
}

// DeleteSnapshot will do that
func DeleteSnapshot(region, snapid string) error {
	var svc *ec2.EC2
	if svc = getEC2(region); svc == nil {
		return fmt.Errorf("  failed to get aws client for %s ", region)
	}

	// Create and send request to delete snapshot
	req := svc.DeleteSnapshotRequest(&ec2.DeleteSnapshotInput{SnapshotId: aws.String(snapid)})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  # failed response to request %v", err)
		return nil
	}
	log.Debugf("  got result %v from region %s ", result, region)
	log.Fatalf("  result %+v", result)
	return nil
}
