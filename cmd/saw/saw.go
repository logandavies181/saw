package main

import (
	"github.com/logandavies181/slaw/config"
	"github.com/spf13/cobra"
)

// SawCommand is the main top-level command
var sawCommand = &cobra.Command{
	Use:   "saw <command>",
	Short: "A fast, multipurpose tool for AWS CloudWatch Logs",
	Long:  "Saw is a fast, multipurpose tool for AWS CloudWatch Logs.",
	Example: `  saw groups
  saw streams production
  saw watch production`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var awsConfig config.AWSConfiguration

func init() {
	sawCommand.AddCommand(groupsCommand)
	sawCommand.AddCommand(streamsCommand)
	sawCommand.AddCommand(versionCommand)
	sawCommand.AddCommand(watchCommand)
	sawCommand.AddCommand(getCommand)
	sawCommand.AddCommand(queryCommand)
	sawCommand.PersistentFlags().StringVar(&awsConfig.Region, "region", "", "override profile AWS region")
	sawCommand.PersistentFlags().StringVar(&awsConfig.Profile, "profile", "", "override default AWS profile")
}
