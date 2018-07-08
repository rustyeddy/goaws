package goaws

import (
	log "github.com/rustyeddy/logrus"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// FetchInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func GetInstances(region string) (vms map[string]*VM) {
	log.Debugln("~~> GetInstances for region ", region)
	defer log.Debugln("  <~~ return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	e := getEC2(region)
	if e == nil {
		log.Errorf("  failed to get an EC2 client for ", region)
		return nil
	}

	// Look for a cached version of the object
	idxname := region + "-instances"
	err := cache.FetchObject(idxname, &vms)
	if err == nil && vms != nil {
		log.Debugf("  found cached version of %s .. ", idxname)
		return vms
	}

	log.Debugf("  fetch instance data from AWS %s ", region)
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	if vms = vmsFromAWS(result); vms == nil {
		log.Errorf("  failed to get vms from AWS %v", err)
		return nil
	}

	// Store the object for later cache usage
	go func() {
		obj, err := cache.StoreObject(idxname, vms)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
		}
		log.Debugf("  object cached at path %s ", obj.Path)
	}()
	return vms
}

func vmsFromAWS(result *ec2.DescribeInstancesOutput) (vms map[string]*VM) {
	vms = make(map[string]*VM, 100)
	nextToken := result.NextToken
	resvs := result.Reservations
	for _, resv := range resvs {
		for _, inst := range resv.Instances {
			vm := &VM{
				InstanceId: *inst.InstanceId,
				State:      *inst.State,
				KeyName:    *inst.KeyName,
			}
			for _, bdm := range inst.BlockDeviceMappings {
				vm.VolumeId = *bdm.Ebs.VolumeId
			}
			vms[vm.InstanceId] = vm
		}
	}
	if nextToken != nil {
		panic("next token != nil ")
	}
	return vms
}

// DeleteInstance
func DeleteInstance(instId string) error {

	/*

		var svc *ec2.EC2
		if svc = getEC2(region); svc == nil {
			log.Errorf("  failed to get aws client for %s ", region)
			return nil
		}

		log.Debugf("  sending request to delete instance region ", region)
		req := svc.DeleteInstanceRequest(&ec2.DeleteInstanceInput{VolumeId: aws.String(instId)})
		result, err := req.Send()
		if err != nil {
			return fmt.Errorf("  # failed response to request %v", err)
		}
		log.Debugf("  got result %v from region %s ", result, region)
		log.Fatalf("  result %+v", result)
	*/
	return nil
}
