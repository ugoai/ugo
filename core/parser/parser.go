package parser

import (
	"fmt"
	"log"

	prose "github.com/jdkato/prose/v2"
)

type Parser struct {
	QuestionsWords []string
}

func (parser Parser) Sentence(sentence string) {
	extracted, proseError := prose.NewDocument(sentence)
	if proseError != nil {
		log.Fatal(proseError)
	}
	fmt.Println("Tokens: ")
	var words []string
	for _, token := range extracted.Tokens() {
		fmt.Println("\t", token.Text, token.Tag, token.Label)
		words = append(words, token.Text)
	}

	fmt.Println("Entities: ")
	for _, entity := range extracted.Entities() {
		fmt.Println("\t", entity.Text, entity.Label)
	}

	fmt.Println("Sentences: ")
	for _, sentence := range extracted.Sentences() {
		fmt.Println("\t", sentence.Text)
	}
	parser.ParseNonWordsInSentence(words)
}

func (parser Parser) ParseNonWordsInSentence(words []string) {
	for word := range words {
		fmt.Println(words[word])
	}
}
