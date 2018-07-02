/*
Goa commands aws management utilities.
*/
package cmd

import (
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command{
		Use:   "goa --help",
		Short: "Manage AWS Instances and Volumes",
		Long:  `Gather AWS Volume and Image information, do stuff with it`,
		Run:   GoaDo,
	}
)

// Execute the RootCommand
func Execute() {

	// Best place to set up cache???
	cache := goaws.Cache()
	if cache == nil {
		log.Info("Root initConfig = calling UseStore ", Basedir)
		cache = store.UseStore(Basedir)
	}
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// RunRoot runs the root command
func GoaDo(cmd *cobra.Command, args []string) {
	cache := goaws.Cache() // ???

	log.Println("Welcome to Goa! ")
	log.Println("  basedir  ", Basedir)
	if cache == nil {
		log.Println("  store .. no cache ")
	} else {
		log.Printf("  store %s ", cache.String())
	}
	log.Println(".. Aog! ")
}
