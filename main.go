package main

import (
	"fmt"
	"github.com/cdipaolo/sentiment"
)

func main() {
	// Initialize the sentiment analysis model
	model, err := sentiment.Restore()
	if err != nil {
		fmt.Println("Error loading model:", err)
		return
	}

	// Sample text for sentiment analysis
	text := "I love programming in Golang!"

	// Perform sentiment analysis
	analysis := model.SentimentAnalysis(text, sentiment.English)

	// Display analysis results
	fmt.Printf("Score: %v\n", analysis.Score)
	fmt.Printf("Sentence: %s\n", text)

	// Interpretation of the score for sentiment
	if analysis.Score == 0 {
		fmt.Println("Neutral Sentiment")
	} else if analysis.Score < 0 {
		fmt.Println("Negative Sentiment")
	} else {
		fmt.Println("Positive Sentiment")
	}
}
