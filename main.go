package main

import (
	"flag"
	"fmt"
	"os"
	"present-aur/aur"
	"present-aur/utils"
)

func main() {
	// Parse the sub-commands and their flags respectively.
	search := flag.NewFlagSet("search", flag.ExitOnError)
	inverseSearch := search.Bool("inverse", false, "Show the results inverse to reduce scrolling time")
	
	install := flag.NewFlagSet("install", flag.ExitOnError)
	
	uninstall := flag.NewFlagSet("uninstall", flag.ExitOnError)
	dependencies := uninstall.Bool("dependencies", false, "Uninstall the dependencies alongside the packages you wish to uninstall (identical to using the -Rs flag in pacman)")

	help := flag.NewFlagSet("help", flag.ExitOnError)


	// If there's no arguments, exit the program.
	if len(os.Args) < 1 {
		fmt.Println(utils.Prefix + "You need to use a subcommand")
		os.Exit(0)
	}

	// Check which subcommand is given.
	subcommand := os.Args[1]
	switch subcommand {
	case "search":
		search.Parse(os.Args[2:])
		if len(search.Args()) == 0 {
			fmt.Println(utils.Prefix + "You need to enter a search query")
			os.Exit(0)
		}
		aur.SearchPackages(search.Args()[0], *inverseSearch)
		break
	case "install":
		install.Parse(os.Args[2:])
		// Clear the cache just in case, the cache should not include anything important as it's just clones of package repos.
		utils.ClearCache()
		if len(install.Args()) == 0 {
			fmt.Println(utils.Prefix + "You need to enter a package you wish to install")
			os.Exit(0)
		}
		aur.InstallPackages(install.Args())
		break
	case "uninstall":
		uninstall.Parse(os.Args[2:])
		if len(uninstall.Args()) == 0 {
			fmt.Println(utils.Prefix + "You need to enter a package you wish to uninstall")
			os.Exit(0)
		}
		aur.UninstallPackages(uninstall.Args(), *dependencies)
		break
	case "help":
		// To-do: A dynamic way, not hardcoded method to add all of the subcommands to the help output.
		help.Parse(os.Args[2:])
		fmt.Println(utils.Prefix + "Usage: presaur [install <packages> | search <query> | uninstall <packages> | help]\n\nFor more detailed information and possible flags for every sub-command, you can use the '-help' flag (e.g. 'presaur uninstall -help')")
		break
	default:
		fmt.Println(utils.Prefix + "Invalid usage; run 'presaur help' to learn more.")
		os.Exit(0)
	}
}
