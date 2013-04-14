// go run rar_renamer.go /path_to_dir/ old new 
package main

import (
  "fmt"
  "io/ioutil" 
  "strings"
  "os"
  ) 

func main() {
  root := os.Args[1]
  to_be_replaced := os.Args[2]
  replace_by := os.Args[3]
  dir, err := ioutil.ReadDir(root)
  if err != nil { 
    fmt.Println("something went wrong")
  }
  for _, file := range dir {
    old_name := file.Name()
    new_name := strings.Replace(old_name, to_be_replaced, replace_by, 1)
    fmt.Println("Renaming " + old_name + " to " + new_name)
    os.Rename(root + old_name, root + new_name)
  }
}
