package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sih-ipfs",
	Short: "SIH - IPFS project",
	Long:  "SIH - IPFS project to demonstrate compression and data transfer through a network",
}

var configFile string

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// we'll define global configs in here.
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file to be used by the application")
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sih")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using Config file: ", viper.ConfigFileUsed())
	}
}
