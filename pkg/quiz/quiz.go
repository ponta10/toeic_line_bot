package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"
)

type Word struct {
	ID      string
	Word    string
	Meaning string
}

type Quiz struct {
	Words []Word
}

func NewQuiz(filepath string) (*Quiz, error) {
	words, err := readCSV(filepath)
	if err != nil {
		return nil, err
	}
	return &Quiz{Words: words}, nil
}

func readCSV(filename string) ([]Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var words []Word
	for _, record := range records {
		word := Word{
			ID:      record[0],
			Word:    record[1],
			Meaning: record[2],
		}
		words = append(words, word)
	}

	return words, nil
}

func (q *Quiz) Start() error {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	for _, w := range q.Words {
		fmt.Printf("%s の意味は何ですか？\n", w.Word)

		scanner.Scan()
		answer := scanner.Text()

		if answer == w.Meaning {
			fmt.Println("正解！!")
		} else {
			fmt.Printf("不正解！正しい答えは: %s です\n", w.Meaning)
		}
	}

	return nil
}

