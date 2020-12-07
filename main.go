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
	html := read("./empty.html")
	folder := `./`

	files, _ := ioutil.ReadDir(folder)

	sss := ""
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := file.Name()
			reg, _ := regexp.Compile("(\\d{8})(.+)\\.txt")
			if reg.MatchString(name) {
				time := name[:8]
				title := name[8 : len(name)-4]
				time = newtime(time)

				content := read(name)
				if content != "" {

					part := strings.Split(content, "\n")

					parts := strings.Join(part, "</p>\n<p>")

					s := []string{"<article>\n", `<div class="title">`, title, "</div>\n", "<p>"}

					s = append(s, parts[:len(parts)-3])

					s = append(s, `<div class="time">`+time+"</div>\n</article>\n")
					ss := strings.Join(s, "")

					sss = ss + sss

				}
			}
		}
	}

	html = html + sss
	fmt.Println(html)

	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(html)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "字节写入成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func read(name string) string {
	f, err := ioutil.ReadFile("./" + name)
	if err != nil {
		fmt.Println(err)
	}
	return string(f)
}

func newtime(time string) string {
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
