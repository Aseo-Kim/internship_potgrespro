package main

import (
	"reflect"
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	expected := []Config{
		{
			Path: "/path/to/directory",
			Commands: []string{
				"echo 'File changed'",
			},
		},
		{
			Path: "/another/path/to/directory",
			Commands: []string{
				"echo 'File changed in another directory'",
			},
		},
	}

	configArr, err := parseConfiguration("test_config.yaml")
	if err != nil {
		t.Fatalf("parseConfiguration returned error: %v", err)
	}

	if !reflect.DeepEqual(configArr, expected) {
		t.Errorf("parseConfiguration returned unexpected result. Got: %v, expected: %v", configArr, expected)
	}
}

func TestRunCommands(t *testing.T) {
	config := Config{
		Commands: []string{"cmd /c echo Hello", "cmd /c echo World"},
	}
	err := RunCommands(config)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
