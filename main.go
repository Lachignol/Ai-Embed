package main

import (
	"fmt"
	"log"


	"github.com/Lachignol/Ai-Embed/embeding"
	"github.com/parakeet-nest/parakeet/completion"

	"github.com/parakeet-nest/parakeet/embeddings"
	"github.com/parakeet-nest/parakeet/llm"
)

var Store = embeddings.BboltVectorStore{}

func init() {
	Store.Initialize("./embeding/embeddings.db")
	embeding.Embed(Store)

}

func main() {

	ollamaUrl := "http://localhost:11434"
	embeddingsModel := "all-minilm"
	chatModel := "mistral"

	systemContent := `tu est un generateur de resumé base toi uniquement sur ton documentsContent pour repondre`

	// Question for the Chat system
	userContent := `resume moi le document`

	// Create an embedding from the user question
	embeddingFromQuestion, err := embeddings.CreateEmbedding(
		ollamaUrl,
		llm.Query4Embedding{
			Model:  embeddingsModel,
			Prompt: userContent,
		},
		"question",
	)
	if err != nil {
		log.Fatalln("😡:", err)
	}
	fmt.Println("🔎 searching for similarity...")

	similarities, _ := Store.SearchSimilarities(embeddingFromQuestion, 0.3)

	// Generate the context from the similarities
	// This will generate a string with a content like this one:
	// `<context><doc>...<doc><doc>...<doc></context>`
	documentsContent := embeddings.GenerateContextFromSimilarities(similarities)

	fmt.Println("🎉 similarities", len(similarities))

	query := llm.Query{
		Model: chatModel,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "system", Content: documentsContent},
			{Role: "user", Content: userContent},
		},
		Options: llm.Options{
			Temperature: 0.4,
			RepeatLastN: 2,
		},
		Stream: false,
	}

	fmt.Println("😀 Question:",userContent)
	fmt.Println("🤖 Answer:")

	// Answer the question
	_, err = completion.ChatStream(ollamaUrl, query,
		func(answer llm.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	if err != nil {
		log.Fatal("😡:", err)
	}
}
