package main

import (
	//"encoding/json"
	//"fmt"
	//"bytes"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"strings"

	//"strings"

	//"strings"

	//"io/ioutil"
	//"os"

	//"io/ioutil"
	"log"
	//"os"
	"github.com/cdipaolo/sentiment"
	secret "github.com/jjavar1/yt-sentiment-analysis-web/back-end"

	//"io/ioutil"

	"net/http"

	// "net/url"
	// "os"
	// "os/user"
	// "path/filepath"
	"github.com/grassmudhorses/vader-go/lexicon"
    "github.com/grassmudhorses/vader-go/sentitext"
	"google.golang.org/api/option"
	youtube2 "google.golang.org/api/youtube/v3"
)

//Create get struct from front-end
type ytRequest struct {
	Video_Id string `string:"video_ID"`
}

//Create post struct for front-end
type sentScore struct {
	SentScore int
}

var positiveSentimentML int

var negativeSentimentML int

var positiveSentimentLexi int

var negativeSentimentLexi int

var positiveSentimentAverage int

var negativeSentimentAverage int

//Initialize vue serve
func main() {
	//Handle Get Request
	http.HandleFunc("/api/yt", Data_Handler)
	fs := http.FileServer(http.Dir("./sentiment-app/dist"))
	http.Handle("/", fs)
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

//Decode url from front end
func Data_Handler(w http.ResponseWriter, r *http.Request) {
	//Decode json by using struct
	var data ytRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	stringName := data.Video_Id
	Create_API_Request(stringName, secret.GetValue())
}

func Push_Comment_Results() {

}
//Create API request to youtube to retrieve comments using secret token and video url
func Create_API_Request(video_Id string, token string) {
	//create empty context
	ctx := context.Background()
	//create server
	create_serve, err := youtube2.NewService(ctx, option.WithAPIKey(token))
	if err != nil {
		panic(err.Error)
	}
	get_response := create_serve.CommentThreads.List([]string{"snippet"}).MaxResults(5).VideoId(video_Id)
	response, err := get_response.Do()
	if err != nil {
		panic(err.Error())
	}
	var all_decoded []string
	items := response.Items
	for _, item := range items {
		item_info := item.Snippet
		topLevelComment := item_info.TopLevelComment
		comment_info := topLevelComment.Snippet
		sr := comment_info.TextDisplay
		decoded := html.UnescapeString(sr)
		all_decoded = append(all_decoded, decoded, "\n")
	}
	res := strings.Join(all_decoded, "")
	fmt.Println(res)
	ML_Approach(res)
	Rule_Based_Approach(res)
}

//Machine learning approach to comments based on sentiment repo and training data
func ML_Approach(comments string) {
	
	scanner := bufio.NewScanner(strings.NewReader(comments))
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}
	var analysis *sentiment.Analysis

	//Iterate through the comments and incriment positive or negative score
	for scanner.Scan() {
		line := scanner.Text()
		analysis = model.SentimentAnalysis(line, sentiment.English)
		var sentiment string
		if analysis.Score == 1 {
			sentiment = "positive"
			positiveSentimentML++
		} else {
			sentiment = "negative"
			negativeSentimentML++
		}
		fmt.Printf("Review: %s \n and Sentiment:%s\n", line, sentiment)
	}
	fmt.Println(positiveSentimentML, negativeSentimentML)
}

func Rule_Based_Approach(comments string) {
	scanner := bufio.NewScanner(strings.NewReader(comments))
	
	for scanner.Scan() {
		line := scanner.Text()
		parsedtext := sentitext.Parse(line, lexicon.DefaultLexicon)
		sentiment := sentitext.PolarityScore(parsedtext)
		if (sentiment.Compound > 0) {
			positiveSentimentLexi++
		} else {
			negativeSentimentLexi++
		}
		fmt.Println("Review:", sentiment.Compound)
	}
	fmt.Println(positiveSentimentLexi, negativeSentimentLexi)
	Compute_Average()
}

//Takes average of ML approach and Lexi approach to reach a consensus
func Compute_Average() {
	positiveSentimentAverage = (positiveSentimentML + positiveSentimentLexi) / 2
	negativeSentimentAverage = (negativeSentimentML + negativeSentimentLexi) / 2
	fmt.Println(positiveSentimentAverage, negativeSentimentAverage)
}