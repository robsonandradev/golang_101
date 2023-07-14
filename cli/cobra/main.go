package main

import (
  "fmt"
  "github.com/spf13/cobra"
  //"github.com/spf13/viper"
)

func main() {
  version()
}

func version() {
  root := &cobra.Command{
    Use: "version",
    Short: "Get the app version",
    Long: "Get the app version",
    Run: func (cmd *cobra.Command, args []string) {
      fmt.Println("Version 0.0.1")
    },
  }
  //var version string
  root.PersistentFlags().Bool("version", true, "Get the app version")
  if err := root.Execute(); err != nil {
    fmt.Errorf("something want wrong %s", err)
  }
}
