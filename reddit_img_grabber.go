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

 
func main() {
 urls := geturls("hot")
 for _, url := range urls {
    saveimage(url)
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

func saveimage(url string) {
  name := urltoname(url)
  response_img, _ := http.Get(url)
  img, _, _ := image.Decode(response_img.Body)
  if img != nil {
    file, _ := os.Create("/home/stephane/imgs/" + name + ".png")
    err := png.Encode(file, img)
    if err != nil {
      fmt.Println(err)
    }
    file.Close()
  }
}
