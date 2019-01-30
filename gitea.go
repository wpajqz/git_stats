package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Repo repos info from gitea
type Repo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
}

// GetOrgRepos List an organization's repos
func GetOrgRepos(org string) ([]Repo, error) {
	req, err := http.NewRequest(http.MethodGet, "http://git.links123.net/api/v1/orgs/"+org+"/repos", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+c.GiteaToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rab, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	repos := make([]Repo, 0)
	json.Unmarshal(rab, &repos)

	return repos, nil
}

// CloneRepo clone repo when exists update
func CloneRepo(name, url string) error {
	dir := strings.Join([]string{c.GitDir, name}, string(os.PathSeparator))
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		logrus.Info(strings.Join([]string{"git", "clone", url, dir}, " "))
		if execCommand("git", []string{"clone", url, dir}) {
			return nil
		}

		return errors.New("exec git command error")
	}

	logrus.Info(strings.Join([]string{"git", "fetch", dir}, " "))
	if execCommand("git", []string{"fetch", url}) {
		return nil
	}

	return errors.New("exec git command error")
}

// AnalyseRepo git_stats generate -p git/pkg -o git_stats/pkg
func AnalyseRepo(name string) bool {
	gitDir := strings.Join([]string{c.GitDir, name}, string(os.PathSeparator))
	statsDir := strings.Join([]string{c.StatsDir, name}, string(os.PathSeparator))

	return execCommand("git_stats", []string{"generate", "-p", gitDir, "-o", statsDir})
}
