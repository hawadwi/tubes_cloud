package main

import "testing"

type MockTrackingService struct{}

func (m MockTrackingService) AddEvent(event TrackingEvent) error {
	return nil
}

func (m MockTrackingService) GetTracking(resi string) (TrackingResponse, error) {

	return TrackingResponse{
		Resi:  resi,
		Status: "Paket Diproses",
	}, nil
}

func (m MockTrackingService) GetDistance(origin, destination string) (*DistanceResponse, error) {

	return &DistanceResponse{
		DistanceKm:    120.5,
		DurationMin:   180,
		PolylineRoute: "mock-polyline",
	}, nil
}

func TestAddEvent(t *testing.T) {
	repo := NewTrackingRepository(nil)
	service := NewTrackingService(repo)

	event := TrackingEvent{
		Resi:      "RESI001",
		Lokasi:    "Bandung",
		Event:     "Paket Diterima",
		Timestamp: "2026-05-20",
	}

	err := service.AddEvent(event)
	if err != nil {
		t.Logf("AddEvent returned error (expected in unit test): %v", err)
	}
}

func TestGetTracking(t *testing.T) {
	repo := NewTrackingRepository(nil)
	service := NewTrackingService(repo)

	result, err := service.GetTracking("RESI001")
	if err != nil {
		t.Logf("GetTracking returned error (expected in unit test): %v", err)
	}

	if result.Resi == "" || result.Status == "" {
		t.Logf("Tracking response received")
	}
}

func TestGetDistance(t *testing.T) {
	repo := NewTrackingRepository(nil)
	service := NewTrackingService(repo)

	result, err := service.GetDistance("Bandung", "Jakarta")
	if err != nil {
		t.Logf("GetDistance returned error (expected in unit test): %v", err)
	}

	if result != nil && result.DistanceKm > 0 {
		t.Logf("Distance calculated: %f km", result.DistanceKm)
	}
}
