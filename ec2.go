package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

type clientMap map[string]*ec2.EC2
type cloudMap map[string]*AWSCloud
type Instmap map[string]*Instance
type Volmap map[string]*Volume
type Snapmap map[string]*Snapshot

var (
	regions      []string
	awsClients   clientMap
	awsClouds    cloudMap
	allInstances Instmap
	allVolumes   Volmap
	allSnapshots Snapmap
)

func init() {
	awsClients = make(clientMap)
	awsClouds = make(cloudMap)

	allInstances = make(Instmap)
	allVolumes = make(Volmap)
	allSnapshots = make(Snapmap)
}

// AWSCloud is confined to a single region
type AWSCloud struct {
	region string
	Instmap
	Volmap
	Snapmap
}

// GetCloud returns the cloud for the given region
func GetCloud(region string) (cl *AWSCloud) {
	if cl, e := awsClouds[region]; e {
		return cl
	}
	return &AWSCloud{region: region}
}

func (cl *AWSCloud) Volumes() Volmap {
	if cl.Volmap == nil {
		cl.Volmap = GetVolumes(cl.region)
	}
	return cl.Volmap
}

func (cl *AWSCloud) Instances() Instmap {
	if cl.Instmap == nil {
		cl.Instmap = GetInstances(cl.region)
	}
	return cl.Instmap
}

func (cl *AWSCloud) Snapshots() Snapmap {
	if cl.Snapmap == nil {
		cl.Snapmap = GetSnapshots(cl.region)
	}
	return cl.Snapmap
}

// getEC2 returns an ec2 service for the given region ready for use
func getEC2(region string) (ec2Svc *ec2.EC2) {
	log.Debugln("Get EC2 for region ", region)
	defer log.Debugln(" leaving EC2 %v ", ec2Svc)

	// If we have a copy return it
	if svc, e := awsClients[region]; e {
		log.Debugln("  using cached EC2 client for %s ", region)
		return svc
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("  Failed to Load Default AWS Config %q -> %v ", region, err)
		return nil
	}

	cfg.Region = region
	ec2Svc = ec2.New(cfg)
	if ec2Svc == nil {
		log.Fatalf("  expected EC2 client for %s got %s", region, err)
	}
	return ec2Svc
}
