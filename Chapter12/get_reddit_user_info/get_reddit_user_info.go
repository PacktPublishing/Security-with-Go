package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Define the structure of the JSON response
// The json variable names are specified on
// the right since they do not match the
// struct variable names exactly
type redditUserJsonResponse struct {
	Data struct {
		Posts []struct { // Posts & comments
			Data struct {
				Subreddit  string  `json:"subreddit"`
				Title      string  `json:"link_title"`
				PostedTime float32 `json:"created_utc"`
				Body       string  `json:"body"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func printUsage() {
	fmt.Println(os.Args[0] + ` - Print recent Reddit posts by a user

  Usage: ` + os.Args[0] + ` <username>
  Example: ` + os.Args[0] + ` nanodano
`)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}
	url := "https://www.reddit.com/user/" + os.Args[1] + ".json"

	// Make HTTP request and read response
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request. ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP response body. ", err)
	}

	// Decode response into data struct
	var redditUserInfo redditUserJsonResponse
	err = json.Unmarshal(body, &redditUserInfo)
	if err != nil {
		log.Fatal("Error parson JSON. ", err)
	}

	if len(redditUserInfo.Data.Posts) == 0 {
		fmt.Println("No posts found.")
		fmt.Printf("Response Body: %s\n", body)
	}

	// Iterate through all posts found
	for _, post := range redditUserInfo.Data.Posts {
		fmt.Println("Subreddit:", post.Data.Subreddit)
		fmt.Println("Title:", post.Data.Title)
		fmt.Println("Posted:", time.Unix(int64(post.Data.PostedTime), 0))
		fmt.Println("Body:", post.Data.Body)
		fmt.Println("========================================")
	}
}
