package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("go collins 命令行版本")
		fmt.Println("传入第一个参数为要查询的值 eg: hello")
		fmt.Println("author : whisperjie")
		return
	}
	keyword := os.Args[1]
	dom, err := goquery.NewDocument("http://dict.youdao.com/w/" + keyword + "/#keyfrom=dict2.top")
	if err != nil {
		log.Fatalln(err)
	}
	//#collinsResult > div > div > div > div
	dom.Find("#collinsResult > div > div > div > div > ul").Each(func(i int, selection *goquery.Selection) {
		//#collinsResult > div > div > div > div > ul > li:nth-child(1) > div.collinsMajorTrans > p
		str := selection.Text()
		list := strings.FieldsFunc(str, unicode.IsSpace)
		//reg := regexp.MustCompile(`\d\.`)
		//size := len(list)
		//var listMeaning=[size]string
		s := ""
		for i := range list {
			if strings.Index(list[i], ".") == 1 {
				s = s + "\n" + list[i] + " "
			} else {
				s = s + list[i] + " "
			}
		}
		listMeaning := strings.Split(s, "\n")
		for i := range listMeaning {
			fmt.Println(listMeaning[i])
		}
	})
}
