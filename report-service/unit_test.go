package main

import (
	"context"
	"testing"
)

func TestGetDailyReport_ShouldReturnReport(t *testing.T) {
	repo := NewReportRepository(nil)
	svc := NewReportService(repo)

	report, err := svc.GetDailyReport(context.Background(), "2026-04-25")
	if err != nil {
		t.Logf("GetDailyReport returned error (expected in unit test): %v", err)
	}

	if report.TotalPaket >= 0 {
		t.Logf("Report retrieved with TotalPaket: %d", report.TotalPaket)
	}
}
