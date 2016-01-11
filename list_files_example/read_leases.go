package main

import "fmt"
import "io/ioutil"

func main() {
  path := "./tests" 
  files, _ := ioutil.ReadDir(path)
  var all = ""
  for _, f := range files {
    var fileName = f.Name() 
    if( fileName[len(fileName)-7:] == ".leases" ) {
        content,_ := ioutil.ReadFile( path + "/" + fileName)   
        fmt.Println(fileName)
        fmt.Println(string(content))
        all = all + string(content)
    }
  }
  fmt.Println(all)
}
