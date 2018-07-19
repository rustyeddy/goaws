package goaws

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/rustyeddy/logrus"
)

// Object that holds things
type Object struct {
	name string
	path string
	size int
}

// Name of the object
func (o *Object) Name() string {
	return o.name
}

// Path of the object in storage
func (o *Object) Path() string {
	return o.path
}

// Size of the object
func (o *Object) Size() int {
	return o.size
}

// Bucket the name of bucket and objects
type Bucket struct {
	name    string    // bucket name
	created time.Time // creation time
	objects []Object
	s3.Bucket
}

// Name of the bucket
func (b *Bucket) Name() string {
	return b.name
}

// Objects are returned
func (b *Bucket) Objects() []Object {
	return b.objects
}

// Created returns a string representing when the bucket was created
func (b *Bucket) Created() time.Time {
	return b.created
}

// Client get the EC2 Client for this region
func s3svc(region string) (ec *s3.S3) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to get ec2 client for region %s ", region)
	}
	cfg.Region = region
	return s3.New(cfg)
}

// ListBuckets
func ListBuckets() (bkts []Bucket) {
	svc := s3svc(region)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	results, err := req.Send()
	if err != nil {
		log.Errorf("  failed to list buckets")
		return nil
	}
	for _, b := range results.Buckets {
		bkt := bktFromAWS(b)
		if bkt == nil {
			log.Errorf("  failed to create object from bucket ")
		}
		bkts = append(bkts, *bkt)
	}
	return bkts
}

// create a bucket from aws
func bktFromAWS(b s3.Bucket) (bkt *Bucket) {
	bkt = &Bucket{
		name:    *b.Name,
		created: *b.CreationDate,
		objects: nil,
		Bucket:  b,
	}
	return bkt
}
