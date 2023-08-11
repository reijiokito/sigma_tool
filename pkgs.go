package main

var DestFolder string

func installPkgs(pkgNames []string) {
	displayPkgs(pkgNames, "install")

	for _, pkgName := range pkgNames {
		chapLog("=>", "", "Installing %s", pkgName)
		chapLog("==>", "", "Preparing for installation of %s", pkgName)

		initPlugin()

		chapLog("==>", textCol.Green, "Successfully installed %s", pkgName)
	}

	chapLog("=>", textCol.Green, "Success")
	log(0, "Installed all selected packages successfully.")
}

func uninstallPkgs(pkgNames []string) {

}
