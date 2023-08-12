// This file is mainly for long strings of text, like the default config file & help message.
package main

const helpMsg = `Usage: sigma [options...] <command> [args...]

Commands:
  help                            Show this help message.
  install <plugins...>           Installs plugins.
  uninstall <plugins...>         Removes plugins.
  update [plugins...]            Re-downloads the a plugin's info & install instructions. If no plugins are specified, all plugins are updated.
  upgrade [plugins...]           Pulls git repository & recompile's a plugin. If no plugin is specified, all plugins are upgraded.
  info <plugin>                  Displays information about a specific plugin.
  remove-data <plugins...>       Removes plugin data from .sigma. Use this only if a plugin installation has failed and the uninstall command won't work.
  re-clone                        Removes and re-clones sigma source code in the src directory. Useful if you just changed the branch in the config file, or git is throwing errors when updating.
  list                            Lists all installed plugins.
  list-all                        Lists all plugins in all github repositories.
  version                         Shows version.
  init                            Re-generates all the default config files needed for sigma to function properly. This is ran automatically.
  search <query>                  Searches for plugins. NOTE: Only searches repo's that are on Github.
  sigma-update                 Updated source code & re-compiles sigma.
  setup                           Sets up everything for a functional installation of sigma.
  re-compile <plugins...>        Re-compiles plugins. Useful if the info for a plugins are present but the binaries are not.
  env-add                         Adds sigma to several environment variables. This is useful if plugins installed with sigma are not being found.
  repo                            Manages the plugin sources file.
    repo add <url>                Adds a repository to the plugin sources file.
    repo remove <url>             Removes a repository to the plugin sources file.
    repo list                     Lists all repositories in the plugin sources file.]

Developer & Debugging Commands:
  fetch                           Fetches OS & sigma info.
  raw-version                     Shows version without any formatting (For use in scripts).
  github-gen <username> <repo>    Generates a plugin info file for a github repo.
  help2man                        Generates a manpage in the current working directory. NOTE: This only works if your current working directory is in the sigma source directory.

Options:
  -p, --purge                     Removes a plugin's configuration files as well as the plugin itself.
  -d, --debug                     Displays variable & debugging information.
  -y, --assumeyes                 Assumes yes to all prompts. (Use with caution!)
  -f, --force                     Bypasses all checks before preforming an action. Use will almost certainly lead to an error.
  -r, --ignoreroot                Continues even if user is root.
  -n, --nodeps                    Continues even if there are unmet dependencies.

Examples:
  sigma install my-pkg

  sigma uninstall other-pkg

  sigma upgrade third-pkg
`

const defaultConf = `# sigma config file

# Locations for files
# The home directory is prepended automatically for each value, do not include it.
[paths]

# Default: .local
prefix = ".local"

# Configures updating and auto-updating
[updating]

# Default: testing
# The branch to pull from. https://github.com/talwat/sigma#branches for more information about the branches.
branch = "testing"

# Default: true
auto_update = true

# Configures progressbar, purely visual.
[progressbar]

saucer = "[cyan]=[reset]"
saucer_head = "[cyan]>[reset]"
alt_saucer_head = "[cyan]>[reset]"
saucer_padding = " "
bar_start = "("
bar_end = ")"

# Configures github authentication, this is used for querying plugins.
[github]

# Github username
username = ""

# Github access token, this should not expire and have no permissions granted for security purposes.
# Information about getting a token is available here. https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
token = ""
`

const defaultSources = `# Please only add sources you trust.
# This file contains the links to the pkg.json files. If you mess up, you can simply run 'sigma init' to reset it.
# You can also edit this file with 'sigma repo'.

https://raw.githubusercontent.com/talwat/sigma/main/plugins/
# https://raw.githubusercontent.com/talwat/sigma/main/plugins/bin/
# https://raw.githubusercontent.com/talwat/sigma/main/plugins/linux-only/`
