package routes

import (
	"github.com/b2jant/twiss/twiss_backend/gcp"
	"github.com/b2jant/twiss/twiss_backend/twitter"
	"github.com/gin-gonic/gin"
)

func FetchTwitter(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Please Provide A Path Variable",
	})
}

func FetchTwitterWithQuery(c *gin.Context) {
	query := c.Param("query")
	body := twitter.GetTwitterResult(query)
	c.JSON(200, body)
}

func FetchTwitterTest(c *gin.Context) {
	query := `{"Data":[{"id":"1540431656686723072","text":"RT @JonathanClamp1: $SPY $RDFN $AAPL $TSLA $AMZN \nMArket will crash after new home sales data, indication of recession by weak consumer pow…","author_id":"1354932119840071680","created_at":"2022-06-24T20:28:21Z","source":"Twitter for Android","sentiment_score":-0.7,"sentiment_magnitude":0.7},{"id":"1540431580501446656","text":"Las Vegas Sports Betting Giants Roll the Dice on Hollywood Talent  $para $cmcsa $dis $nflx $aapl $t $sne $roku $amzn $disca $lgf $twtr $fox $fb $goog $snap $stx $mgm $quibi $prime $hulu $pluto $tubi $hbo $fb $baba $bidu $iq $twtr $fubo $ea $atvi $msft https://t.co/7zdG0FF6bF","author_id":"302674240","created_at":"2022-06-24T20:28:03Z","source":"Twitter for iPhone","sentiment_score":0.2,"sentiment_magnitude":0.2},{"id":"1540431539057688576","text":"RT @unusual_ape: Kenny G is being very bad introducing Amy into Russell 1000. \n\nIs this how you treat Amy? By shorting her right off bat? B…","author_id":"1422580634854346752","created_at":"2022-06-24T20:27:53Z","source":"Twitter for Android","sentiment_score":-0.4,"sentiment_magnitude":1.8},{"id":"1540431538407743488","text":"Stock\nPositions [06/24/2022] High/Low $AAPL $AMZN $TSLA $ISRG $URI $NVDA $TWTR $CVX $AMAT $HPE  Exited $OXY on Bounce. Trimmed $NVDA on Bounce. Cash 21.38% Watch List; $BABA $BSM $CRWD $DE $GPC $NUE $QCOM $QDEL $ROKU $SNOW $UPST $WMT Holding All Positions. Maintain Cash \u0026gt; 15%.","author_id":"2897340758","created_at":"2022-06-24T20:27:53Z","source":"Twitter for iPhone","sentiment_score":0.1,"sentiment_magnitude":0.7},{"id":"1540431533466845184","text":"$AAPL held $140 high open interest level all day long and then pushed to higher open interest level to $142 at close https://t.co/Ss9mgIsBdc","author_id":"235892590","created_at":"2022-06-24T20:27:52Z","source":"Twitter Web App","sentiment_score":-0.3,"sentiment_magnitude":0.3},{"id":"1540431397906776064","text":"RT @chema_rodriguez: Y así cerramos la semana con $AAPL.\n\nSe suponía que este SPREAD era para un swing que tome en los últimos minutos para…","author_id":"847768752","created_at":"2022-06-24T20:27:19Z","source":"Twitter for Android","sentiment_score":0,"sentiment_magnitude":0.3},{"id":"1540431395536965634","text":"RT @HighStrikeInc: 6/24 Signals Recap💰\n\n$AAPL calls hit out of the gate and returned over 22%\n$F calls returned over 70%\n\n$NFLX was not tra…","author_id":"1285352866715242496","created_at":"2022-06-24T20:27:19Z","source":"Twitter for Android","sentiment_score":-0.2,"sentiment_magnitude":1.2},{"id":"1540431380575690752","text":"HBO Max Is Doing Well — So Of Course Changes Are Coming  $para $cmcsa $dis $nflx $aapl $t $sne $roku $amzn $disca $lgf $twtr $fox $fb $goog $snap $stx $mgm $quibi $prime $hulu $pluto $tubi $hbo $fb $baba $bidu $iq $twtr $fubo $ea $atvi $msft https://t.co/ZdaQ7zESEh","author_id":"302674240","created_at":"2022-06-24T20:27:15Z","source":"Twitter for iPhone","sentiment_score":0,"sentiment_magnitude":0},{"id":"1540431222073135105","text":"I see JPM has loaded calls $SPX $SPY $QQQ $TSLA $AAPL $MSFT https://t.co/Jrh2ZpwFvF","author_id":"795333944933752833","created_at":"2022-06-24T20:26:37Z","source":"Twitter for iPhone","sentiment_score":0,"sentiment_magnitude":0},{"id":"1540431016610959360","text":"$AAPL Def on watch https://t.co/oYPlTmsXyq","author_id":"1304655846979272705","created_at":"2022-06-24T20:25:48Z","source":"Twitter Web App","sentiment_score":0,"sentiment_magnitude":0}]}`
	body := twitter.GetDummyResult(query)
	c.JSON(200, body)
}

func FetchTwitterWithGCP(c *gin.Context) {
	query := c.Param("query")
	body := twitter.GetTwitterResult(query)
	for i := range body.Data {
		gcpResponse := gcp.GetSentimentGCP(body.Data[i].Text)
		gcpResponseScore := gcpResponse.DocumentSentiment.Score
		gcpResponseMagnitude := gcpResponse.DocumentSentiment.Magnitude
		body.Data[i].SentimentScore = gcpResponseScore
		body.Data[i].SentimentMagnitude = gcpResponseMagnitude
	}
	c.JSON(200, body)
}
