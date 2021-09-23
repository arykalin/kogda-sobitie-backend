package sheet

import (
	"io/ioutil"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type sheet struct {
	logger  *zap.SugaredLogger
	service *spreadsheet.Service
}

type Sheet interface {
	GetSheet(id string) (spreadsheet.Spreadsheet, error)
}

func (s sheet) GetSheet(id string) (spreadsheet.Spreadsheet, error) {
	return s.service.FetchSpreadsheet(id)
}

func NewSheetService(logger *zap.SugaredLogger) (Sheet, error) {
	data, err := ioutil.ReadFile("/tmp/client_secret.json")
	if err != nil {
		return &sheet{}, err
	}
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return &sheet{}, err
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	return &sheet{
		logger:  logger,
		service: service,
	}, nil
}
