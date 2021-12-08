package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	// configFile string
	rootCmd = &cobra.Command{
		Use:   "f1",
		Short: "f1-cli is a fast and efficient way to keep up with the Formula 1 events.",
		Long:  "f1-cli is a fast and efficient way to keep up with the Formula 1 events.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)
// 	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.f1.yaml)")
// }

// func initConfig() {
// 	if configFile != "" {
// 		viper.SetConfigFile(configFile)
// 	} else {
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".f1")
// 	}
// 	viper.AutomaticEnv()
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file: ", viper.ConfigFileUsed())
// 	}
// }
