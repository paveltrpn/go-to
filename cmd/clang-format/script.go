package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	var (
		fileList []string
	)

    // Match only c cpp h related files
	cxxFile, err := regexp.Compile(".*\\.(cpp|c|h)$")
	if err != nil {
		log.Fatal(err)
	}

    // Exclude directories that contains /build/ and /extern/
	excludeDir, err := regexp.Compile(".*(\\b/build/\\b)|(\\b/extern/\\b).*")
	if err != nil {
		log.Fatal(err)
	}

	var wd, _ = os.Getwd()

	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			str := path + "/" + info.Name()
			cxx := cxxFile.MatchString(str)
			exclude := excludeDir.MatchString(str)
			if cxx && !exclude {
				fileList = append(fileList, str)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileList {
		fmt.Printf("%s\n", file)
	}
}
