package sheet

import (
	"testing"

	"go.uber.org/zap"
	"google.golang.org/api/sheets/v4"
)

func Test_sheet_GetSheetRange(t *testing.T) {
	answersSheetId := ""
	answersReadRange := ""
	sLoggerConfig := zap.NewDevelopmentConfig()
	sLoggerConfig.DisableStacktrace = true
	sLoggerConfig.DisableCaller = true
	sLogger, err := sLoggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := sLogger.Sugar()
	type args struct {
		spreadsheetId string
		readRange     string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *sheets.ValueRange
		wantErr  bool
	}{
		{name: "get answers", args: args{
			spreadsheetId: answersSheetId,
			readRange:     answersReadRange,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSheetService(logger)
			if err != nil {
				t.Fatalf("failed to init sheet client: %s", err)
			}
			gotResp, err := s.GetSheet(tt.args.spreadsheetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSheetRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got resp: %v", gotResp)
		})
	}
}
