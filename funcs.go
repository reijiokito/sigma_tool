package main

func displayPkgs(pkgNames []string, action string) {
	log(1, "Are you sure you would like to %s the following packages:", bolden(action))

	for _, pkgToDisplay := range pkgNames {
		indent(pkgToDisplay)
	}

	confirm("y", "(y/n)")
}

func fullInit() {
	chapLog("=>", "", "Initializing")
	//Step1
}
