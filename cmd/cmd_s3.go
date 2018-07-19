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

// s3Cmd for s3 buckets
func cmdS3(cmd *cobra.Command, args []string) {
	bkts := goaws.ListBuckets()
	if bkts == nil {
		fmt.Println("No buckets found in AWS account")
		return
	}

	fmt.Printf("Buckets..[%d]\n", len(bkts))
	if bkts := goaws.ListBuckets(); bkts != nil {
		objcount := 0
		totalsize := 0

		for _, bkt := range bkts {
			if objects = goaws.GetObjects(); objects != nil {
				for _, obj := range objects {
					totalsize += obj.Size()
					objcount++
				}
			}
		}
		fmt.Printf(" objects %d size %d \n", objcount, totalsize)
	}
}

// s3ObjectsCmd for s3 objects
func cmdS3Objects(cmd *cobra.Command, args []string) {
	if bkts := goaws.ListBuckets(); bkts != nil {
		for _, bkt := range bkts {
			if objs := goaws.GetObjects( /*bkt.Name*/ ); objs != nil {
				fmt.Println(bkt.Name, objs)
			}
		}
	}
}
