/*
Copyright © 2021 TRH Hosting <software@trhhosting.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/proxyips/proxyupdater/scripts"
	"fmt"
)

// bandwidthCmd represents the bandwidth command
var bandwidthCmd = &cobra.Command{
	Use:   "bandwidth",
	Short: `bandwidth log`,
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := scripts.BandwidthLog(host, path)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bandwidthCmd)
	bandwidthCmd.Flags().StringVar(&path, "path", "data", "path to store the log")
	bandwidthCmd.Flags().StringVar(&host, "host", "", "host to gobetween management api")
}

var path, host string
