package main

import "testing"

func TestRegister(t *testing.T) {
	user, err := Register("Ula", "ula@mail.com", "123", "customer")
	if err != nil {
		t.Errorf("Register returned error: %v", err)
	}

	if user.UserID <= 0 {
		t.Errorf("Expected UserID > 0, got %d", user.UserID)
	}

	if user.Email != "ula@mail.com" {
		t.Errorf("Expected email ula@mail.com, got %s", user.Email)
	}
}

func TestLogin(t *testing.T) {
	// First register a user
	Register("Ula2", "ula2@mail.com", "123", "admin")

	// Then try to login
	user, err := Login("ula2@mail.com", "123")
	if err != nil {
		t.Errorf("Login returned error: %v", err)
	}

	if user.Email != "ula2@mail.com" {
		t.Errorf("Expected email ula2@mail.com, got %s", user.Email)
	}
}

func TestLoginFailed(t *testing.T) {
	// Try to login with wrong password
	_, err := Login("seed@mail.com", "wrongpassword")
	if err == nil {
		t.Error("Expected login to fail with wrong password")
	}
}

func TestUpdateProfile(t *testing.T) {
	user, err := Register("TestUser", "test@mail.com", "123", "customer")
	if err != nil {
		t.Errorf("Register returned error: %v", err)
	}

	ok := UpdateProfile(user.UserID, "Bandung", "Fast Delivery")
	if !ok {
		t.Error("UpdateProfile should return true")
	}

	// Verify profile was updated
	profile := GetProfile(user.UserID)
	if profile == nil {
		t.Error("GetProfile returned nil")
	}

	if profile.Alamat != "Bandung" {
		t.Errorf("Expected alamat Bandung, got %s", profile.Alamat)
	}

	if profile.Preferensi != "Fast Delivery" {
		t.Errorf("Expected preferensi Fast Delivery, got %s", profile.Preferensi)
	}
}

func TestGetProfile(t *testing.T) {
	user, _ := Register("ProfileTest", "profile@mail.com", "123", "customer")

	profile := GetProfile(user.UserID)
	if profile == nil {
		t.Error("GetProfile returned nil")
	}

	if profile.Name != "ProfileTest" {
		t.Errorf("Expected name ProfileTest, got %s", profile.Name)
	}
}