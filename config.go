package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GitRepositoryUrl string `json:"git_repository_url"`
	TargetWord       string `json:"target_word"`
	TargetDirPath    string `json:"target_dir_path"`
	ClientId         string `json:"client_id"`
	ClientSecret     string `json:"client_secret"`
	SpreadsheetId    string `json:"spreadsheet_id"`
	SpreadsheetName  string `json:"spreadsheet_name"`
	FileSearchRegexp string `json:"file_search_regexp"`
}

func parseConfigJson(fileName string) (*Config, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
