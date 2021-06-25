package cmd

import (
	"github.com/spf13/cobra"
	"wm-infoflow-api-go/common/log"
	"wm-infoflow-api-go/conf"
	"wm-infoflow-api-go/core"
)

var (
	frontendCmd = &cobra.Command{
		Use: "api",
	}

	frontendStartCmd = &cobra.Command{
		Use: "start",
		Run: apiStart,
	}

	switchCron int
)

func init() {
	rootCmd.AddCommand(frontendCmd)
	frontendCmd.AddCommand(frontendStartCmd)

	frontendStartCmd.PersistentFlags().IntVarP(&switchCron, "switchCron", "c", 0, "开启定时任务")
}

func apiStart(cmd *cobra.Command, args []string) {
	fe := api.New(conf.Config)
	if switchCron == 1 {
		fe.StartCron()
	}
	if err := fe.Start(); err != nil {
		log.Error(err.Error())
	}
}
