package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	"github.com/spf13/cobra"
)

var (
	// GoaCmd is the root command
	s3Cmd = cobra.Command{
		Use:     "s3",
		Aliases: []string{"buckets"},
		Short:   "bucket information",
		Run:     cmdS3,
		Version: "2018-07-18",
	}

	s3ObjectsCmd = cobra.Command{
		Use:     "objects",
		Aliases: []string{"objs"},
		Short:   "objects in bucket",
		Run:     cmdS3Objects,
		Version: "2018-07-19",
	}
)

func init() {
	s3Cmd.AddCommand(&s3ObjectsCmd)
}

// s3Cmd for s3 buckets
func cmdS3(cmd *cobra.Command, args []string) {
	fmt.Println("CMD S3 ")
	bkts := goaws.ListBuckets()
	if bkts == nil {
		fmt.Println("No buckets found in AWS account")
		return
	}
	fmt.Printf("Buckets..[%d]\n", len(bkts))
	if bkts := goaws.ListBuckets(); bkts != nil {
		for _, bkt := range bkts {
			bname := bkt.Name()
			fmt.Println("Bucket ", bname)
		}
		fmt.Printf("\n")
	}
}

// s3ObjectsCmd for s3 objects
func cmdS3Objects(cmd *cobra.Command, args []string) {
	fmt.Println("CMD S3 Objects ")
	if bkts := goaws.ListBuckets(); bkts != nil {
		for _, bkt := range bkts {

			for _, region := range goaws.RegionNames {
				// Must be in correct region ...
				if objs, err := goaws.GetObjects(region, bkt.Name()); err == nil {
					fmt.Println(bkt.Name, objs)

					if reg := goaws.RegionMap.Get(region); reg != nil {
						reg.Buckets[bkt.Name()] = &bkt
					}
				}
			}
		}
	}
}
