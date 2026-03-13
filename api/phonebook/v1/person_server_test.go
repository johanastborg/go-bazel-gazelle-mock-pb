package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestPersonHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDatabase(ctrl)

	p := &Person{
		Name:   "Johan",
		Number: "123456",
	}

	mockDB.EXPECT().GetPerson(gomock.Any(), "Johan").Return(p, nil)

	server := NewServer(mockDB)
	req := httptest.NewRequest(http.MethodGet, "/person?name=Johan", nil)
	w := httptest.NewRecorder()

	server.PersonHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.Status)
	}

	var got Person
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Name != p.Name || got.Number != p.Number {
		t.Errorf("expected person %+v, got %+v", p, got)
	}
}
