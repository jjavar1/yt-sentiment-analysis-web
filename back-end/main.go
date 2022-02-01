package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

var api_key := "AIzaSyAr7qU0S_FAdk1aFu825DwYrDMYoUj13j4"

func get_comments(video_id string) {
	replies := string[]

	
}
