package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// global variable --------------------------------
var configFile string

var rootCmd = &cobra.Command{
	Use:   "PCA_Pi_Server",
	Short: "PCA Server on RaspberryPi",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// ------------------------------------------------

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("./configs/develop")
		viper.AddConfigPath(".")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&configFile,
		"config", "", "config file (default is configs/develop.yml)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// TODO: fix to log
		fmt.Println(err)
		os.Exit(1)
	}

}
