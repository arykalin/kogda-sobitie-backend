package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/arykalin/kogda-sobitie-backend/tests/migration/events"
	"go.uber.org/zap"

	sheet "github.com/arykalin/kogda-sobitie-backend/tests/migration/scheet"
)

func main() {

	url := "http://localhost:8080/event"
	method := "POST"

	sLoggerConfig := zap.NewDevelopmentConfig()
	sLoggerConfig.DisableStacktrace = true
	sLoggerConfig.DisableCaller = true
	sLogger, err := sLoggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := sLogger.Sugar()

	s, err := sheet.NewSheetService(logger)
	if err != nil {
		log.Fatalf("failed to init sheet client: %s", err)
	}

	spreadsheet, err := s.GetSheet("14lqlXGtiT6vHi6O4X3iGDy9fATI3kghb8bS1_szNA_o")
	if err != nil {
		logger.Fatalf("failed to get sheet data: %s", err)
	}

	gotSheet, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		logger.Fatalf("failed to get sheet data: %s", err)
	}

	sheetConfig := events.SheetConfig{
		DateIdx:        1,
		TitleIdx:       3,
		DescriptionIdx: 9,
		DurationIdx:    2,
		LinkIdx:        4,
		WhoManagesIdx:  5,
		ForWhomIdx:     6,
		WhereIdx:       8,
		Skip:           2,
	}

	e := events.NewEvents()
	err = e.AddEvents(gotSheet, &sheetConfig)
	if err != nil {
		logger.Fatalf("failed to add events: %s", err)
	}

	payload := strings.NewReader(`{
    "date":"24.09.2021",
    "title":"Чтения на среднем",
    "duration":"3-5 часов",
    "link":"https://vk.com/4tenia",
    "who_manages":"Татка",
    "for_whom":"для всех",
    "where":"Дом Средний",
    "description":"четния по пятницам на среднем",
    "wanting_people":"15"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
