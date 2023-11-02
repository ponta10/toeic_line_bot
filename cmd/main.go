package main

import (
	"fmt"
	"log"
	"github.com/ponta10/toeic_line_bot/pkg/quiz"
)

func main() {
	q, err := quiz.NewQuiz("../data/words.csv")
	if err != nil {
		log.Fatalf("クイズの初期化に失敗しました: %s", err)
	}

	err = q.Start()
	if err != nil {
		fmt.Println(err)
	}
}