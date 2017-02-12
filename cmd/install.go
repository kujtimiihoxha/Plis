// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kujtimiihoxha/plis/fs"
	"github.com/kujtimiihoxha/plis/logger"
	"github.com/kujtimiihoxha/plis/config"
	"encoding/json"
	"github.com/kujtimiihoxha/plis/api"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all the plis dependencies",
	Long: `Install all the plis dependencies`,
	Run: func(cmd *cobra.Command, args []string) {
		installDependencies()
	},
}

func installDependencies() {
	fsApi := api.NewFsAPI(fs.GetCurrentFs())
	data,err:=fsApi.ReadFile("plis.json")
	if err != nil {
		logger.GetLogger().Fatal("Could not read pplis configuration file : ", err)
	}
	plisConfig := config.PlisConfig{}
	err = json.Unmarshal([]byte(data),&plisConfig)
	if err != nil {
		logger.GetLogger().Fatal("Could not decode plis configurations : ", err)
	}
	for _,v := range plisConfig.Dependencies{
		getGenerator(v.Repository,v.Branch)
	}
}

func init() {
	RootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}