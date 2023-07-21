package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  //"github.com/spf13/viper"
)

func Version(root *cobra.Command) {
  versionCmd := &cobra.Command{
    Use: "version",
    Short: "Get the app version",
    Long: "Get the app version",
    Run: func (cmd *cobra.Command, args []string) {
      fmt.Println("Version 0.0.1")
    },
  }
  //root.PersistentFlags().BoolP("version", "v", true, "Get the app version")
  root.AddCommand(versionCmd)
}

func VersionFlag(root *cobra.Command) {
  versionFlag := root.Flag("version")
  root.Flags().AddFlag(versionFlag)
}
