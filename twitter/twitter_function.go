package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type TwitterResponse struct {
	Data []struct {
		ID                 string    `json:"id"`
		Text               string    `json:"text"`
		AuthorId           string    `json:"author_id"`
		CreatedAt          time.Time `json:"created_at"`
		Source             string    `json:"source"`
		SentimentScore     float64   `json:"sentiment_score"`
		SentimentMagnitude float64   `json:"sentiment_magnitude"`
	}
}

type DummyTwissResponse struct {
	Data []struct {
		ID                 string    `json:"id"`
		Text               string    `json:"text"`
		AuthorId           string    `json:"author_id"`
		CreatedAt          time.Time `json:"created_at"`
		Source             string    `json:"source"`
		SentimentScore     float64   `json:"sentiment_score"`
		SentimentMagnitude float64   `json:"sentiment_magnitude"`
	}
}

func GetTwitterResult(query string) TwitterResponse {
	token, ok := os.LookupEnv("TWITTERBEARERTOKEN")
	if !ok {
		log.Fatal("TWITTERBEARERTOKEN not found")
	}
	bearer_token := "Bearer " + token
	twitter_query := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?query=%s&tweet.fields=attachments,author_id,created_at,public_metrics,source", query)
	r, _ := http.NewRequest("GET", twitter_query, nil)
	r.Header.Add("Authorization", bearer_token)
	client := &http.Client{}
	resp, err := client.Do(r)
	CheckError(err)
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var twitter_response TwitterResponse
	json.Unmarshal(body, &twitter_response)
	return twitter_response
}

func GetDummyResult(query string) DummyTwissResponse {
	var twitter_response DummyTwissResponse
	json.Unmarshal([]byte(query), &twitter_response)
	return twitter_response
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
