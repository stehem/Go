package main
 
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

 

type JsonResponse struct {
  Kind string
  Data DataType
  //Databases []DatabasesType
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
  response, _ := http.Get("http://reddit.com/r/hot.json")
  defer response.Body.Close()
  contents, _ := ioutil.ReadAll(response.Body)
  var resp JsonResponse
  err := json.Unmarshal(contents, &resp)
  if err != nil {
    fmt.Println(err)
  }
  children := resp.Data.Children
  var urls []string 
  for _, val := range children {
    urls = append(urls, val.Data.Url)
  }
  fmt.Println(urls)
}
