package main

type Bin struct {
	Installed []string
	InSource  []string `json:"in_source"`
}

type Commands struct {
	Install   []string
	Uninstall []string
	Update    []string
}

type OSCommands struct {
	All    *Commands
	Linux  *Commands
	Darwin *Commands
}

type Deps struct {
	All    []string
	Linux  []string
	Darwin []string
}

type Plugin struct {
}
