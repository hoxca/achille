/*
Copyright © 2020 Hugues Obolonsky <hugh@atosc.org>

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

// Package cmd for 'switch roofIsOpen' command
package cmd

import (
	"achille/roof"
	"fmt"

	"github.com/spf13/cobra"
)

// roofIsOpenCmd represents the roofIsOpen command
var roofIsOpenCmd = &cobra.Command{
	Use:   "roofIsOpen",
	Short: "Switch confirm that the roof is open",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		talonSwitch := new(roof.Switch)
		talonSwitch = talonSwitch.Set(splitedlog[5])

		if talonSwitch.IsRoofOpen() {
			fmt.Printf("OK, switch confirm that roof is open")
		} else {
			fmt.Printf("Nope!")
		}

	},
}

func init() {
	switchCmd.AddCommand(roofIsOpenCmd)
}
