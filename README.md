# Sigma Tool

## _Version 0.0.1_

### Introduction
**Sigma Tool** is a tool which play a role as a package manager that use OS-level. It supports **Sigma Plugin Manager** to set up plugins in client machine.
It pulls plugin(zip file) from **Sigma Plugin Registry**, then delivers and sets up the plugin called sigma-plugin in machine.

### Feature
1. **Help**
    - _Description_: show help
    - _Command_ `sigma --help`
2. **Set destination**
    - _Description_: set root destination's plugins
    - _Command_ `sigma install plg_name ...`
3. **Install Plugin**
    - _Description_: install plugins
    - _Command_ `sigma install plg_name ...`    
4. **Uninstall Plugin**
    - _Description_: uninstall plugins
    - _Command_ `sigma uninstall plg_name ...`
5. **Upgrade**
    - _Description_: use to upgrade sigma tool
    - _Command_ `sigma upgrade`
6. **Publish**
    - _Description_: use to publish plugin from client
    - _Command_ `sigma publish plg_name link_plg@version`
