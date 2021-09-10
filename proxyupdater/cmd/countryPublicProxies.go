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
	"fmt"
)

// countryPublicProxiesCmd represents the countryPublicProxies command
var countryPublicProxiesCmd = &cobra.Command{
	Use:   "country-public-proxies",
	Short: "Country Public Proxies Json format",
	Long: `Public proxies are scraped and updated every 10 min.  We are using several paid proxy scraper programs worth hundreds of dollars and running them on a server 24/7. These are good for testing.
For posting on BHW
`,
	Run: func(cmd *cobra.Command, args []string) {
		if country != ""{
			scripts.PublicCountryProxies(country)
		} else {
			fmt.Println("must provide two letter country code")
		}
	},
}

func init() {
	rootCmd.AddCommand(countryPublicProxiesCmd)
	countryPublicProxiesCmd.Flags().StringVarP(&country, "country", "c", "", "two letter country code for proxies")
}
