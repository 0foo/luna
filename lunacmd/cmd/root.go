package cmd

import (
	"fmt"
	"luna/cmd/lunadb"
	"luna/cmd/lunamigrate"
	"luna/cmd/lunaseed"
	"luna/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RootCmd = &cobra.Command{
		Use:   "luna",
		Short: "Luna CLI tool",
		Long:  "Luna: a CLI tool for managing migrations and future extensions.",
	}
	configFile     string
)

func Execute() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config file (default is ./config.yaml)")

	RootCmd.AddCommand(lunamigrate.Cmd)
	RootCmd.AddCommand(lunaseed.Cmd)
	RootCmd.AddCommand(lunadb.Cmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig loads the configuration
func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config.ConfigValues); err != nil {
		fmt.Println("Unable to decode config into struct:", err)
		os.Exit(1)
	}

}
