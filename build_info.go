//go:generate bash -c "mkdir -p tmp"
package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:generate bash -c "date --iso-8601=seconds > tmp/build_time"
//go:embed tmp/build_time
var buildTime string

// Local time of build
var BuildTime = strings.TrimSpace(buildTime)

//go:generate bash -c "git rev-parse HEAD > tmp/commit_sha"
//go:embed tmp/commit_sha
var commitSHA string

// Commit SHA
var CommitSHA = strings.TrimRight(commitSHA, "\n")

//go:generate bash -c "git rev-parse --short=8 HEAD > tmp/short_commit_sha"
//go:embed tmp/short_commit_sha
var shortCommitSHA string

// Short Commit SHA
var ShortCommitSHA = strings.TrimRight(shortCommitSHA, "\n")

//go:generate bash -c "git rev-parse --abbrev-ref HEAD > tmp/branch"
//go:embed tmp/branch
var branch string

// Build branch (if applicable)
var Branch = strings.TrimRight(branch, "\n")

//go:generate bash -c "hostname > tmp/hostname"
//go:embed tmp/hostname
var buildHostname string

// Hostname on which build was done
var BuildHostname = strings.TrimRight(buildHostname, "\n")

//go:generate bash -c "echo $USER > tmp/username"
//go:embed tmp/username
var buildUsername string

// Username who did the build
var BuildUsername = strings.TrimRight(buildUsername, "\n")

//go:generate bash -c "git tag --points-at $(git rev-parse HEAD) > tmp/commit_tag"
//go:embed tmp/commit_tag
var commitTag string

// Tag of build (if applicable)
var CommitTag = strings.TrimRight(commitTag, "\n")

//go:generate bash -c "git remote -v | head -1 | cut -f2 | cut -d' ' -f1 > tmp/repository"
//go:embed tmp/repository
var repository string

// Repository remote
var Repository = strings.TrimSpace(repository)

// BuildInfo holds the name and value of a build time var
type BuildInfo struct {
	Name  string
	Value string
}

// GetBuildInfo returns a list of BuildInfo containing build time vars
func GetBuildInfo() []BuildInfo {
	return []BuildInfo{
		{Name: "Repository", Value: Repository},
		{Name: "CommitSHA", Value: CommitSHA},
		{Name: "ShortCommitSHA", Value: ShortCommitSHA},
		{Name: "Branch", Value: Branch},
		{Name: "CommitTag", Value: CommitTag},
		{Name: "BuildHostname", Value: BuildHostname},
		{Name: "BuildTime", Value: BuildTime},
		{Name: "BuildUsername", Value: BuildUsername},
	}
}

// returns the length of the longest heading and value
func longestHeadingAndValue(info []BuildInfo) (int, int) {
	var heading, value int
	for _, item := range info {
		if len(item.Name) > heading {
			heading = len(item.Name)
		}
		if len(item.Value) > value {
			value = len(item.Value)
		}
	}
	return heading, value
}

// GetBuildInfoFormatted returns a list of strings for either printing or logging
func GetBuildInfoFormatted() []string {
	buildInfo := GetBuildInfo()

	// build format strings based on longest heading and value
	headingLength, valueLength := longestHeadingAndValue(buildInfo)
	format := "| %-" + fmt.Sprintf("%d", headingLength+1) + "s %-" + fmt.Sprintf("%d", valueLength) + "s |"

	// top and bottom borders need an additional four chars to account for spaces
	lineFormat := "+%" + fmt.Sprintf("%d", headingLength+valueLength+4) + "s+"
	longLine := strings.Replace(fmt.Sprintf(lineFormat, ""), " ", "-", -1)
	
	// build formatted strings
	content := []string{longLine}
	for _, item := range buildInfo {
		content = append(content, fmt.Sprintf(format, item.Name+":", item.Value))
	}
	return append(content, longLine)
}

// PrintBuildInfo prints out the build info
func PrintBuildInfo() {
	for _, line := range GetBuildInfoFormatted() {
		fmt.Println(line)
	}
}
