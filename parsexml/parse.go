package parsexml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/pengxianghu/tsungo/defs"
)

func parse(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file error ", err)
	}

	//将文件转换为对象
	var tsung defs.Tsung
	err = xml.Unmarshal(content, &tsung)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result:", tsung)

	fmt.Println("读取完成。。。。")
}
