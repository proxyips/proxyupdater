/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/proxyips/proxyupdater/scripts"
)

// publicproxyCmd represents the publicproxy command
var publicproxyCmd = &cobra.Command{
	Use:   "publicproxy",
	Short: "Updates Public Proxy List",
	Long: scripts.Countries,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Public proxies are scraped and updated every 10 min.  We are using several paid proxy scraper programs worth hundreds of dollars and running them on a server 24/7. These are good for testing.

For high-quality Mobile, Residential, and Datacenter proxies check out https://proxyips.net and use discount code: build. `)

		scripts.UpdateGobetween(filename, country)

		fmt.Println("List provided by ProxyIps.net")

	},
}

func init() {
	rootCmd.AddCommand(publicproxyCmd)

	publicproxyCmd.Flags().StringVarP(&filename, "filename", "f","/opt/gobetween/gobetween.json", "location of the configuration gobetween")
	publicproxyCmd.Flags().StringVarP(&country, "country", "c","", "")
}

var filename, country string