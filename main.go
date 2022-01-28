package main

import (
	"bytes"
	pt "github.com/monochromegane/the_platinum_searcher"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	config, err := parseConfigJson("./config.json")
	if err != nil {
		log.Fatal("parse config error ", err.Error())
	}
	if config == nil {
		log.Fatal("empty config")
	}
	err = exec.Command("git", "clone", config.GitRepositoryUrl, "temp").Run()
	if err != nil {
		log.Fatal("git clone error ", err.Error())
	}

	buf := new(bytes.Buffer)
	searcher := pt.PlatinumSearcher{Out: buf, Err: os.Stderr}
	args := []string{config.TargetWord, "./temp/" + config.TargetDirPath, "-i", "--count", "--nocolor"}
	if config.FileSearchRegexp != "" {
		args = append(args, "--file-search-regexp="+config.FileSearchRegexp)
	}

	exitCode := searcher.Run(args)
	result := buf.String()
	resultList := strings.Split(result, "\n")
	sum := 0
	for _, r := range resultList {
		if len(r) <= 0 {
			continue
		}
		count, err := strconv.Atoi(strings.Split(r, ":")[1])
		if err != nil {
			log.Fatal("parse error", err.Error())
		}
		sum += count
	}
	record(config, sum)

	err = os.RemoveAll("./temp")
	if err != nil {
		log.Fatal("remove dir error ", err.Error())
	}
	os.Exit(exitCode)
}
