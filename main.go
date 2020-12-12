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
	//如果要修改主题，直接修改 empty.html 就行
	emptyHTML := read("./empty.html")
	//所有文件都在同一级目录下
	files, _ := ioutil.ReadDir(`./`)
	sss := ""
	//匹配以 20201212 这样日期开头的 txt 文件
	reg, _ := regexp.Compile("(\\d{8})(.+)\\.txt")
	for number, file := range files {
		fmt.Println(number)
		name := file.Name()
		if (!file.IsDir()) && (reg.MatchString(name)) {
			content := read(name)
			//内容为空的跳过
			if content == "" {
				continue
			}
			//日期就是文件名的前八位
			time := chartime(name[:8])
			//标题就是日期之后 .txt之前的那一部分
			title := name[8 : len(name)-4]
			//在读取的 txt 最后加上换行，方便下一步按照换行分割 txt
			part := strings.Split(content+"\n", "\n")
			//再把分割来的片段用 <p> 组合起来
			//在这里可以看出，我只对换行做了变换，所以 HTML 的部分标签在不换行的情况下是可以用的
			//当然这个标签会被 <p> 包裹，但是一般不影响使用
			//没有内容的空 <p> 标签不会单独占一空行，会默认无视
			parts := strings.Join(part, "</p>\n<p>")
			//再加点修饰，把时间和标题加上
			article := "<article>\n" + `<div class="title">` + title + "</div>\n<p>" + parts[:len(parts)-3] + `<div class="time">` + time + "</div>\n</article>\n"
			//这一步是调整顺序（for range 是按照从小到大排的，不调整的话，新的日期就会排在后面）
			sss = article + sss
		}
	}
	//大功告成
	indexHTML := emptyHTML + sss
	//创建 index.html
	f, err := os.Create("index.html")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	//写入 index.html
	_, err = f.WriteString(indexHTML)
	if err != nil {
		fmt.Println(err)
	}
}

//这个函数是读取文件名，返回文件内容
func read(name string) string {
	f, err := ioutil.ReadFile("./" + name)
	if err != nil {
		fmt.Println(err)
	}
	return string(f)
}

//这个函数是用来转换日期的，效果见下
//20201212 -> 二〇二〇年十二月十二日
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
