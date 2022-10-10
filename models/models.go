package models

import "time"

type TwitterSearch struct {
	Query string `json:"query"`
}

type TwitterResponseWithGCPSentimentScore struct {
	Data []struct {
		ID            string `json:"id"`
		PublicMetrics struct {
			RetweetCount int `json:"retweet_count"`
			ReplyCount   int `json:"reply_count"`
			LikeCount    int `json:"like_count"`
			QuoteCount   int `json:"quote_count"`
		} `json:"public_metrics"`
		Text        string    `json:"text"`
		AuthorId    string    `json:"author_id"`
		CreatedAt   time.Time `json:"created_at"`
		Source      string    `json:"source"`
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments,omitempty"`
		SentimentScore     float64 `json:"sentiment_score"`
		SentimentMagnitude float64 `json:"sentiment_magnitude"`
	}
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	}
}

type TwitterResponseWithGCPSentimentScoreTest struct {
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
