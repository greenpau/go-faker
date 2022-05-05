// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/greenpau/versioned"
)

var (
	app        *versioned.PackageManager
	appVersion string
	gitBranch  string
	gitCommit  string
	buildUser  string
	buildDate  string
	sh         *cli.App
)

func init() {
	app = versioned.NewPackageManager("faker")
	app.Description = "Fake user data generator"
	app.Documentation = "https://github.com/greenpau/go-faker/"
	app.SetVersion(appVersion, "1.0.1")
	app.SetGitBranch(gitBranch, "main")
	app.SetGitCommit(gitCommit, "2a68448")
	app.SetBuildUser(buildUser, "")
	app.SetBuildDate(buildDate, "")

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(os.Stdout, "%s\n", app.Banner())
	}

	sh = cli.NewApp()
	sh.Name = app.Name
	sh.Version = app.Version
	sh.Usage = app.Description
	sh.Description = app.Documentation
	sh.HideHelp = false
	sh.HideVersion = false
	sh.Flags = append(sh.Flags, &cli.StringFlag{Name: "domain", Usage: "set email domain"})
	sh.Flags = append(sh.Flags, &cli.IntFlag{Name: "count", Usage: "set user count", Value: 1})
	sh.Commands = []*cli.Command{
		{
			Name:   "user",
			Usage:  "generate user",
			Action: NewUser,
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "password", Usage: "set password"},
				&cli.StringFlag{Name: "email-prefix", Usage: "set email address prefix"},
				&cli.StringSliceFlag{Name: "roles", Usage: "set user roles", Value: cli.NewStringSlice("viewer")},
			},
		},
	}
}

func main() {
	sh.Run(os.Args)
}
