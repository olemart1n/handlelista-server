package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gage-technologies/mistral-go"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	client := MistralClient()

	router.GET("/api/prompt1", func(c *gin.Context) {
        recommendation := ""
        recommendation, err := mistralChat(client, prompt1())
        if err != nil {
            c.Set("Content-Type", "application/json")
            c.JSON(http.StatusServiceUnavailable, gin.H{
                "error": "Mistral service error",
            })
            return
        }
        c.Writer.Header().Set("Content-Type", "application/json")
        c.IndentedJSON(http.StatusOK, recommendation)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }

    // Start the server with the retrieved port
    addr := fmt.Sprintf(":%s", port)
    http.ListenAndServe(addr, router)
}

func prompt1 () (mistral.ChatMessage) {
    currentTime := time.Now().Format(time.RFC3339)
    content := currentTime + " Generer ett forslag til en middag. Svar i from av et objekt som består av navnet på middagen, og en liste med ingredienser. Svar med bare objektet og bruk små bokstaver, slik som dette {middag: \"\", ingredienser: []}"
    prompt1 := mistral.ChatMessage {
		Role: "user",
		Content: content,
	}
    return prompt1
}

func mistralChat(m *mistral.MistralClient, prompt mistral.ChatMessage) (string, error) {
    response, err := m.Chat("mistral-tiny", []mistral.ChatMessage{prompt}, nil)
    if err != nil {
        fmt.Print(err)
        return "", err
    }
    if response.Choices[0].Message.Content == "" {
        fmt.Println("Empty response received from Mistral")
    }
    return response.Choices[0].Message.Content, nil
}