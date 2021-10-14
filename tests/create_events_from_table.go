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

	//url := "http://95.216.158.138:80/event"
	url := "http://127.0.0.1:8080/event"
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

	gotEvents := e.GetEvents()

	for _, event := range gotEvents {
		payload := strings.NewReader(fmt.Sprintf(`{
    "date":"%s",
    "title":"%s",
    "duration":"%s",
    "link":"%s",
    "org":"%s",
    "target":"%s",
    "where":"%s",
    "description":"%s",
    "amount":"%s"
}`, event.Date,
			event.Title,
			event.Duration,
			event.Link,
			event.WhoManages,
			event.ForWhom,
			event.Where,
			event.Description,
			event.WantingPeople))

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

}
