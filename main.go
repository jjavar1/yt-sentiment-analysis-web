package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/cdipaolo/sentiment"
	secret "github.com/jjavar1/yt-sentiment-analysis-web/back-end"

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
	SentScore 					int
	TotalPositiveComments 		int 
	TotalNegativeComments 		int 	
	PositiveMLComments 			int	
	NegativeMLComments 			int
	PositiveLexiComments 		int
	NegativeLexiComments 		int
	MLComments 					[]string
	LexComments					[]string
}

var positiveSentimentML int

var negativeSentimentML int

var positiveSentimentLexi int

var negativeSentimentLexi int

var positiveSentimentAverage int

var negativeSentimentAverage int

var totalPositive int

var totalNegative int

var totalAverage int

var mlComments []string

var lexiComments []string

//Initialize vue serve
func main() {
	//Handle Get Request
	
	http.HandleFunc("/api/yt", Data_Handler)
	
	http.HandleFunc("/api/yt/get", Send_Post_Request)
	
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
	//reset global values 
	//change this garbage to a struct later please
	
	stringName := data.Video_Id
	Create_API_Request(stringName, secret.GetValue())
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
	get_response := create_serve.CommentThreads.List([]string{"snippet"}).MaxResults(40).VideoId(video_Id)
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

	ML_Approach(res)
	Rule_Based_Approach(res)
	Compute_Average()
	
}

//Machine learning approach to comments based on sentiment repo and training data
func ML_Approach(comments string) {
	
	scanner := bufio.NewScanner(strings.NewReader(comments))
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}
	var analysis *sentiment.Analysis
	s := []string{}
	//Iterate through the comments and incriment positive or negative score
	for scanner.Scan() {
		line := scanner.Text()
		analysis = model.SentimentAnalysis(line, sentiment.English)
		
		if analysis.Score == 1 {
			positiveSentimentML++
		} else {
			negativeSentimentML++
		}
		s = append(s, fmt.Sprintf("Comment: %s \n Score:%d\n", line, analysis.Score))
	}
	mlComments = s
	fmt.Println(positiveSentimentML, negativeSentimentML)
}

//Lexicon based approach to get sentiment consensus
func Rule_Based_Approach(comments string) {
	scanner := bufio.NewScanner(strings.NewReader(comments))
	
	s := []string{}
	//Scan through comments and incriment sentiment
	for scanner.Scan() {
		line := scanner.Text()
		parsedtext := sentitext.Parse(line, lexicon.DefaultLexicon)
		sentiment := sentitext.PolarityScore(parsedtext)
		if (sentiment.Compound > 0) {
			positiveSentimentLexi++
		} else {
			negativeSentimentLexi++
		}
		s = append(s, fmt.Sprintf("Comment: %s \n Score:%f\n", line, sentiment.Compound))
	}
	lexiComments = s
	fmt.Println(positiveSentimentLexi, negativeSentimentLexi)
}

//Takes average of ML approach and Lexi approach to reach a consensus
//-1 negative sentiment, 1 positive sentiment
func Compute_Average() {
	positiveSentimentAverage = (positiveSentimentML + positiveSentimentLexi) / 2
	negativeSentimentAverage = (negativeSentimentML + negativeSentimentLexi) / 2
	totalPositive = (positiveSentimentML + positiveSentimentLexi)
	totalNegative = (negativeSentimentML + negativeSentimentLexi)
	if (positiveSentimentAverage > negativeSentimentAverage) {
		totalAverage = 1
	} else {
		totalAverage = -1
	}
	fmt.Printf("total average:%d",  totalAverage)
}

func Send_Post_Request(w http.ResponseWriter, r *http.Request) {
	
	
	struct_data := sentScore{SentScore: totalAverage,
	TotalPositiveComments: totalPositive,
	TotalNegativeComments: totalNegative,
	PositiveMLComments: positiveSentimentML,
	NegativeMLComments: negativeSentimentML,
	PositiveLexiComments: positiveSentimentLexi,
	NegativeLexiComments: negativeSentimentLexi,
	MLComments: mlComments,
	LexComments: lexiComments}

	json_data, err := json.Marshal(&struct_data)
	if err != nil {
		fmt.Println(err)
		return
	}
	positiveSentimentML = 0
	negativeSentimentML = 0
	positiveSentimentLexi = 0
	negativeSentimentLexi = 0
	positiveSentimentAverage = 0
	negativeSentimentAverage = 0
	fmt.Fprintf(w, `{ "sentScore": %s }`, json_data)
}