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
		Short:   "bucket information",
		Run:     cmdS3,
		Version: "2018-07-18",
	}
)

// s3Cmd for s3 buckets
func cmdS3(cmd *cobra.Command, args []string) {
	if bkts := goaws.ListBuckets(); bkts != nil {
		for _, bkt := range bkts {
			fmt.Printf("  %s - %v\n", bkt.Name(), bkt.Created())
		}
	}
}
