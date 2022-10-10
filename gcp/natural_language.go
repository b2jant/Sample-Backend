package gcp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GCPNaturalLanguage struct {
	Document struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"document"`
	EncodingType string `json:"encodingType"`
}

type GCPNaturalLanguageResponse struct {
	DocumentSentiment struct {
		Magnitude float64 `json:"magnitude"`
		Score     float64 `json:"score"`
	} `json:"documentSentiment"`
	Language  string `json:"language"`
	Sentences []struct {
		Text struct {
			Content     string `json:"content"`
			BeginOffset int    `json:"beginOffset"`
		} `json:"text"`
		Sentiment struct {
			Magnitude float64 `json:"magnitude"`
			Score     float64 `json:"score"`
		} `json:"sentiment"`
	} `json:"sentences"`
}

func GetSentimentGCP(text string) GCPNaturalLanguageResponse {
	gcp_api_key, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !ok {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS not found")
	}

	gcp_api_url := "https://language.googleapis.com/v1/documents:analyzeSentiment?key=" + gcp_api_key
	gcp_api_body := []byte("{\"document\":{\"type\":\"PLAIN_TEXT\",\"content\":\"" + text + "\"},\"encodingType\":\"UTF8\"}")
	resp, err := http.Post(gcp_api_url, "application/json", bytes.NewBuffer(gcp_api_body))
	if err != nil {
		log.Fatal(err)
	}
	gcp_api_response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gcp_response GCPNaturalLanguageResponse
	json.Unmarshal(gcp_api_response, &gcp_response)

	return gcp_response
}

func GetSentimentGCPScore(response GCPNaturalLanguageResponse) float64 {
	return response.DocumentSentiment.Score
}

func PostGCP(url string, body string) string {
	return ""
}
