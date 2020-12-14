package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
		cs := strings.Split(line, " ")
		if len(cs) != 10 {
			panic("invalid line")
		}
		h := hand{
			player1: cards{
				card(cs[0]),
				card(cs[1]),
				card(cs[2]),
				card(cs[3]),
				card(cs[4]),
			},
			player2: cards{
				card(cs[5]),
				card(cs[6]),
				card(cs[7]),
				card(cs[8]),
				card(cs[9]),
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
	player1 cards
	player2 cards
}

func (hand) isPlayer1Win() bool {
	return false
}

type card string

func (c card) number() int {
	i, _ := strconv.Atoi(string(c[0]))
	return i
}

func (c card) suit() string {
	return string(c[1])
}

type cards [5]card

func (c cards) isSampleSuit() bool {
	s := c[0].suit()
	for _, card := range c {
		if card.suit() != s {
			return false
		}
	}
	return true
}
