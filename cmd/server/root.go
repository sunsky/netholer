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

package server

import (
	"context"
	"fmt"
	"os"

	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
)

func Run(cfgFile string, serverCfg v1.ServerConfig, strictConfigMode bool) error {
	var (
		svrCfg         *v1.ServerConfig
		isLegacyFormat bool
		err            error
	)
	if cfgFile != "" {
		svrCfg, isLegacyFormat, err = config.LoadServerConfig(cfgFile, strictConfigMode)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isLegacyFormat {
			fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
				"please use yaml/json/toml format instead!\n")
		}
	} else {
		serverCfg.Complete()
		svrCfg = &serverCfg
	}

	warning, err := validation.ValidateServerConfig(svrCfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		log.Infof("mates uses config file: %s", cfgFile)
	} else {
		log.Infof("mates uses command line arguments for config")
	}

	if err := runServer(svrCfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func runServer(cfg *v1.ServerConfig) (err error) {
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)

	svr, err := server.NewService(cfg)
	if err != nil {
		return err
	}
	log.Infof("mates started successfully")
	svr.Run(context.Background())
	return
}
