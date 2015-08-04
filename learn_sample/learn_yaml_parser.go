//http://stackoverflow.com/questions/28682439/golang-parse-yaml-file
package main

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func main() {
	m := make(map[interface{}]interface{})
	filename, _ := filepath.Abs("./application.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	err = yaml.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
}



















