package main

import (
	"github.com/ymiz/go-spreadsheet/service"
	"golang.org/x/oauth2"
	"google.golang.org/api/sheets/v4"
	"log"
	"time"
)

func record(conf *Config, count int) {
	config := oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		Scopes:      []string{"https://www.googleapis.com/auth/spreadsheets"},
	}

	srv, err := service.SheetServiceCreator{Config: &config}.Create()
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := conf.SpreadsheetId
	sheetName := conf.SpreadsheetName
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			generateSheetTitles(conf.TargetWord),
		},
	}
	_, err = srv.Spreadsheets.Values.Update(
		spreadsheetId,
		sheetName+"!A1",
		valueRange,
	).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatal("spreadsheet title error ", err.Error())
	}

	valueRange = &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			generateSheetResult(count),
		},
	}
	_, err = srv.Spreadsheets.Values.Append(
		spreadsheetId,
		sheetName+"!A2",
		valueRange,
	).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Println("Unable to write result value ", err)
	}
}

func generateSheetTitles(targetKeyword string) []interface{} {
	return []interface{}{
		"date",
		targetKeyword,
	}
}

func generateSheetResult(count int) []interface{} {
	jst := time.FixedZone("Asia/Tokyo", 96060)
	return []interface{}{
		time.Now().In(jst).Format("2006-01-02 15:04:05"),
		count,
	}
}
