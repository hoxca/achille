# achille

Achille decode information from the Talon6 driver log file
Achille will verify that the last log line contain information is up to date before processing it.
Achille Talon print and report Talon6 RoR status from a command line and be able to be used in imaging sequencers to correctly confirm different roof status.

At Multiscale Detection, we successfully use achille to manage emergency situation in our spain observatory.


# Usage

As far as we use cobra to code the CLI all command are self documented ;)
there are 2 major command to get informations from roof and switch

Here is the help for the 'achille roof' command:

```
hugh⨕shula:achille|● ./bin/achille roof --help
Command return the different roll of roof status

Usage:
  achille roof [flags]
  achille roof [command]

Available Commands:
  isClose     achille command to confirm that talon roof is close
  isFullyOpen achille command to confirm that talon roof is fully open
  isOpen      achille command to confirm that talon roof is open

Flags:
  -h, --help   help for roof

Global Flags:
      --config string         config file (default is .achille.yaml in program location directory)
      --delay int             delay in seconds before considering last log line too old to process (default 30)
      --location string       location timezone (default is "Europe/Paris")
      --talonLogFile string   talon log file location (default is Talon6_ROR.log in program location directory)
      --verbosity string      Log level (debug, info, warn, error) (default "warn")

Use "achille roof [command] --help" for more information about a command.
```

Here is the help for the 'achille switch' command:

```
hugh⨕shula:achille|● ./bin/achille switch --help
Command return the different talon switch status

Usage:
  achille switch [flags]
  achille switch [command]

Available Commands:
  mountAtPark   Switch confirm that mount is parked
  roofIsClose   Switch confirm that the roof is close
  roofIsOpen    Switch confirm that the roof is open
  systemOnPower switch confirm that system on power
  weatherIsSafe switch confirm that weather is safe

Flags:
  -h, --help   help for switch

Global Flags:
      --config string         config file (default is .achille.yaml in program location directory)
      --delay int             delay in seconds before considering last log line too old to process (default 30)
      --location string       location timezone (default is "Europe/Paris")
      --talonLogFile string   talon log file location (default is Talon6_ROR.log in program location directory)
      --verbosity string      Log level (debug, info, warn, error) (default "warn")

Use "achille switch [command] --help" for more information about a command.
```
