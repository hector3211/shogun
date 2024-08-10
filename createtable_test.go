package shogun

import "testing"

func TestCreateTable(t *testing.T) {
	query := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY,name TEXT NOT NULL);"
	ct := NewCreateTableBuilder().
		CreateTable("users").
		IfNotExists().
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		Define("name", "TEXT", "NOT NULL").
		String()

	if ct != query {
		t.Fatalf("TestCreateTable failed wanted: %s, got %s", query, ct)
	}
}

func TestCreateTableBuild(t *testing.T) {
	query := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY,name TEXT NOT NULL);"
	ct := NewCreateTableBuilder().
		CreateTable("users").
		IfNotExists().
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		Define("name", "TEXT", "NOT NULL").
		Build()

	if ct != query {
		t.Fatalf("TestCreateTableBuild failed test wanted: %s, got %s", query, ct)
	}
}

func TestCreateTableTwo(t *testing.T) {
	query := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY);"
	ct := NewCreateTableBuilder().
		CreateTable("users").
		IfNotExists().
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		String()

	if ct != query {
		t.Fatalf("TestCreateTableTwo failed wanted: %s, got %s", query, ct)
	}
}

func TestCreateTableThree(t *testing.T) {
	query := "CREATE TABLE users (id INT NOT NULL PRIMARY KEY);"
	ct := NewCreateTableBuilder().
		CreateTable("users").
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		String()

	if ct != query {
		t.Fatalf("TestCreateTableThree failed wanted: %s, got %s", query, ct)
	}
}

func TestCreateTableFour(t *testing.T) {
	query := "CREATE TABLE users (id INT NOT NULL PRIMARY KEY);"
	ct := CreatTable("users").
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		String()

	if ct != query {
		t.Fatalf("TestCreateTableThree failed wanted: %s, got %s", query, ct)
	}
}
