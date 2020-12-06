package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, err := os.Open("./p054_poker.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var input []hand
	count := 0
	for {
		a, _, e := br.ReadLine()
		if e == io.EOF {
			break
		}
		line := string(a)
		cards := strings.Split(line, " ")
		if len(cards) != 10 {
			panic("invalid line")
		}
		h := hand{
			player1: [5]string{
				cards[0],
				cards[1],
				cards[2],
				cards[3],
				cards[4],
			},
			player2: [5]string{
				cards[5],
				cards[6],
				cards[7],
				cards[8],
				cards[9],
			},
		}
		input = append(input, h)
		if h.isPlayer1Win() {
			count++
		}
	}
	fmt.Println(count)
}

type hand struct {
	player1 [5]string
	player2 [5]string
}

func (hand) isPlayer1Win() bool {
	return false
}
