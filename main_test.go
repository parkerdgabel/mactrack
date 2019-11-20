package main

import (
	"os"
	"testing"
)

func TestDatabasePath(t *testing.T) {
	os.Unsetenv("MACTRACK_DATABASE")
	expected := "~/.config/mactrack/mactrack.db"
	actual := getDatabasePath()
	if actual != expected {
		t.Errorf("Mactrack datapath expected to be %v but was %v", expected, actual)
	}
}

func TestDatabasePathWithEnviromentVariableSet(t *testing.T) {
	expected := "~/Dropbox/mactrack/mactrack.db"
	os.Setenv("MACTRACK_DATABASE", expected)
	actual := getDatabasePath()
	if actual != expected {
		t.Errorf("Mactrack datapath expected to be %v but was %v", expected, actual)
	}
}
