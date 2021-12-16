package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/arykalin/kogda-sobitie-backend/models"
	"go.uber.org/zap"
	"gopkg.in/Iwark/spreadsheet.v2"
)

const (
	userDataJsonFile = "user_data.json"
)

type Event = map[string]models.Event

type SheetConfig struct {
	DateIdx        int
	TitleIdx       int
	DurationIdx    int
	LinkIdx        int
	WhoManagesIdx  int
	ForWhomIdx     int
	WhereIdx       int
	DescriptionIdx int
	AmountIdx      int
	Skip           int
}

type events struct {
	events []models.Event
	logger *zap.SugaredLogger
}

type UsersInt interface {
	AddEvents(sheet *spreadsheet.Sheet, config *SheetConfig) (err error)
	GetEvents() []models.Event
	DumpUsers() error
}

func (u *events) AddEvents(sheet *spreadsheet.Sheet, config *SheetConfig) (err error) {
	for i := range sheet.Rows {
		if i < config.Skip {
			// skip
			continue
		}
		event, err := u.makeEvent(sheet.Rows[i], config)
		if err != nil {
			u.logger.Errorf("failed to make event with title: %s err: %w", event.Title, err)
			continue
		}
		u.events = append(u.events, event)
	}
	return err
}

func (u events) GetEvents() []models.Event {
	return u.events
}

func (u events) DumpUsers() error {
	file, err := json.MarshalIndent(u.events, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(userDataJsonFile, file, 0644) //nolint:gosec
	return err
}

func (u *events) makeEvent(cell []spreadsheet.Cell, config *SheetConfig) (event models.Event, err error) {
	//	TitleIdx       int
	var title string
	if len(cell) > config.TitleIdx {
		title = cell[config.TitleIdx].Value
	}
	event.Title = title

	//	DateIdx        int
	var date string
	if len(cell) > config.DateIdx {
		date = cell[config.DateIdx].Value
		cellLayout := "02.01.2006"
		t, err := time.Parse(cellLayout, date)
		if err != nil {
			return event, fmt.Errorf("failed to parse date: %s", err)
		}
		date = t.Format(cellLayout)
	}
	event.Date = date

	//	DurationIdx    int
	var duration string
	if len(cell) > config.DurationIdx {
		duration = cell[config.DurationIdx].Value
	}
	event.Duration = duration

	//	LinkIdx        int
	var link string
	if len(cell) > config.LinkIdx {
		link = cell[config.LinkIdx].Value
	}
	event.Link = link

	//	WhoManagesIdx  int
	var whoManages string
	if len(cell) > config.WhoManagesIdx {
		whoManages = cell[config.WhoManagesIdx].Value
	}
	event.Org = whoManages

	//	ForWhomIdx     int
	var forWhom string
	if len(cell) > config.ForWhomIdx {
		forWhom = cell[config.ForWhomIdx].Value
	}
	event.Target = forWhom

	//	WhereIdx     int
	where := "неизвестно"
	if len(cell) > config.WhereIdx {
		where = cell[config.WhereIdx].Value
	}
	event.Where = where

	//	AmountIdx     int
	amount := "10"
	if len(cell) > config.AmountIdx {
		amount = cell[config.AmountIdx].Value
	}
	event.Amount = amount

	//	DescriptionIdx int
	var description string
	if len(cell) > config.DescriptionIdx {
		description = cell[config.DescriptionIdx].Value
	}
	if description == "" {
		description = ""
	}
	event.Description = description
	return u.validate(event)
}

func (u *events) validate(event models.Event) (models.Event, error) {
	if event.Title == "" {
		return event, fmt.Errorf("missing title")
	}
	if event.Date == "" {
		return event, fmt.Errorf("missing date")
	}
	return event, nil
}

func NewEvents() UsersInt {
	sLoggerConfig := zap.NewDevelopmentConfig()
	sLoggerConfig.DisableStacktrace = true
	sLoggerConfig.DisableCaller = true
	sLogger, err := sLoggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := sLogger.Sugar()
	return &events{
		events: []models.Event{},
		logger: logger,
	}
}
