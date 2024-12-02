// Copyright 2018 fatedier, fatedier@gmail.com
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
	"netholer/client"
	"netholer/server"
	"os"

	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/spf13/cobra"
)

var (
	cfgFile          string
	cfgDir           string
	showVersion      bool
	strictConfigMode bool
	rootCmd          = &cobra.Command{
		Use:   "netholer",
		Short: "netholer is the net holer",
	}
	serverCfg v1.ServerConfig
	isServer  bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./conf.toml", "config file")
	rootCmd.PersistentFlags().StringVarP(&cfgDir, "config_dir", "d", "", "config directory, run one matec service for each file in config directory")
	rootCmd.PersistentFlags().BoolVarP(&isServer, "server", "s", false, "is server")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolVarP(&strictConfigMode, "strict_config", "", false, "strict config parsing mode, unknown fields will cause an errors")
	config.RegisterServerConfigFlags(rootCmd, &serverCfg)
}

func main() {
	system.EnableCompatibilityMode()
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version.Full())
			return nil
		}

		if isServer {
			fmt.Println("run as server")
			server.Run(cfgFile, serverCfg, strictConfigMode)
			return nil
		}
		fmt.Println("run as client")
		for _, arg := range cmd.Flags().Args(){
			fmt.Println(arg)
		}

		client.Run(cfgFile, cfgDir, strictConfigMode)
		// client.Register(rootCmd)
		return nil
	}
	execute()
}
func execute() {
	rootCmd.SetGlobalNormalizationFunc(config.WordSepNormalizeFunc)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
