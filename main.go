package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("开始运行:")
	emptyHTML := read("./empty.html")
	files, _ := ioutil.ReadDir(`./`)
	sss := ""
	reg, _ := regexp.Compile("(\\d{8})(.+)\\.txt")
	for number, file := range files {
		fmt.Println(number)
		name := file.Name()
		if (!file.IsDir()) && (reg.MatchString(name)) {
			content := read(name)
			if content == "" {
				continue
			}
			time := chartime(name[:8])
			title := name[8 : len(name)-4]
			part := strings.Split(content+"\n", "\n")
			parts := strings.Join(part, "</p>\n<p>")
			article := "<article>\n" + `<div class="title">` + title + "</div>\n<p>" + parts[:len(parts)-3] + `<div class="time">` + time + "</div>\n</article>\n"
			sss = article + sss
		}
	}

	indexHTML := emptyHTML + sss

	f, err := os.Create("index.html")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(indexHTML)
	if err != nil {
		fmt.Println(err)
	}
}

func read(name string) string {
	f, err := ioutil.ReadFile("./" + name)
	if err != nil {
		fmt.Println(err)
	}
	return string(f)
}

func chartime(time string) string {
	time2 := ""
	// time 20200820
	// 		01234567
	for i := 0; i < len(time); i++ {
		if i == 4 {
			time2 = time2 + "年"
		}
		if i == 4 && time[i] == '0' {
			continue
		}
		if i == 4 && time[4] == '1' {
			time2 = time2 + "十"
			continue
		}
		if i == 6 {
			time2 = time2 + "月"
		}
		if i == 6 && (time[i] == '0' || time[i] == '1') {
			continue
		}
		if i == 7 && time[6] != '0' {
			time2 = time2 + "十"
		}
		if i == 7 && time[7] == '0' {
			continue
		}

		switch time[i] {
		case '0':
			time2 = time2 + "〇"
		case '1':
			time2 = time2 + "一"
		case '2':
			time2 = time2 + "二"
		case '3':
			time2 = time2 + "三"
		case '4':
			time2 = time2 + "四"
		case '5':
			time2 = time2 + "五"
		case '6':
			time2 = time2 + "六"
		case '7':
			time2 = time2 + "七"
		case '8':
			time2 = time2 + "八"
		case '9':
			time2 = time2 + "九"
		}
	}
	time2 = time2 + "日"
	return time2
}
