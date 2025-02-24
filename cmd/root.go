package cmd

import (
	"os"
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "hani",
	Short: "HAnime downloader",
	Long:  `HAnime downloader. Repo: https://github.com/acgtools/hanime-hunter`,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	oldMask := syscall.Umask(0)
	fmt.Printf("Current UMASK: %04o\n", oldMask)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("log-level", "info", "log level, options: debug, info, warn, error, fatal")

	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.AddCommand(verCmd, dlCmd)
}
