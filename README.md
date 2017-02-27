# tt
[![Go Report Card](https://goreportcard.com/badge/github.com/adolphlwq/tt)](https://goreportcard.com/report/github.com/adolphlwq/tt)

A tool for generating template Markdown file for [translate](https://github.com/adolphlwq/translate)

## Usage
### install
If you have installed [Golang](https://golang.org)
```
go install github.com/adolphlwq/tt
```
else
```
wget https://github.com/adolphlwq/tt/releases/download/0.1/tt
sudo chmod +x tt
sudo mv tt /usr/local/bin
```

### Usage
- see version
    ```
    translate git:(master) tt version 
    tt version: 0.1
    go version: go1.7.4
    os/arch:    linux/amd64
    ```
- new a template post with default arguments
    ```
    translate git:(master) tt new
    file "00-example.md" created successfully on "/home/geek/workspace/githubwp/translate/2017"!
    ```
- new a template post with customed arguments
  ```
  ➜  translate git:(master) ✗ tt new -f adolphlwq-tt -n 10 -y 2017 -t "dive into python"
  file "10-adolphlwq-tt.md" created successfully on "/home/geek/workspace/githubwp/translate/2017"!
  ➜  translate git:(master) ✗ cat 2017/10-adolphlwq-tt.md 
  # dive into python

  ## 说明
  - [原文链接]()
  - [翻译：@adolphlwq](https://github.com/adolphlwq)
  - [项目地址](https://github.com/adolphlwq/translate)
  - <a rel="license" href="http://creativecommons.org/licenses/by-nc/4.0/"><img alt="知识共享许可协议" style="border-width:0" src="https://i.creativecommons.org/l/by-nc/4.0/80x15.png" />
  ```

### Options
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

## TODOs
- [X] new: new a post
- [X] version: show version info