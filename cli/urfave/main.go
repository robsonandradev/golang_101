package main

import (
  "os"
  "fmt"
  "github.com/urfave/cli/v2"
)

type ArgsVars struct {
  version bool
}

func main() {
  argsVars := ArgsVars{}
  versionFlagData := versionFlag(&argsVars)
  flags := []cli.Flag{versionFlagData}
  app := createApp(&argsVars, flags)
  err := app.Run(os.Args)
  if err != nil {
    panic(err)
  }
}

func versionFlag(argsVars *ArgsVars) *cli.BoolFlag {
  return &cli.BoolFlag{
    Name: "version",
    Value: false,
    Usage: "shows app version",
    Aliases: []string{"v"},
    Destination: &argsVars.version,
  }
}

func createApp(argsVars *ArgsVars, flags []cli.Flag) *cli.App {
  return &cli.App{
    Name: "app",
    Usage: "My awsome cli",
    Flags: flags,
    Action: func(ctx *cli.Context) error {
      return cmdAction(ctx, argsVars)
    },
  }
}

func cmdAction(ctx *cli.Context, argsVars *ArgsVars) error {
      if argsVars.version {
        fmt.Println("version 1.0")
        return nil
      }
      return cli.ShowAppHelp(ctx)
}
