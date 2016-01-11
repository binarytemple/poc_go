package main

import "fmt"
import "io/ioutil"

func main() {
  
  files, _ := ioutil.ReadDir("./tests")
  for _, f := range files {
    var fileName = f.Name() 
    if( fileName[len(fileName)-7:] == ".leases" ) {
        fmt.Println(fileName)
    }
  }
}
