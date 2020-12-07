# 1a 极简单页博客生成器

如果您觉得不满意可以联系我，如果依然不满意或者喜欢markdown可以去使用Hugo、Jekyll、Hexo

不支持图片等高级格式，支持纯文本

这个主题来自 [寫嘢 - An beatutiful Hexo Theme](https://github.com/eatradish/Seje)

[DEMO](http://cccc.press)

# 使用

所有资源文件均在同一文件夹下

- `empty.html` 空白主题，注意修改 `html-title`
- `index.html` 生成的目标文件，上传到oss或者gh-page
- `20200827名称.txt` 博客内容，日期必须为8位，名称为 `index.html` 中的 `article-title` ，必须以 `.txt` 结尾
- `1a` 程序本体，双击或在命令行均可

# 主题

自行修改 `empty.html`

# 编译

```
go build main.go
```

# Todo

目前还有个bug，txt最后必须以回车结尾（也就是说最后要空一行），不然会吞字