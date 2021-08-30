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
	"github.com/spf13/cobra"
	"github.com/proxyips/proxyupdater/scripts"
)

// publicproxyjsonCmd represents the publicproxyjson command
var publicproxyjsonCmd = &cobra.Command{
	Use:   "public=proxy-json",
	Short: "Public Proxies Json format",
	Long: `Public proxies are scraped and updated every 10 min.  We are using several paid proxy scraper programs worth hundreds of dollars and running them on a server 24/7. These are good for testing.

For high-quality Mobile, Residential, and Datacenter proxies check out https://proxyips.net and use discount code: build..`,
	Run: func(cmd *cobra.Command, args []string) {
		scripts.PublicProxies()
	},
}

func init() {
	rootCmd.AddCommand(publicproxyjsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publicproxyjsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publicproxyjsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
