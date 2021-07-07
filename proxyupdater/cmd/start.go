// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"github.com/proxyips/proxyupdater/scripts"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Proxy Updater API service",
	Run: func(cmd *cobra.Command, args []string) {
		scripts.StartApiService(ipAddress, port)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&ipAddress, "ipaddress", "i", "0.0.0.0", "IP address to listen on")
	startCmd.Flags().StringVarP(&port, "port", "p", "8009", "Port to listen on")

}

var (
	ipAddress string
	port      string
)
