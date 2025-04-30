package sheets

import (
	"context"
	"fmt"

	"github.com/awwwm4n/expense_tracker/src/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetSheetsService(cfg config.Config) (*sheets.Service, error) {

	ctx := context.Background()
	creds := []byte(cfg.GoogleCredsJSON)

	config, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, fmt.Errorf("error parsing service account JSON: %w", err)
	}

	client := config.Client(ctx)
	return sheets.NewService(ctx, option.WithHTTPClient(client))
}
