package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"./esv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Verse to find: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	verse, err := esv.GetVerse(url.QueryEscape(text))
	if err != nil {
		panic(err)
	}

	println(verse)
}
