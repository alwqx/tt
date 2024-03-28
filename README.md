# tt

[![Go Report Card](https://goreportcard.com/badge/github.com/alwqx/tt)](https://goreportcard.com/report/github.com/alwqx/tt)

A tool for generating template Markdown file for [translate](https://github.com/alwqx/translate)

## Usage

### install

If you have installed [Golang](https://golang.org)

```
go install github.com/alwqx/tt@latest
```

else

```
wget https://github.com/alwqx/tt/releases/download/0.3.0/tt
sudo chmod +x tt
sudo mv tt /usr/local/bin
```

### Usage

- see version

  ```
  translate git:(master) tt version
  tt version: 0.3.0
  go version: go1.20.3
  os/arch:    darwin/arm64
  ```

- new a template post with default arguments

  ```
  translate git:(master) tt new
  file "00-example.md" created successfully on "/home/geek/workspace/githubwp/translate/2017"!
  ```

- new a template post with customed arguments

  ```
  ➜  translate git:(master) ✗ tt new -f alwqx-tt -n 10 -y 2017 -t "dive into python"
  file "10-alwqx-tt.md" created successfully on "/home/geek/workspace/githubwp/translate/2017"!
  ➜  translate git:(master) ✗ cat 2017/10-alwqx-tt.md
  # dive into python

  ## 说明
  - [原文链接]()
  - [翻译：@alwqx](https://github.com/alwqx)
  - [项目地址](https://github.com/alwqx/translate)
  - [tt](https://github.com/alwqx/tt)：自动生成翻译模板
  - 用时：
  - <a rel="license" href="http://creativecommons.org/licenses/by-nc/4.0/"><img alt="知识共享许可协议" style="border-width:0" src="https://i.creativecommons.org/l/by-nc/4.0/80x15.png" /></a>
  ```

### Options

- new a post

  ```
  create a new template file.

  Usage:
  tt new [flags]

  Flags:
  -f, --file-name string   file name of your post (default "example")
  -n, --number string      number of your post (default "00")
  -t, --title string       title of your post (default "example")
  -y, --year string        year of directory contains posts (default "2017")
  ```

- list issue (defalut alwqx/translate)

  ```
  Usage:
  tt issue [flags]

  Flags:
  -r, --repo string    name of your repo. (default "translate")
  -s, --state string   state of issue to list. (default "open")
  -u, --user string    short name of your github account. (default "alwqx")
  ```

## TODOs

- [x] new: new a post.
- [x] version: show version info.
- [x] issue: list issue.
