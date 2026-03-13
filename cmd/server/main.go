package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	v1 "asymptotic.world/api/phonebook/v1"
)

type InMemoryDB struct {
	mu     sync.RWMutex
	people map[string]*v1.Person
}

func (db *InMemoryDB) GetPerson(ctx context.Context, name string) (*v1.Person, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	p, ok := db.people[name]
	if !ok {
		return nil, fmt.Errorf("person not found")
	}
	return p, nil
}

func (db *InMemoryDB) SavePerson(ctx context.Context, p *v1.Person) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if db.people == nil {
		db.people = make(map[string]*v1.Person)
	}
	db.people[p.Name] = p
	return nil
}

func main() {
	db := &InMemoryDB{
		people: map[string]*v1.Person{
			"Johan": {Name: "Johan", Number: "123456"},
		},
	}
	server := v1.NewServer(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/person", server.PersonHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
