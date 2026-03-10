package v1

import (
	"testing"
)

func TestCreatePerson(t *testing.T) {
	p := &Person{
		Name:   "John Doe",
		Number: "123-456-7890",
	}

	if p.Name != "John Doe" {
		t.Errorf("expected p.Name John Doe, got %s", p.Name)
	}
	if p.Number != "123-456-7890" {
		t.Errorf("expected p.Number 123-456-7890, got %s", p.Number)
	}
}
