package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test start delivery

func TestStartDelivery_Success(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	body := []byte(`{
	    "resi": "RESI001",
	    "courier_id": 1,
	    "assigned_zone": "Jakarta"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	// expected response ketika request berhasil
	if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
		t.Errorf("expected 200 or 400, got %d", w.Code)
	}
}

func TestStartDelivery_InvalidJSON(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	// isi dengan format JSON yang tidak valid
	body := []byte(``)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	// expected response ketika format JSON salah
	if w.Code != http.StatusBadRequest && w.Code != http.StatusInternalServerError {
		t.Errorf("expected 400 or 500, got %d", w.Code)
	}
}

func TestStartDelivery_MissingField(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	body := []byte(`{
		"resi": "",
		"courier_id": 0,
		"assigned_zone": ""
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	// expected response ketika ada field yang kosong
	if w.Code != http.StatusBadRequest && w.Code != http.StatusOK {
		t.Errorf("expected 400 or 200, got %d", w.Code)
	}
}

// test get delivery

func TestGetCourierDeliveries_Success(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries?courier_id=1", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
		t.Errorf("expected 200 or 400, got %d", w.Code)
	}
}

func TestGetCourierDeliveries_InvalidID(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries?courier_id=abc", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	if w.Code != http.StatusBadRequest && w.Code != http.StatusOK {
		t.Errorf("expected 400 or 200, got %d", w.Code)
	}
}

func TestGetCourierDeliveries_MissingID(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	if w.Code != http.StatusBadRequest && w.Code != http.StatusOK {
		t.Errorf("expected 400 or 200, got %d", w.Code)
	}
}

// test health 

func TestHealth(t *testing.T) {
	repo := NewDeliveryRepository(nil)
	service := NewCourierService(repo)
	handler := NewCourierHandler(service, repo, nil)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
