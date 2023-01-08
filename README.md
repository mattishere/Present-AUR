# Present AUR
![screenshot](media/screenshot-transparent.png)

## Note
- **This is not a fully released program and is not guaranteed to be fully stable**. Using this program **may lead to data loss or corruption** as it's not a fully tested/developed project.

## Dependencies
- An Arch Linux system (of course).
- `sudo pacman -S base-devel git go`.

## Installation
- Clone the repository: `git clone https://github.com/mattishere/present-aur`.
- Change your directory into it and build it: `cd present-aur && make install` (You likely have to run make with elevated privileges).

## Uninstallation
- Go into the cloned repository and run `make uninstall` (You likely have to run make with elevated privileges).

## Usage
|Command    |Explanation
|-----------|--------
|`presaur install [packages]`|Install any packages from the AUR.
|`presaur uninstall [packages]`|Uninstall any packages you have installed locally (this command makes use of the `pacman` package manager)
|`presaur search [query]`|Search the AUR with your query to find available packages.
|`presaur help`|Shows you all of the possible sub-commands available.

To check for extra flags that are available for each sub-command, you can use the `--help` flag (e.g. `presaur uninstall --help` will show you that there's a `--dependencies` flag available & its meaning).

## To-Do
- All of the major features of a standard AUR helper.
- A TUI to make searching for packages nicer..?
- A lot of bug fixes.
- A bunch of stuff (there's too much to write down here).

## Known issues
- When installing/uninstalling, you may not get some confirmation text (e.g. Do you wish to proceed [Y/n]?). If you figure out that it's waiting for your confirmation, you can input your decision and it will work as intended.