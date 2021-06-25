package cmd

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"wm-infoflow-api-go/common/log"
	"wm-infoflow-api-go/conf"
)

var (
	cfgFile string
	env     string

	rootCmd = &cobra.Command{
		Use:   "wm-infoflow-api-go",
		Short: "信息流API",
	}

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)


func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./conf/config.toml)")
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "prod", "env setting")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if _, err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Config)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
