package main

import _ "embed"

type Build struct {
	Commit string `json:"commit"`
	Branch string `json:"branch"`
	CiPlat string `json:"ci_plat"`
}

//go:embed .commit
var commitId string

//go:embed .branch
var branch string

//go:embed .ci_plat
var ciPlat string

func getBuildInfo() Build {
	return Build{
		commitId,
		branch,
		ciPlat,
	}
}
