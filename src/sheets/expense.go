package sheets

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/sheets/v4"
)

func AppendExpense(srv *sheets.Service, spreadsheetID, category string, amount float64, description, paidBy, method string) error {
	ist, _ := time.LoadLocation("Asia/Kolkata")
	currentMonth := time.Now().In(ist).Month().String()

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, fmt.Sprintf("%v!B5:B1000", currentMonth)).Do()
	if err != nil {
		return fmt.Errorf("error fetching values: %w", err)
	}

	nonEmptyRows := 0
	for _, row := range resp.Values {
		if len(row) > 0 && row[0] != "" {
			nonEmptyRows++
		}
	}
	nextRow := nonEmptyRows + 5

	writeRange := fmt.Sprintf("%v!B%d:G%d", currentMonth, nextRow, nextRow)

	date := time.Now().In(ist).Format("1/2/2006")
	values := [][]interface{}{{
		date,
		fmt.Sprintf("%.2f", amount),
		description,
		category,
		paidBy,
		method,
	}}

	vr := &sheets.ValueRange{Values: values}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, vr).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return fmt.Errorf("failed to append expense: %w", err)
	}

	log.Printf("Expense: %v added in %v sheet", values, currentMonth)

	return nil
}
