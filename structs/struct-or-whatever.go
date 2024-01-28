package main

import (
	"errors"
	"fmt"
	"time"
)

type Memes struct {
	firstMeme   string
	anotherMeme string
	createdAt   time.Time
}

func main() {

	var userMemes Memes

	// userMemes.firstMeme = getUserData("Please enter a first meme: ")
	// userMemes.anotherMeme = getUserData("Please enter a another meme: ")

	firstMeme := getUserData("Please enter a first meme: ")
	anotherMeme := getUserData("Please enter a another meme: ")

	userMemes = newMemes(firstMeme, anotherMeme)

	userMemes.outputUserMemes()
	userMemes.clearUserMemes()
	userMemes.outputUserMemes()
}

// Method of Memes struct
func (m Memes) outputUserMemes() {
	fmt.Println("First meme: ", m.firstMeme)
	fmt.Println("Another meme: ", m.anotherMeme)
	fmt.Println("Created at: ", m.createdAt)
}

func (m *Memes) clearUserMemes() {
	m.firstMeme = ""
	m.anotherMeme = ""
}

func newMemes(firstMeme string, anotherMeme string) (*Memes, error) {
	if (firstMeme == "") || (anotherMeme == "") {
		return nil, errors.New("Please enter a valid meme")
	}
	return &Memes{
		firstMeme,
		anotherMeme,
		time.Now(),
	}, nil
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
