package main

import (
	"bytes"
	"crypto/rand"
	"io"
	"math/big"
	"os"
	"strings"
)

// GeneratePassword generates a xkcd style password.
func GeneratePassword(wordCount int) string {
	// Read the word file
	f, err := os.Open("words.txt")
	checkAndPanic(err)
	numberOfWords, _ := lineCounter(f)
	// If there are less words than the requested number, make the requested number
	// the number of words so we don't have an error.
	if wordCount > numberOfWords {
		wordCount = numberOfWords
	}
	// Pick 4 integers exclusively from 0 to number of words - 1 in the file
	indexMap := make(map[int]string)
	var biggestRandomPicked int
	for len(indexMap) < wordCount {
		randBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(numberOfWords)))
		randInt := int(randBigInt.Int64())
		if _, selected := indexMap[randInt]; !selected {
			indexMap[randInt] = ""
		}
		// Save the biggest random number generated which we will use when we traverse the file.
		if randInt > biggestRandomPicked {
			biggestRandomPicked = randInt
		}
	}
	// Extract these words from the file
	scanner, file := getWordFileScanner("words.txt")
	defer file.Close()

	lineNumber := 0
	for scanner.Scan() {
		if _, selected := indexMap[lineNumber]; selected {
			indexMap[lineNumber] = scanner.Text()
		}
		lineNumber++
		if lineNumber > biggestRandomPicked {
			break
		}
	}
	err = scanner.Err()
	checkAndLog(err)
	// Form the words into one string.
	var result string
	for _, v := range indexMap {
		result += " " + strings.ToLower(v)
	}
	// Remove the space character at the start of the string.
	result = strings.TrimLeft(result, " ")
	return result
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	// Loop until io.EOF
	for {
		// Read r into buf with c being the number of bytes read.
		c, err := r.Read(buf)
		// Count the number of new line characters.
		count += bytes.Count(buf[:c], lineSep)

		switch {
		// at End Of File return the count.
		case err == io.EOF:
			return count, nil
		// If there is an error, return the current count and the error.
		case err != nil:
			return count, err
		}
	}
}
