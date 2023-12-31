// This file is mainly for long strings of text, like the default config file & help message.
package main

const helpMsg = `Usage: sigma [options...] <command> [args...]

Commands:
  help                            Show this help message.
  install <plugins...>           Installs plugins.
  uninstall <plugins...>         Removes plugins.
  upgrade [plugins...]           Pulls git repository & recompile's a plugin. If no plugin is specified, all plugins are upgraded.

Options:
  -f, --force                     Bypasses all checks before preforming an action. Use will almost certainly lead to an error.

Examples:
  sigma install my-plg

  sigma uninstall other-plg

  sigma upgrade third-plg
`
