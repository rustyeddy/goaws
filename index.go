package goaws

/*
type Inventory struct {
	IV map[string]*ec2.InstanceId
}

// Index the inventory
func (inv *Inventory) Index() {

	// index the instances
	for _, inst := range inv.Instances {
		iid := *aws.String(*inst.InstanceId)
		vid := *aws.String(*inst.BlockDeviceMappings[0].Ebs.VolumeId)
		inv.IV[iid] = vid
	}

	// read and process volumes
	for _, vol := range inv.Volumes {
		vid := *aws.String(*vol.VolumeId)
		iid := *aws.String(*vol.Attachments[0].InstanceId)
		inv.VI[vid] = iid
	}
}

// indexInstances
func (inv *Inventory) indexInstances(rlist []ec2.RunInstancesOutput) {
	for _, ilist := range rlist {
		for _, inst := range ilist.Instances {
			iid := *inst.InstanceId
			inv.Instances[iid] = HostFromInstance(&inst)
			if inst.BlockDeviceMappings != nil {
				bm0 := inst.BlockDeviceMappings[0]
				ebs := bm0.Ebs
				vid := ebs.VolumeId
				inv.IV[iid] = *vid
			} else {
				inv.IV[iid] = ""
			}
		}
	}
}

// indexVolumes
func (inv *Inventory) indexVolumes(vols []ec2.CreateVolumeOutput) {
	// Index teh volmes and volume to image map
	for _, vol := range vols {
		vid := *vol.VolumeId
		inv.Volumes[vid] = DiskFromVolume(&vol)
		if vol.Attachments != nil {
			a := vol.Attachments[0]
			inv.VI[vid] = *a.InstanceId
		} else {
			inv.VI[vid] = ""
		}
	}
}
*/
