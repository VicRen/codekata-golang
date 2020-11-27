package main

import (
	"reflect"
	"testing"
)

func Test_hasCross(t *testing.T) {
	type args struct {
		orders string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"NES"},
			false,
		},
		{
			"test2",
			args{"NESWW"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCross(tt.args.orders); got != tt.want {
				t.Errorf("hasCross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPos(t *testing.T) {
	type args struct {
		current [2]int
		order   string
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			"test_E",
			args{[2]int{0, 0}, "E"},
			[2]int{1, 0},
		},
		{
			"test_S",
			args{[2]int{0, 0}, "S"},
			[2]int{0, -1},
		},
		{
			"test_W",
			args{[2]int{0, 0}, "W"},
			[2]int{-1, 0},
		},
		{
			"test_N",
			args{[2]int{0, 0}, "N"},
			[2]int{0, 1},
		},
		{
			"test1",
			args{[2]int{0, 0}, "N"},
			[2]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPos(tt.args.current, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pos_Move(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		order string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *pos
	}{
		{
			"test_E",
			fields{0, 0},
			args{"E"},
			&pos{1, 0},
		},
		{
			"test_S",
			fields{0, 0},
			args{"S"},
			&pos{0, -1},
		},
		{
			"test_W",
			fields{0, 0},
			args{"W"},
			&pos{-1, 0},
		},
		{
			"test_N",
			fields{0, 0},
			args{"N"},
			&pos{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pos{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			p.Move(tt.args.order)
			if !p.Equals(tt.want) {
				t.Errorf("pos Move()=%v, want %v", p, tt.want)
			}
		})
	}
}

func Test_hasCross2(t *testing.T) {
	type args struct {
		orders string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"NES"},
			false,
		},
		{
			"test2",
			args{"NESWW"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCross2(tt.args.orders); got != tt.want {
				t.Errorf("hasCross2() = %v, want %v", got, tt.want)
			}
		})
	}
}
