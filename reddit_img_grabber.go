// concurrently download images from any subreddit

package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "image"
  "image/png"
  _ "image/jpeg"
  "os"
  "strings"
)


type JsonResponse struct {
  Kind string
  Data DataType
}

type DataType struct {
  Children []ChildrenType
}

type ChildrenType struct {
  Data DataType2
}

type DataType2 struct {
  Url string
}


type RedditImage struct {
  url string
  name string
  http_resp *http.Response
}


func main() {
  var subreddit directory string
  fmt.Print("Subreddit: ")
  fmt.Scanln(&subreddit)
  fmt.Print("Path where to save: ")
  fmt.Scanln(&directory)
  grab(subreddit, directory)
}

func grab(sub string, dir string) {
  ch := make(chan RedditImage)
  urls := geturls(sub)

  for _, url := range urls {
    reddit := RedditImage{url: url, name: urltoname(url)}
    go fetchimage(reddit, ch)
  }

  for {
    select {
      case redditimg := <-ch:
        saveimage(redditimg, dir)
    }
  }
}

func geturls(subreddit string) []string  {
  response, _ := http.Get("http://reddit.com/r/" + subreddit + ".json")
  defer response.Body.Close()
  contents, _ := ioutil.ReadAll(response.Body)
  var resp JsonResponse
  err := json.Unmarshal(contents, &resp)
  if err != nil {
    fmt.Println("shit happened")
  }
  children := resp.Data.Children
  var urls []string
  for _, val := range children {
    urls = append(urls, val.Data.Url)
  }
  return urls
}

func urltoname(url string) string {
  split := strings.Split(url, "/")
  lenofsplit := len(split)
  indexoflast := lenofsplit - 1
  namewithext := split[indexoflast]
  split2 := strings.Split(namewithext, ".")
  name := split2[0]
  return name
}

func fetchimage(redditimg RedditImage, ch chan RedditImage) {
  resp, _ := http.Get(redditimg.url)
  redditimg.http_resp = resp
  ch <- redditimg
}

func saveimage(redditimg RedditImage, dir string) {
  name := redditimg.name
  img, _, _ := image.Decode(redditimg.http_resp.Body)
  if img != nil {
    file, _ := os.Create(dir + name + ".png")
    err := png.Encode(file, img)
    if err != nil {
      fmt.Println(err)
    }
    file.Close()
  }
}



