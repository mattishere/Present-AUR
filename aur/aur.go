package aur

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"present-aur/utils"
	"strconv"
	"time"
)

var baseUrl = "https://aur.archlinux.org/rpc/v5/"
var client = &http.Client{Timeout: 15 * time.Second}

type Package struct {
	Name        string `json:"Name"`
	Description string `json:"description"`
	Maintainer  string `json:"maintainer"`
	Version     string `json:"Version"`
}

type SearchResults struct {
	ResultsCount int       `json:"resultcount"`
	Results      []Package `json:"results"`
}

func AurSearch(arg, endpoint string) SearchResults {

	res, err := client.Get(baseUrl + endpoint + "/" + arg)

	if err != nil {
		fmt.Println(utils.Prefix+"There was an error getting search data from the AUR:\n", err)
		os.Exit(0)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(utils.Prefix+"There was an error parsing the JSON from the AUR:\n", err)
		os.Exit(0)
	}

	var results SearchResults
	if err := json.Unmarshal(data, &results); err != nil {
		fmt.Println(utils.Prefix+"There was an error parsing JSON data from the AUR API:\n", err)
		os.Exit(0)
	}
	return results
}

func InstallPackages(packages []string) {
	for i := 0; i < len(packages); i++ {
		install(packages[i])
	}
}

func UninstallPackages(packages []string, dependencies bool) {
	var flag string
	if dependencies == true {
		flag = "-Rs"
	} else {
		flag = "-R"
	}

	arguments := []string{"sudo", "pacman", flag}
	arguments = append(arguments, packages...)

	pacman := exec.Command(arguments[0], arguments[1:]...)
	pacman.Stdin = os.Stdin
	pacman.Stdout = os.Stdout
	err := pacman.Run()
	if err != nil {
		fmt.Println(utils.Prefix+"Error running the uninstall command (you likely don't have the package installed!):\n", err)
	}
}

func SearchPackages(pkg string, inverse bool) {
	results := AurSearch(pkg, "search")
	if results.ResultsCount == 0 {
		fmt.Println(utils.Prefix + "Unable to find any packages in the AUR based on your search query '" + pkg + "'")
		os.Exit(0)
	}

	count := strconv.Itoa(results.ResultsCount)

	resultsText := "\n------------------------------------\n"+utils.Green + count +  utils.Reset+" results found in the AUR for the search query '"+pkg+utils.Reset+"'\n"

	if inverse == true {
		fmt.Println(resultsText)
		for _, result := range results.Results {
			fmt.Println(packageInfo(result))
		}
	} else {
		for i := results.ResultsCount - 1; i >= 0; i-- {
			result := results.Results[i]
			fmt.Println(packageInfo(result))
		}
		fmt.Println(resultsText)
	}
}

// private helper functions

func packageInfo(result Package) string {
	return utils.Blue + result.Name + utils.Yellow + "\n- description: " + utils.Reset + result.Description + utils.Yellow + "\n- version: " + utils.Reset + result.Version + "\n"	
}

func install(pkg string) {
	results := AurSearch(pkg, "info")
	if results.ResultsCount == 0 {
		fmt.Println(utils.Prefix + "Could not find package '" + pkg + "' in the AUR")
		os.Exit(0)
	}
	cache, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(utils.Prefix+"Error getting user's cache directory:\n", err)
		os.Exit(0)
	}
	fmt.Println(utils.Prefix + "Found cache directory at '" + cache + "'")

	path := cache + "/present-aur/" + pkg
	mkdirErr := os.MkdirAll(path, os.ModePerm)
	if mkdirErr != nil {
		utils.ClearCache()
		fmt.Println(utils.Prefix+"Error while making a directory at path '"+path+"':\n", mkdirErr)
		os.Exit(0)
	}
	fmt.Println(utils.Prefix + "Created temporary directory at '" + path + "'")

	url := "https://aur.archlinux.org/" + pkg + ".git"
	cloneCmd := exec.Command("git", "clone", url, path)
	_, cloneErr := cloneCmd.Output()
	if cloneErr != nil {
		utils.ClearCache()
		fmt.Println(utils.Prefix+"Unable to clone package '"+pkg+"' in '"+path+"':\n", cloneErr)
		os.Exit(0)
	}
	fmt.Println(utils.Prefix + "Cloned '" + pkg + "' into '" + path + "'")

	makeCmd := exec.Command("makepkg", "-si")
	makeCmd.Dir = path
	makeCmd.Stdin = os.Stdin
	makeCmd.Stdout = os.Stdout

	makeErr := makeCmd.Run()
	if makeErr != nil {
		utils.ClearCache()
		fmt.Println(utils.Prefix+"Unable to build the package '"+pkg+"':\n", err)
		os.Exit(0)
	}
	fmt.Println(utils.Prefix + "You have installed '" + pkg + "'")
	utils.ClearCache()
}
