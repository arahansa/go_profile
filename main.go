// main.go
package main

import (
	"fmt"
	"flag"
	"goProfile/go/config"
	"goProfile/go/repository"
)



func main(){
	propertySource := config.GetPropertySource()
	
	fmt.Println("Program Start!")
	wordPtr := flag.String("p", "test", "a string")
	flag.Parse()
	
	propertySource.SetProfile(*wordPtr)
	fmt.Println("Get flag is :", propertySource.GetProfile())
	repository.JustSayDbProperties()
	fmt.Println("Program End!")
}



















