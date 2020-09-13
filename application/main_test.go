package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	var tt = []struct {
		id       string
		input    string
		expected string
	}{
		{"1", "-c=testconfig.json", "Configuration path : testconfig.json\n"},
		{"2", "-c=c", "Configuration path : c\n"},
		{"3", "-c=", "Configuration path : \n"},
		{"4", "", "Configuration path : config.json\n"},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			var stdout bytes.Buffer
			args := []string{"program", tc.input}
			err := run(args, &stdout)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			actual := stdout.String()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, actual)
			}
		})
	}
}
