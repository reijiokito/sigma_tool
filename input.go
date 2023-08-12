package main

import (
	"os"
	"strings"
)

const version = "0.0.1"

var (
	purge      bool = false
	debug      bool = false
	assumeYes  bool = false
	force      bool = false
	noDeps     bool = false
	ignoreRoot bool = false
)

var optionToOthers, optionToOther bool = false, false

func checkFlag(flag string) {
	switch flag {
	case "-help", "--help":
		rawLog(helpMsg)

		os.Exit(0)
	case "-version", "--version":
		rawLog(version)

		os.Exit(0)
	default:
		errorLogRaw("Flag %s not found", bolden(flag))
		os.Exit(1)
	}
}

func checkCommand(other string, others []string, index int, args []string) {
	checkForOptions := func(errSpecify string, commandPartsCount int) {
		if len(others[index+commandPartsCount:]) < 1 {
			errorLogRaw("No %s specified", errSpecify)

			os.Exit(1)
		}
	}

	switch other {
	case "install":
		checkForOptions("plugin names", 1)

		optionToOthers = true

		installPkgs(others[index+1:])

	case "uninstall":
		checkForOptions("plugin names", 1)

		optionToOthers = true

		uninstallPkgs(others[index+1:])

	case "upgrade":

	case "usname":
	case "help":
		rawLog(helpMsg)
	case "dest":
		setDest(others[index+1])
	default:
		errorLogRaw("Command %s not found", bolden(other))
		os.Exit(1)
	}
}

func parseInput() {
	args := os.Args[1:]

	var (
		flags  []string
		others []string
	)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			others = append(others, arg)
		}
	}

	for _, flag := range flags {
		checkFlag(flag)
	}

	for i, other := range others {
		if !optionToOthers && !optionToOther {
			checkCommand(other, others, i, args)
		}

		optionToOther = true
	}

	if len(others) < 1 {
		log(1, "Indiepkg Version %s, run %s for usage.", bolden(version), bolden("indiepkg help"))
	}
}
