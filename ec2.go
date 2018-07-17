package goaws

var (
	region    string // current region
	regions   []string
	instances map[string]*Instance
	volumes   map[string]*Volume
)

func init() {
}

// Instances returns the Instmap
func Instances(region string) map[string]*Instance {
	if instances == nil {
		instances = FetchInstances(region)
	}
	return instances
}

// Volumes returns the Volumemap
func Volumes(region string) map[string]*Volume {
	if volumes == nil {
		volumes = FetchVolumes(region)
	}
	return volumes
}
