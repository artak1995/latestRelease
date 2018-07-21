package main

import (
	"context"
	"fmt"
	"os"
	"bufio"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"

)

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version

	semver.Sort(releases)
	// fmt.Printf("sorted Releases: %s\n", releases)
	// Sort all of the releases for easy comparison
	tempLatestPatch := minVersion
	// dummy var for comparing highest patch of the same release
	for i, release := range releases {
		if minVersion.LessThan(*releases[i]){

			if release.Major == tempLatestPatch.Major && release.Minor == tempLatestPatch.Minor && release.Patch >= tempLatestPatch.Patch {						
				if len(versionSlice) > 0 {
					versionSlice = versionSlice[:len(versionSlice)-1]
				}
				versionSlice = append(versionSlice, release)
				continue

			}
			// Replace if there is a higher patch of the same release
		versionSlice = append(versionSlice, release)
		tempLatestPatch = release
		}
		

	}
 	for i := len(versionSlice)/2-1; i >= 0; i-- {
	opp := len(versionSlice)-1-i
	versionSlice[i], versionSlice[opp] = versionSlice[opp], versionSlice[i]
	}
	// Reverse the versionSlice in order to display it in descending order
	return versionSlice
}

func main() {
	filepath := os.Args[1:]
	inFile, _ := os.Open(filepath[0])
  	defer inFile.Close()
  	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
  
	// Loop through every repo listed in the txt file
	for scanner.Scan() {
		if (scanner.Text() == "repository,min_version") {
			continue
		}	
		repoPath1 := strings.Split(scanner.Text(), "/")
	    repoPath2 := strings.Split(repoPath1[1], ",")

	    // initialize Github repo releases
	    client := github.NewClient(nil)
		ctx := context.Background()
		opt := &github.ListOptions{PerPage: 10}
		releases, _, err := client.Repositories.ListReleases(ctx, repoPath1[0], repoPath2[0], opt)
		if err != nil {

			fmt.Printf("Error found in %s/%s: %s\n",repoPath1[0], repoPath2[0], err)
			continue
			// log the error and continue to print the releases of next repo
		}
		minVersion := semver.New(repoPath2[1])
		allReleases := make([]*semver.Version, len(releases))
		for i, release := range releases {
			versionString := *release.TagName
			if versionString[0] == 'v' {
				versionString = versionString[1:]
			}
			allReleases[i] = semver.New(versionString)
		}
		versionSlice := LatestVersions(allReleases, minVersion)

		fmt.Printf("latest versions of %s/%s: %s\n",repoPath1[0],repoPath2[0], versionSlice)
	  	}

}
