package lesson_21

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestListHandler tests the /list endpoint
func TestListHandler(t *testing.T) {
	tests := []struct {
		name           string
		db             database
		method         string
		expectedStatus int
		expectedError  string
		expectedItems  []Item
	}{
		{
			name:           "success with items",
			db:             database{"shoes": 50, "socks": 5},
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedItems:  []Item{{Name: "shoes", Price: 50}, {Name: "socks", Price: 5}},
		},
		{
			name:           "empty database",
			db:             database{},
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedError:  "no items in database",
		},
		{
			name:           "invalid method",
			db:             database{"shoes": 50},
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "method must be GET",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/list", nil)
			w := httptest.NewRecorder()

			tt.db.list(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var resp ListResponse
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if tt.expectedError != "" {
				if resp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, resp.Error)
				}
			} else {
				if resp.Error != "" {
					t.Errorf("unexpected error: %q", resp.Error)
				}
				if len(resp.Data) != len(tt.expectedItems) {
					t.Errorf("expected %d items, got %d", len(tt.expectedItems), len(resp.Data))
				}
				for i, item := range resp.Data {
					if item.Name != tt.expectedItems[i].Name || item.Price != tt.expectedItems[i].Price {
						t.Errorf("expected item %v, got %v", tt.expectedItems[i], item)
					}
				}
			}
		})
	}
}

// TestAddHandler tests the /create endpoint
func TestAddHandler(t *testing.T) {
	tests := []struct {
		name           string
		db             database
		url            string
		method         string
		expectedStatus int
		expectedMsg    string
		expectedError  string
		expectedDB     database
	}{
		{
			name:           "success",
			db:             database{"shoes": 50},
			url:            "/create?item=shirt&price=20",
			method:         http.MethodPost,
			expectedStatus: http.StatusCreated,
			expectedMsg:    "item shirt added with price $20.00",
			expectedDB:     database{"shoes": 50, "shirt": 20},
		},
		{
			name:           "missing item",
			db:             database{"shoes": 50},
			url:            "/create?price=20",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item must be specified",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "missing price",
			db:             database{"shoes": 50},
			url:            "/create?item=shirt",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "price must be specified",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "invalid price",
			db:             database{"shoes": 50},
			url:            "/create?item=shirt&price=invalid",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid price: invalid",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "negative price",
			db:             database{"shoes": 50},
			url:            "/create?item=shirt&price=-10",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "price cannot be negative",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "item exists",
			db:             database{"shoes": 50},
			url:            "/create?item=shoes&price=60",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item already exists: shoes",
			expectedDB:     database{"shoes": 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			tt.db.add(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var resp Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if tt.expectedError != "" {
				if resp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, resp.Error)
				}
			} else {
				if resp.Error != "" {
					t.Errorf("unexpected error: %q", resp.Error)
				}
				if resp.Message != tt.expectedMsg {
					t.Errorf("expected message %q, got %q", tt.expectedMsg, resp.Message)
				}
			}

			if len(tt.db) != len(tt.expectedDB) {
				t.Errorf("expected db %v, got %v", tt.expectedDB, tt.db)
			}
			for k, v := range tt.expectedDB {
				if dbV, ok := tt.db[k]; !ok || dbV != v {
					t.Errorf("expected db[%q] = %v, got %v", k, v, dbV)
				}
			}
		})
	}
}

