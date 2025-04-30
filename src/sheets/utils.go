package sheets

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func GetCategories(srv *sheets.Service, spreadsheetID string) ([]string, error) {
	readRange := "Main!I3:I"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	var categories []string
	for _, row := range resp.Values {
		if len(row) > 0 {
			categories = append(categories, fmt.Sprintf("%v", row[0]))
		}
	}
	return categories, nil
}

func GetPaymentMethods(srv *sheets.Service, spreadsheetID string) ([]string, error) {
	readRange := "Main!O3:O6"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}

	var methods []string
	for _, row := range resp.Values {
		if len(row) > 0 {
			methods = append(methods, fmt.Sprintf("%v", row[0]))
		}
	}
	return methods, nil
}
