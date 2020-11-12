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

// Package cmd for cobra root command
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	Log "github.com/apatters/go-conlog"

	"achille/roof"
	"achille/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var talonLogFile string
var myLocation string
var delay = 30
var verbosity = "warn"
var splitedlog []string
var cfgFileNotFound = false

// rootCmd represents the base command when 'achille' is called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "achille",
	Short: "Achille decode informations from the Talon6 driver log file",
	Long: `Achille decode information from the Talon6 driver log file
Achille will verify that the last log line contain information
is up to date before processing it.

Achille Talon print and report Talon6 RoR status from a command line
and be able to be used in imaging sequencers to correctly confirm different roof status.

At Multiscale Detection, we successfully use achille to manage emergency situation in our spain observatory.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Achille Talon")
		fmt.Println("\nRoof status:")
		talonRoof := new(roof.Status)
		talonRoof = talonRoof.Set(splitedlog[3], splitedlog[2])
		talonRoof.PrintStatus()

		fmt.Println("\nSwitches status:")
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

// Execute is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is achille.yaml in program location or conf directory of program location)")
	rootCmd.PersistentFlags().StringVar(&talonLogFile, "talonLogFile", "", "talon log file location (default is Talon6_ROR.log in program location directory)")
	rootCmd.PersistentFlags().StringVar(&myLocation, "location", "", "location timezone (default is \"Europe/Paris\")")
	rootCmd.PersistentFlags().IntVar(&delay, "delay", 30, "delay in seconds before considering last log line too old to process")
	rootCmd.PersistentFlags().StringVar(&verbosity, "verbosity", "warn", "Log level (debug, info, warn, error)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	setUpLogs()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Switch to default program path
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		confdir := fmt.Sprintf("%s/conf", dir)
		// Search yaml config file in program path with name "achille.yaml".
		viper.AddConfigPath(dir)
		viper.AddConfigPath(confdir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("achille")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cfgFileNotFound = true
			Log.Debug("Config file not found")
		} else {
			Log.Debug("Something look strange")
			Log.Debugf("error: %v\n", err)
		}
	} else {
		Log.Debug("Using config file:", viper.ConfigFileUsed())
	}

	manageDefault()
	parseTalonLastLogline()

}

func setUpLogs() {

	formatter := Log.NewStdFormatter()
	formatter.Options.LogLevelFmt = Log.LogLevelFormatLongTitle
	Log.SetFormatter(formatter)
	switch verbosity {
	case "debug":
		Log.SetLevel(Log.DebugLevel)
	case "info":
		Log.SetLevel(Log.InfoLevel)
	case "warn":
		Log.SetLevel(Log.WarnLevel)
	case "error":
		Log.SetLevel(Log.ErrorLevel)
	default:
		Log.SetLevel(Log.WarnLevel)
	}

}

func manageDefault() {

	var talonLogFileViper = viper.GetString("talonlogfile")
	if talonLogFile == "" && talonLogFileViper == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		talonLogFile = fmt.Sprintf("%s/%s", dir, "Talon6_ROR.log")
	}
	if talonLogFile == "" {
		talonLogFile = talonLogFileViper
	}

	if myLocation == "" && viper.GetString("location") == "" {
		myLocation = "Europe/Paris"
	}
	if myLocation == "" {
		myLocation = viper.GetString("location")
	}

	if delay == 30 && viper.GetInt("delay") != 0 {
		delay = viper.GetInt("delay")
	}

}

func parseTalonLastLogline() {

	if _, err := os.Stat(talonLogFile); err != nil {

		Log.Errorf("Talong log file: %s not found !\n", talonLogFile)
		if cfgFileNotFound {
			Log.Warning("You must provide talonLogFile via configuration file")
			Log.Warning("or provide a valid talon log location to achille using command line flags !")
		} else {
			Log.Warning("Talon log defined path is not valid !")
		}
		Log.Error("Critical: Exiting, cannot find talon log file")
		os.Exit(1)
	}

	utils.MyLocation = myLocation
	utils.Delay = delay

	logline := utils.GetLastLineWithSeek(talonLogFile)
	Log.Debugf("%s", logline)
	splitedlog = utils.ParseLogline(logline)

	fresh, seconds := utils.IsFreshlog(splitedlog[1])
	if !fresh {
		Log.Debugf("Last log date: %v\n", splitedlog[1])
		Log.Warningf("Log is from %d seconds back.", seconds)
		Log.Error("Talon status may not be accurate...\nExiting!!!")
		os.Exit(1)
	}
	Log.Infof("log is fresh %d seconds back.\n", seconds)

}
