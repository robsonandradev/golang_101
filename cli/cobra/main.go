package main

import (
  "github.com/spf13/cobra"
  "robsonandradev/cmd"
)

func main() {
  testCMD()
}

func testCMD() {
  root := &cobra.Command{Use: "app"}
  cmd.Version(root)
  //cmd.VersionFlag(root)
  root.Execute()
}
