package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func generateComponent(name string) {
	originFile := "./templates/component.tpl"
	data, err := ioutil.ReadFile(originFile)
	check(err)
	dataStr := string(data)
	res := strings.Replace(dataStr, "${name}", name, -1)
	outName := "./src/components/" + name + ".js"
	err2 := ioutil.WriteFile(outName, []byte(res), 0666)
	check(err2)
	fmt.Printf("生成%v组件成功", name)
}
func generateView(name string) {
	originFile := "./templates/component.tpl"
	data, err := ioutil.ReadFile(originFile)
	check(err)
	dataStr := string(data)
	res := strings.Replace(dataStr, "${name}", name, -1)
	outName := "./src/views/" + name + ".js"
	err2 := ioutil.WriteFile(outName, []byte(res), 0666)
	check(err2)
	fmt.Printf("生成%v组件成功", name)
}
func main() {
	fileName := ""
	isComponent := false
	for k, v := range os.Args {
		if v == "-vn" {
			fileName = os.Args[k+1]
			break
		}
		if v == "-cn" {
			fileName = os.Args[k+1]
			isComponent = true
			break
		}
	}
	if isComponent {
		generateComponent(fileName)
	} else {
		generateView(fileName)
	}

}
