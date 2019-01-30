package main

import (
	"github.com/sirupsen/logrus"
)

// RunGitStats run git status command
func RunGitStats() {
	orgs := []string{"links123.com", "cloudteam"}
	for _, org := range orgs {
		repos, err := GetOrgRepos(org)
		if err != nil {
			continue
		}

		for _, repo := range repos {
			go func(repoName, cloneURL string) {
				err := CloneRepo(repoName, cloneURL)
				if err != nil {
					logrus.Info("clone repo " + repoName + "is faild")
					return
				}

				logrus.Info("clone repo " + repoName + "is successful")

				if !AnalyseRepo(repoName) {
					logrus.Info("analyse repo " + repoName + "is faild")
					return
				}

				logrus.Info("analyse repo " + repoName + "is successful")
			}(repo.Name, repo.CloneURL)
		}
	}
}
