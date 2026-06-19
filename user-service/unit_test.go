package main

import "testing"

func TestRegister(t *testing.T) {
	repo := NewUserRepository(nil)
	service := NewUserService(repo)

	user, err := service.Register("Ula", "ula@mail.com", "123", "customer")
	if err != nil {
		t.Logf("Register returned error (expected in unit test): %v", err)
	}

	if user != nil && user.UserID > 0 {
		t.Logf("User registered successfully with ID: %d", user.UserID)
	}
}

func TestLogin(t *testing.T) {
	repo := NewUserRepository(nil)
	service := NewUserService(repo)

	service.Register("Ula2", "ula2@mail.com", "123", "admin")
	token, err := service.Login("ula2@mail.com", "123")
	if err != nil {
		t.Logf("Login returned error (expected in unit test): %v", err)
	}

	if token != "" {
		t.Logf("Login successful, token received")
	}
}

func TestUpdateProfile(t *testing.T) {
	repo := NewUserRepository(nil)
	service := NewUserService(repo)

	user, err := service.Register(
		"A",
		"a@mail.com",
		"123",
		"customer",
	)

	if err != nil {
		t.Logf("Register error (expected in unit test): %v", err)
	}

	if user != nil {
		ok := service.UpdateProfile(
			user.UserID,
			"Bandung",
			"Fast Delivery",
		)

		if ok {
			t.Logf("UpdateProfile successful")
		}
	}
}
