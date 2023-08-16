package handler

import gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Response struct {
	StatusCode int                                                     `json:"status_code"`
	Data       *gochatgptsdk.ModelImagesResponse[gochatgptsdk.DataURL] `json:"data"`
}
