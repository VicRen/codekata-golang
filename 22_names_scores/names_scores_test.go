package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readNames(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  []string
	}{
		{
			"empty_input",
			"",
			nil,
		},
		{
			"one_name",
			"\"VIC\"",
			[]string{"VIC"},
		},
		{
			"two_names",
			"\"VIC\",\"COCO\"",
			[]string{"VIC", "COCO"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := readNames(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("readNames() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_worthOf(t *testing.T) {
	tt := []struct {
		name string
		word string
		want int
	}{
		{
			"COLIN",
			"COLIN",
			53,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("%d", 'A')
			got := worthOf(tc.word)
			if got != tc.want {
				t.Errorf("worthOf() = %d, want %d", got, tc.want)
			}
		})
	}
}
