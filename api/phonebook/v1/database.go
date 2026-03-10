package v1

import (
	"context"
)

//go:generate mockgen -destination=mock_database_test.go -package=v1 . Database
type Database interface {
	GetPerson(ctx context.Context, name string) (*Person, error)
	SavePerson(ctx context.Context, p *Person) error
}
