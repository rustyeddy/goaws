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

// Bucket the name of bucket and objects
type Bucket struct {
	name    string    // bucket name
	created time.Time // creation time
	objects []Object
	s3.Bucket
}

var (
	allBuckets []Bucket
)

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

	svc := s3svc(currentRegion)
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
		allBuckets = append(allBuckets, *bkt)
	}
	return allBuckets
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

// GetObjects returns all (up to 1000) objects from a specific bucket
func GetObjects(region, bktname string) (olist []Object, err error) {
	svc := s3svc(region)
	req := svc.ListObjectsRequest(&s3.ListObjectsInput{Bucket: &bktname})
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	for _, cont := range resp.Contents {
		//obj := objFromContent(cont)
		//olist = append(olist, &obj)
		log.Fatalf(" %+v ", cont)
	}
	return olist, nil
}

// Objects from Content
func objFromContent(cont interface{}) (obj *Object) {
	log.Fatalf("  %+v ", cont)
	return obj
}
