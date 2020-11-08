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

// Package cmd for achille 'switch' command
package cmd

import (
	"achille/roof"
	"fmt"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:       "switch",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"status", "mountAtPark", "weatherIsSafe", "systemOnPower", "roofIsOpen", "roofIsClose"},
	Short:     "Command return the different talon switch status",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Achille Talon switch status\n\n")
		talonSwitch := new(roof.Switch)
		talonSwitch = talonSwitch.Set(splitedlog[5])
		talonSwitch.PrintSwitchesStatus()
		if talonSwitch.IsSystemOnPower() {
			fmt.Println("Sensor confirm that system on power")
		}
		if talonSwitch.IsWeatherSafe() {
			fmt.Println("Sensor confirm that weather is safe")
		}

	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
