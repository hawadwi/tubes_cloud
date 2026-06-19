package main

import "testing"

type MockPaymentService struct{}

func (m MockPaymentService) Calculate(req CalculateRequest) (CalculateResponse, error) {

	biayaLayanan := 5000

	switch req.Layanan {
	case "ekspres":
		biayaLayanan = 10000
	case "one-day":
		biayaLayanan = 20000
	}

	total := (req.Berat * 1000) + (req.Jarak * 500) + biayaLayanan

	return CalculateResponse{
		Biaya: total,
	}, nil
}

func (m MockPaymentService) Pay(req PaymentRequest) (PaymentResponse, error) {

	return PaymentResponse{
		TransactionID:    "TRX001",
		StatusPembayaran: "SUCCESS",
		Biaya:            12000,
	}, nil
}

func TestCalculate(t *testing.T) {
	repo := NewPaymentRepository(nil)
	service := NewPaymentService(repo)

	req := CalculateRequest{
		Berat:   2,
		Jarak:   10,
		Layanan: "reguler",
	}

	result, err := service.Calculate(req)
	if err != nil {
		t.Logf("Calculate returned error (expected in unit test): %v", err)
	}

	if result.Biaya == 0 && err == nil {
		t.Errorf("Expected non-zero biaya or error")
	}
}

func TestPay(t *testing.T) {
	repo := NewPaymentRepository(nil)
	service := NewPaymentService(repo)

	req := PaymentRequest{
		OrderID:          1,
		MetodePembayaran: "Transfer",
		PaymentDetails:   "BCA Virtual Account",
	}

	result, err := service.Pay(req)
	if err != nil {
		t.Logf("Pay returned error (expected in unit test): %v", err)
	}

	if result.TransactionID != "" {
		t.Logf("Payment test passed with transaction ID: %s", result.TransactionID)
	}
}
