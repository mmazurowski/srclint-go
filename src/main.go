package main

import (
	"log"
	"os"
	"path"

	"srclint/src/config"
	"srclint/src/crawler"
	"srclint/src/matcher"
	"srclint/src/printer"
)

func main() {

	var checkedDirectory = getCheckedPath("01")
	var configPath = getConfig(checkedDirectory)

	var configContent, err = config.ReadConfig(configPath)

	if err != nil {
		os.Exit(1)
	}

	var files, crawlErr = crawler.Crawl(checkedDirectory, configContent.Ignore)

	if crawlErr != nil {
		log.Fatalln(crawlErr.Error())
	}

	var missingRequired = matcher.MatchPaths(configContent.Required, files)
	var structure = matcher.MatchStructure(append(configContent.Structure, configContent.Required...), files)

	if len(missingRequired) > 0 {
		printer.PrintMissingRequired(missingRequired)
		os.Exit(1)
	} else {
		printer.PrintSuccessfulRequired()
	}

	if len(structure) > 0 {
		printer.PrintStructureFailed(structure)
		os.Exit(2)
	} else {
		printer.PrintStructureSuccess()
	}

	os.Exit(0)
}

func getCheckedPath(testCase string) string {
	var pwd, _ = os.Getwd()
	return path.Join(pwd, "tests", testCase)
}

func getConfig(pwd string) string {
	return path.Join(pwd, "srclint.yml")
}
