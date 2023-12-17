package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	// Replace with the actual import path
)

func TestCreateCustomer(t *testing.T) {
	app := yourapp.New()

	payload := []byte(`{"owner": "John Doe", "type": "Sedan", "color": "Blue", "vehicle_no": "ABC123", "model": "XYZ", "defective_part": "Engine", "amount": "100"}`)
	req, err := http.NewRequest("POST", "/customer", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Add additional checks based on your application's behavior
	// For example, you might want to check the response body or the database state
}

func TestGetCustomers(t *testing.T) {
	app := yourapp.New()

	req, err := http.NewRequest("GET", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Add additional checks based on your application's behavior
	// For example, you might want to check the response body or the returned data
}

func TestUpdateCustomer(t *testing.T) {
	app := yourapp.New()

	// Assuming you have an existing customer with ID 1
	updatePayload := []byte(`{"item": "owner", "new_item": "New Owner"}`)
	req, err := http.NewRequest("PUT", "/customer/1", bytes.NewBuffer(updatePayload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Add additional checks based on your application's behavior
	// For example, you might want to check the response body or the updated data in the database
}

func TestDeleteCustomer(t *testing.T) {
	app := yourapp.New()

	// Assuming you have an existing customer with ID 1
	req, err := http.NewRequest("DELETE", "/customer/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Add additional checks based on your application's behavior
	// For example, you might want to check the response body or the deleted data in the database
}
