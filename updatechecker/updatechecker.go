package updatechecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const Filename string = "latestcheck.json"
const DateFormat string = "2006-01-02T15-04-05"

type CheckData struct {
	Timestamp   string `json:"timestamp"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GithubApiResponseData struct {
	Version     string `json:"tag_name"`
	Name        string `json:"name"`
	Description string `json:"body"`
}

func processError(err error, verbose bool) {
	if !verbose {
		return
	}

	fmt.Println("ERROR: " + err.Error())
}

func requestLatest(owner string, repo string, verbose bool) (GithubApiResponseData, error) {
	// call https://api.github.com/repos/{owner}/{repo}/releases/latest
	var apiResponse GithubApiResponseData

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, requestErr := client.Get("https://api.github.com/repos/" + owner + "/" + repo + "/releases/latest")

	if requestErr != nil {
		processError(requestErr, verbose)
		return apiResponse, requestErr
	}

	body, ioUtilErr := ioutil.ReadAll(resp.Body)
	if ioUtilErr != nil {
		processError(ioUtilErr, verbose)
		return apiResponse, ioUtilErr
	}

	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		processError(jsonErr, verbose)
		return apiResponse, jsonErr
	}

	if verbose {
		fmt.Println("GitHub API Response:")
		fmt.Println(apiResponse)
	}

	return apiResponse, nil
}

func loadFile(filename string, verbose bool) (CheckData, error) {
	var latestCheck CheckData

	file, err := os.Open(filename)
	if err != nil {
		processError(err, verbose)
		return latestCheck, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	jsonErr := json.Unmarshal(byteValue, &latestCheck)
	if jsonErr != nil {
		processError(err, verbose)

		epoch := time.Time{}
		latestCheck.Timestamp = epoch.Format(DateFormat)
		return latestCheck, jsonErr
	}

	return latestCheck, nil
}

func writeLatestCheckFile(checkData CheckData, verbose bool) error {
	file, jsonErr := json.Marshal(checkData)
	if jsonErr != nil {
		processError(jsonErr, verbose)
		return jsonErr
	}

	fileErr := ioutil.WriteFile(Filename, file, 0644)
	if fileErr != nil {
		processError(fileErr, verbose)
		return fileErr
	}

	return nil
}

func hasLatestVersion(currentVersion string, availableVersion string) bool {
	//TODO
	return false
}

func canCheck(latestCheckTimestamp string, minIntervalDays int) bool {
	//TODO
	return true
}

func printUpdateMessage(checkData CheckData, owner string, repo string) {
	fmt.Println("==========================================")
	fmt.Println("=== INFO: A new update is available... ===")
	fmt.Println()
	fmt.Println("Version: " + checkData.Version)
	fmt.Println()
	fmt.Println("Title: " + checkData.Name)
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println(checkData.Description)
	fmt.Println()
	fmt.Println("Download the latest version here:")
	fmt.Println("https://github.com/" + owner + "/" + repo + "/releases")
	fmt.Println("==========================================")
	fmt.Println()
}

func IsUpdateAvailable(owner string, repo string, currentVersion string, minDaysInterval int, verbose bool) bool {
	latestCheck, fileErr := loadFile(Filename, verbose)
	if fileErr != nil {
		//return false
	}

	if verbose {
		fmt.Println("Returned from loadFile():")
		fmt.Println(latestCheck)
	}

	if canCheck(latestCheck.Timestamp, minDaysInterval) {
		apiResponse, apiErr := requestLatest(owner, repo, verbose)
		if apiErr == nil {
			now := time.Now()
			snow := now.Format(DateFormat)

			latestCheck.Timestamp = snow
			latestCheck.Version = apiResponse.Version
			latestCheck.Name = apiResponse.Name
			latestCheck.Description = apiResponse.Description

			writeLatestCheckFile(latestCheck, verbose)
		}
	} else if verbose {
		fmt.Println("Didn't request GitHub API because interval isn't over yet...")
	}

	if hasLatestVersion(currentVersion, latestCheck.Version) {
		return false
	}

	printUpdateMessage(latestCheck, owner, repo)
	return true
}
