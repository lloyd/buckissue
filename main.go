package main

import (
       "os"
       "fmt"
       "io/ioutil"
)

func main() {
     fmt.Printf("Hello World!\n")

     // we expect to be running via `buck run` in a generated directory
     dir,err := os.Getwd()
     if err != nil {
         panic("oops")
     }
     fmt.Printf("We're in %s\n", dir)
     dat, err := ioutil.ReadFile("generated_asset")
     if err != nil {
         fmt.Fprintf(os.Stdout, "Can't find my asset: %s\n", err)
	 os.Exit(1)
     }
     fmt.Printf("My Asset: %s\n", dat)
}
