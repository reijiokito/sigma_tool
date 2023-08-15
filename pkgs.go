package main

func installPlugins(plgNames []string) {
	displayPlugins(plgNames, "install")

	for _, plgName := range plgNames {
		chapLog("=>", "", "Installing %s", plgName)
		chapLog("==>", "", "Preparing for installation of %s", plgName)

		initPlg(plgName)

		chapLog("==>", textCol.Green, "Successfully installed %s", plgName)
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Installed all selected plugins successfully.")
}

func uninstallPlugins(plgNames []string) {
	displayPlugins(plgNames, "uninstall")

	for _, plgName := range plgNames {
		chapLog("=>", "", "Uninstalling %s", plgName)
		removePlg(plgName)

		chapLog("==>", textCol.Green, "Successfully uninstalled %s", plgName)
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Uninstalled all selected plugins successfully.")
}

func publishPlugin(plgName string, pluginDir string, version string) {
	chapLog("=>", "", "Publishing plugin %s", plgName)

	publishPlg(plgName, pluginDir, version)

	chapLog("=>", textCol.Green, "Success")
	log(0, "Uninstalled all selected plugins successfully.")
}
