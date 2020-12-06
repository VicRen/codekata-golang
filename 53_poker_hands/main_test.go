package main

import "testing"

func Test_hand_isPlayer1Win(t *testing.T) {
	type fields struct {
		player1 [5]string
		player2 [5]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"test1",
			fields{
				player1: [5]string{"5H", "5C", "6S", "7S", "KD"},
				player2: [5]string{"2C", "3S", "8S", "8D", "TD"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ha := hand{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			if got := ha.isPlayer1Win(); got != tt.want {
				t.Errorf("isPlayer1Win() = %v, want %v", got, tt.want)
			}
		})
	}
}
