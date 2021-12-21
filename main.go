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
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func generateComponent(name string) {
	const originFile string = "./templates/component.tpl"
	valueMap := make(map[string]string)
	valueMap["name"] = name
	res := formartTpl(originFile, valueMap)
	outName := "./src/components/" + name + ".js"
	err2 := ioutil.WriteFile(outName, []byte(res), 0666)
	check(err2)
	fmt.Printf("生成%v组件成功\n", name)
}
func generateView(name string) {
	originFile := "./templates/page.tpl"
	valueMap := make(map[string]string)
	valueMap["name"] = name
	res := formartTpl(originFile, valueMap)
	// 创建文件夹
	outDir := "./src/views/" + strings.ToLower(name)
	exis, err := PathExists(outDir)
	if !exis {
		err = os.Mkdir(outDir, 0666)
		check(err)
	}
	//创建组件
	outName := outDir + "/" + name + ".js"
	err = ioutil.WriteFile(outName, []byte(res), 0666)
	check(err)
	// 创建css
	outName = outDir + "/index.css"
	err = ioutil.WriteFile(outName, []byte(""), 0666)
	check(err)
	//创建index.js
	originFile = "./templates/index.tpl"
	res = formartTpl(originFile, valueMap)
	outName = outDir + "/index.js"
	err = ioutil.WriteFile(outName, []byte(res), 0666)
	check(err)
	fmt.Printf("生成%v页面成功\n", name)

}
func formartTpl(path string, valueMap map[string]string) string {
	data, err := ioutil.ReadFile(path)
	check(err)
	dataStr := string(data)
	var res string
	for key := range valueMap {
		res = strings.Replace(dataStr, "${"+key+"}", valueMap[key], -1)
	}
	return res
}
func main() {
	operator := "-cn"
	for k, v := range os.Args {
		if k == 0 {
			continue
		}
		if find := strings.Contains(v, "-"); find {
			operator = v
		} else {
			if operator == "-cn" {
				generateComponent(v)
			} else {
				generateView(v)
			}
		}

	}

}
