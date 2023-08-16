package handler

import (
	"encoding/json"
	"net/http"
	"os"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Response struct {
	StatusCode int                                                     `json:"status_code"`
	Data       *gochatgptsdk.ModelImagesResponse[gochatgptsdk.DataURL] `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := c.ImagesGenerations(gochatgptsdk.ModelImages{
		Prompt: "Generate random images",
	})
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error.Message,
		})
	}

	json.NewEncoder(w).Encode(Response{
		StatusCode: http.StatusOK,
		Data:       resp,
	})
}
