package main

import (
	"context"
	"testing"
)

func TestFillDatabase(t *testing.T) {
	ctx := context.Background()
	path := t.TempDir()
	db, err := InitDB(path + "/db.sqlite")
	if err != nil {
		t.Fatalf("Failed to create db tables: %v", err)
	}
	if err := AddDemoData(ctx, db); err != nil {
		t.Error(err)
	}
}
