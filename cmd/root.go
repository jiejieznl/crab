package main

import (
	"crab/cmd/internal"
	"crab/pkg"
	"crab/pkg/config"
	"crab/pkg/server"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	showVersion bool
	isMigrate   bool
	isGenConfig bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file of crab")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version of crab")
	rootCmd.PersistentFlags().BoolVarP(&isMigrate, "migrate", "m", false, "Migrate database")
	rootCmd.PersistentFlags().BoolVarP(&isGenConfig, "genconfig", "g", false, "Generate default configuration file")
}

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "Crab",
	Short: "o.0?",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version)
			return nil
		}

		if isGenConfig {
			internal.GenDefaultConfig()
			return nil
		}

		if len(cfgFile) == 0 {
			config.InitConfig("")
		} else {
			config.InitConfig(cfgFile)
		}

		if isMigrate {
			internal.InstallTables()
			return nil
		}

		start()
		return nil
	},
}

func execute() {
	err := rootCmd.Execute()
	rootCmd.AddCommand()
	if err != nil {
		panic(err)
	}
}

func start() {
	pkg.Initialize()
	app := server.New()
	moduleInstall(app)
	server.Run(app)
}
