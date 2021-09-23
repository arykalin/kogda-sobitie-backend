package events

import (
	"encoding/json"
	"io/ioutil"

	"github.com/arykalin/kogda-sobitie-backend/models"
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
	Skip           int
}

type events struct {
	events []models.Event
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
		event := models.Event{}

		//	DateIdx        int
		var date string
		if len(sheet.Rows[i]) > config.DateIdx {
			date = sheet.Rows[i][config.DateIdx].Value
		}
		event.Date = date

		//	TitleIdx       int
		var title string
		if len(sheet.Rows[i]) > config.TitleIdx {
			title = sheet.Rows[i][config.TitleIdx].Value
		}
		event.Title = title

		//	DurationIdx    int
		var duration string
		if len(sheet.Rows[i]) > config.DurationIdx {
			duration = sheet.Rows[i][config.DurationIdx].Value
		}
		event.Duration = duration

		//	LinkIdx        int
		var link string
		if len(sheet.Rows[i]) > config.LinkIdx {
			link = sheet.Rows[i][config.LinkIdx].Value
		}
		event.Link = link

		//	WhoManagesIdx  int
		var whoManages string
		if len(sheet.Rows[i]) > config.WhoManagesIdx {
			whoManages = sheet.Rows[i][config.WhoManagesIdx].Value
		}
		event.WhoManages = whoManages

		//	ForWhomIdx     int
		var forWhom string
		if len(sheet.Rows[i]) > config.ForWhomIdx {
			forWhom = sheet.Rows[i][config.ForWhomIdx].Value
		}
		event.ForWhom = forWhom

		//	WhereIdx     int
		var where string
		if len(sheet.Rows[i]) > config.WhereIdx {
			forWhom = sheet.Rows[i][config.WhereIdx].Value
		}
		event.Where = where

		//	DescriptionIdx int
		var description string
		if len(sheet.Rows[i]) > config.DescriptionIdx {
			description = sheet.Rows[i][config.DescriptionIdx].Value
		}
		if description == "" {
			description = "description"
		}
		event.Description = description

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
func NewEvents() UsersInt {
	return &events{
		events: []models.Event{},
	}
}
