//https://github.com/go-yaml/yaml
package main

import (
	"fmt"
	

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

var data2 = `
c: Difficult?
d:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		C int
		D []int ",flow"
	}
}

func main() {
	// We need from here!!! (omit error )
	m := make(map[interface{}]interface{})
	
	yaml.Unmarshal([]byte(data), &m)
	
	fmt.Printf("--- m:\n%v\n\n", m)

	yaml.Unmarshal([]byte(data2), &m)
	
	d, _ := yaml.Marshal(&m)
	
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}






















