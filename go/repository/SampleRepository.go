// BasicRepository.go
package repository

import (
	"fmt"
	"goProfile/go/config"
)

func JustSayDbProperties(){
	
	propertySource := config.GetPropertySource()
	fmt.Println("repository : "+propertySource.GetProfile())
}