// TestUpdateHandler tests the /update endpoint
func TestUpdateHandler(t *testing.T) {
	tests := []struct {
		name           string
		db             database
		url            string
		method         string
		expectedStatus int
		expectedMsg    string
		expectedError  string
		expectedDB     database
	}{
		{
			name:           "success",
			db:             database{"shoes": 50},
			url:            "/update?item=shoes&price=60",
			method:         http.MethodPut,
			expectedStatus: http.StatusOK,
			expectedMsg:    "item shoes updated with price $60.00",
			expectedDB:     database{"shoes": 60},
		},
		{
			name:           "missing item",
			db:             database{"shoes": 50},
			url:            "/update?price=60",
			method:         http.MethodPut,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item must be specified",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "missing price",
			db:             database{"shoes": 50},
			url:            "/update?item=shoes",
			method:         http.MethodPut,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "price must be specified",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "invalid price",
			db:             database{"shoes": 50},
			url:            "/update?item=shoes&price=invalid",
			method:         http.MethodPut,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid price: invalid",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "negative price",
			db:             database{"shoes": 50},
			url:            "/update?item=shoes&price=-10",
			method:         http.MethodPut,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "price cannot be negative",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "item not found",
			db:             database{"shoes": 50},
			url:            "/update?item=shirt&price=60",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedError:  "item not found: shirt",
			expectedDB:     database{"shoes": 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			tt.db.update(w, req)

			var resp Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if len(tt.db) != len(tt.expectedDB) {
				t.Errorf("expected db %v, got %v", tt.expectedDB, tt.db)
			}
			for k, v := range tt.expectedDB {
				if dbV, ok := tt.db[k]; !ok || dbV != v {
					t.Errorf("expected db[%q] = %v, got %v", k, v, dbV)
				}
			}
		})
	}
}

// TestFetchHandler tests the /read endpoint
func TestFetchHandler(t *testing.T) {
	tests := []struct {
		name           string
		db             database
		url            string
		method         string
		expectedStatus int
		expectedMsg    string
		expectedError  string
	}{
		{
			name:           "success",
			db:             database{"shoes": 50},
			url:            "/read?item=shoes",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedMsg:    "item shoes has price $50.00",
		},
		{
			name:           "missing item",
			db:             database{"shoes": 50},
			url:            "/read",
			method:         http.MethodGet,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item must be specified",
		},
		{
			name:           "item not found",
			db:             database{"shoes": 50},
			url:            "/read?item=shirt",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedError:  "item not found: shirt",
		},
		{
			name:           "invalid method",
			db:             database{"shoes": 50},
			url:            "/read?item=shoes",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "method must be GET",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			tt.db.fetch(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var resp Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if tt.expectedError != "" {
				if resp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, resp.Error)
				}
			} else {
				if resp.Error != "" {
					t.Errorf("unexpected error: %q", resp.Error)
				}
				if resp.Message != tt.expectedMsg {
					t.Errorf("expected message %q, got %q", tt.expectedMsg, resp.Message)
				}
			}
		})
	}
}

// TestDropHandler tests the /delete endpoint
func TestDropHandler(t *testing.T) {
	tests := []struct {
		name           string
		db             database
		url            string
		method         string
		expectedStatus int
		expectedMsg    string
		expectedError  string
		expectedDB     database
	}{
		{
			name:           "success",
			db:             database{"shoes": 50, "socks": 5},
			url:            "/delete?item=shoes",
			method:         http.MethodDelete,
			expectedStatus: http.StatusOK,
			expectedMsg:    "item shoes deleted",
			expectedDB:     database{"socks": 5},
		},
		{
			name:           "missing item",
			db:             database{"shoes": 50},
			url:            "/delete",
			method:         http.MethodDelete,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item must be specified",
			expectedDB:     database{"shoes": 50},
		},
		{
			name:           "item not found",
			db:             database{"shoes": 50},
			url:            "/delete?item=shirt",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedError:  "item not found: shirt",
			expectedDB:     database{"shoes": 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			tt.db.drop(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var resp Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if tt.expectedError != "" {
				if resp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, resp.Error)
				}
			} else {
				if resp.Error != "" {
					t.Errorf("unexpected error: %q", resp.Error)
				}
				if resp.Message != tt.expectedMsg {
					t.Errorf("expected message %q, got %q", tt.expectedMsg, resp.Message)
				}
			}

			if len(tt.db) != len(tt.expectedDB) {
				t.Errorf("expected db %v, got %v", tt.expectedDB, tt.db)
			}
			for k, v := range tt.expectedDB {
				if dbV, ok := tt.db[k]; !ok || dbV != v {
					t.Errorf("expected db[%q] = %v, got %v", k, v, dbV)
				}
			}
		})
	}
}
