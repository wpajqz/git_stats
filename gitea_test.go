package main

import (
	"testing"
)

func TestGetOrgsRepos(t *testing.T) {
	d, err := GetOrgRepos("cloudteam")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(d)
}

func TestCloneRepo(t *testing.T) {
	err := CloneRepo("pkg", "git@git.links123.net:links123.com/pkg.git")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAnalyseRepo(t *testing.T) {
	r := AnalyseRepo("pkg")
	t.Log(r)
}
