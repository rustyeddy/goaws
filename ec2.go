package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// AWSCloud is confined to a single region
type AWSCloud struct {
	region string
	imap   Instmap
	vmap   Volmap
	smap   Snapmap

	ec2Svc *ec2.EC2

	*log.Logger
}

// TODO - Move ec2.EC2 into AWSCloud ??
type Cloudmap map[string]*AWSCloud
type Instmap map[string]*Instance
type Volmap map[string]*Volume
type Snapmap map[string]*Snapshot
type InstRegion map[string]string

var (
	regions []string

	AWSClouds Cloudmap
	I2R       InstRegion
)

func init() {
	AWSClouds = make(Cloudmap, 20)
	I2R = make(InstRegion, 10)
}

// GetCloud returns the cloud for the given region
func GetCloud(region string) (cl *AWSCloud) {
	if AWSClouds == nil {
		AWSClouds = make(Cloudmap)
	} else if cl, e := AWSClouds[region]; e {
		return cl
	}
	return &AWSCloud{
		region: region,
		imap:   nil,
		vmap:   nil,
		smap:   nil,
	}
}

// Volumes returns the Volumemap
func (cl *AWSCloud) Volumes() Volmap {
	if cl.vmap == nil {
		cl.vmap = GetVolumes(cl.region)
	}
	return cl.vmap
}

// Instances returns the Instmap
func (cl *AWSCloud) Instances() Instmap {
	if cl.imap == nil {
		cl.GetInstances()
	}
	return cl.imap
}

// Snapshots returns the snapshots from AWS
func (cl *AWSCloud) Snapshots() Snapmap {
	if cl.smap == nil {
		cl.smap = GetSnapshots(cl.region)
	}
	return cl.smap
}

// Client get the EC2 Client for this region
func (cl *AWSCloud) Client() (ec *ec2.EC2) {
	if cl.ec2Svc == nil {
		if cfg, err := external.LoadDefaultAWSConfig(); err == nil {
			cfg.Region = cl.region
			cl.ec2Svc = ec2.New(cfg)
		} else {
			log.Fatalf(" Error loading config ")
		}
	}
	return cl.ec2Svc
}

// getEC2 returns an ec2 service for the given region ready for use
func getEC2(region string) (e *ec2.EC2) {
	if cfg, err := external.LoadDefaultAWSConfig(); err == nil {
		cfg.Region = region
		e = ec2.New(cfg)
	}
	return e
}
