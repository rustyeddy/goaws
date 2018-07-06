package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func GetVolumes(region string) (vdisks map[string]*VDisk) {
	log.Println("GetVolumes from ", region)
	defer log.Printf("  return GetVolumes region %s ", region)

	idxname := region + "-volumes"
	log.Debugf("  looking in local cache for %s ", idxname)

	err := cache.FetchObject(idxname, &vdisks)
	if err == nil && vdisks != nil {
		log.Debugf("  Found cached version of %s .. ", idxname)
		return vdisks
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

	vdisks = vdisksFromAWS(result)
	if vdisks == nil {
		log.Errorf("failed to get vdisks from aws ")
		return nil
	}

	go func() {
		// Save the results in the storage cache
		obj, err := cache.StoreObject(idxname, vdisks)
		log.Infoln("  cache store object idx %s -> %v ", idxname, obj)
		if err != nil || obj == nil {
			log.Errorf("  failed to cache object idx %s -> %v ", idxname, obj)
			return
		}
	}()
	return vdisks
}

func vdisksFromAWS(result *ec2.DescribeVolumesOutput) (vdisks map[string]*VDisk) {
	vdisks = make(map[string]*VDisk, 10)
	volumes := result.Volumes
	for _, vol := range volumes {
		for _, att := range vol.Attachments {
			vd := &VDisk{
				raw:         &vol,
				VolumeId:    *vol.VolumeId,
				SnapshotId:  *vol.SnapshotId,
				Size:        *vol.Size,
				State:       vol.State,
				InstanceId:  *att.InstanceId,
				AttachState: att.State,
				AvailZone:   *vol.AvailabilityZone,
			}
			vdisks[vd.VolumeId] = vd
		}
	}
	return vdisks
}
