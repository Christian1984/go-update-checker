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
	PreRelease  bool   `json:"prerelease"`
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

func writeLatestCheckFile(checkData CheckData) error {
	//fmt.Println(checkData)
	//TODO: write file

	return nil
}

func hasLatestVersion(currentVersion string, availableVersion string) bool {
	//TODO
	return true
}

func canCheck(latestCheckTimestamp string, minIntervalDays int) bool {
	//TODO
	return true
}

func printUpdateMessage(checkData CheckData) {
	fmt.Println("A new Update is available...")
	//TODO
}

func IsUpdateAvailable(owner string, repo string, currentVersion string, minDaysInterval int, verbose bool) bool {
	fmt.Println("CheckForUpdate called...")

	latestCheck, fileErr := loadFile(Filename, verbose)
	if fileErr != nil {
		//return false
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

			writeLatestCheckFile(latestCheck)
		}
	}

	if hasLatestVersion(currentVersion, latestCheck.Version) {
		return false
	}

	printUpdateMessage(latestCheck)
	return true
}
