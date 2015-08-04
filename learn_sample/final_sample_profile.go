// newfilereader.go
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var configmap map[interface{}]interface{}
var flagProfile string

func main() {
	SetProfile("test")
	var currentProfile string
	configmap = make(map[interface{}]interface{})
	fmt.Println("Hello Golang File Reader!")

	// 01. yml 파일을 읽어들여서 먼저 파싱을 한다.
	filename, _ := filepath.Abs("./application.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	checkErrorWithMsg(err, "can't not find application.yml")

	fmt.Println("unmarshall yaml file ==== ")
	err = yaml.Unmarshal(yamlFile, &configmap)
	checkErrorWithMsg(err, "unmarshalling error")
	fmt.Println("=== map ===\n ", configmap)

	if flagProfile != "" {
		currentProfile = flagProfile
	} else {
		var ok bool
		activeProfile := configmap["go.profiles.active"]
		currentProfile, ok = activeProfile.(string)
		if !ok {
			checkErrorWithMsg(nil, "active profile is not string")
		}
	}
	// 02. go.profiles 혹은 active된 프로파일을 읽는다.
	// 해당 부분이 없는 경우도 생각을 해야 한다. 일단은 이대로 구현을 하겠다.
	if currentProfile != "" {
		filecontens := string(yamlFile)
		//		fmt.Println("\n==== contents ==== \n", filecontens)
		searchWorkd := strings.Join([]string{"go.profiles: ", currentProfile}, "")
		fmt.Println("===================\n")
		fmt.Println("search word is :", searchWorkd)

		// 03. ---로 구분을 해서 알맞는 환경변수를 찾아보자.
		slices := strings.Split(filecontens, "---")
		for k, v := range slices {
			fmt.Println(k, " and ", v)
			if strings.Contains(v, searchWorkd) {
				// 04. 구분을 하여서 그 부분을 잘라서 yaml 마샬링에 추가
				fmt.Println("find :", v)
				err = yaml.Unmarshal([]byte(v), &configmap)
				checkErrorWithMsg(err, "unmarshalling error")
				fmt.Println("=== last map ===\n ", configmap)
				break
			}
			if k == len(slices)-1 {
				log.Fatal("can not find your profile")
			}
		}

	}
	// 05. map을 내부에 가지고 있고, 외부에서 호출을 하는데, 키를 가지지 않은 경우 fatal 에러처리를 해서
	// 초반에 프로그램이 기동이 되면서 잘못된 키를 불러왔을 때는 기동이 안되게 한다.

	fmt.Println(Get("go.datasource"))
	fmt.Println(Get("go.robot"))
	fmt.Println("end...")
}
func SetProfile(profile string) {
	flagProfile = profile
}
func Get(key string) interface{} {
	value := configmap[key]
	if value == nil {
		log.Fatal(key, "is invalid key")
	}
	return value
}

func checkErrorWithMsg(err error, msg string) {
	if err != nil {
		log.Fatal(msg, "\n", err)
	}
}
