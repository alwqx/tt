package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"strings"

	"github.com/spf13/cobra"
)

var year, number, fileName, title string

var newCmd = &cobra.Command{
	Use:   "new [flags]",
	Short: "create a new template file.",
	Run:   newRun,
}

func newRun(cmd *cobra.Command, args []string) {
	absDir, err := os.Getwd()
	if err != nil {
		qerr(err)
	}

	aimDir := filepath.Join(absDir, year)
	if !exists(aimDir) {
		if err := os.Mkdir(aimDir, 0755); err != nil {
			qerr(err)
		}
	}

	aimFileName := formatFileName(number, fileName)
	if exists(filepath.Join(aimDir, aimFileName)) {
		qerr(fmt.Sprintf("file %s is exits in %s", aimFileName, aimDir))
	}

	createPostFile(aimDir, aimFileName)
}

func qerr(input interface{}) {
	fmt.Println(input)
	os.Exit(-1)
}

//Check if a file or directory exists.
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	qerr(err)
	return false
}

func init() {
	newCmd.Flags().StringVarP(&year, "year", "y", "2017", "year of directory contains posts")
	newCmd.Flags().StringVarP(&number, "number", "n", "00", "number of your post")
	newCmd.Flags().StringVarP(&fileName, "file-name", "f", "example", "file name of your post")
	newCmd.Flags().StringVarP(&title, "title", "t", "example", "title of your post")
	RootCmd.AddCommand(newCmd)
}

func createPostFile(dir, fileName string) {
	tmpl := `# {{.title}}
	
## 说明
- [原文链接]()
- [翻译：@adolphlwq](https://github.com/adolphlwq)
- [项目地址](https://github.com/adolphlwq/translate)
- <a rel="license" href="http://creativecommons.org/licenses/by-nc/4.0/"><img alt="知识共享许可协议" style="border-width:0" src="https://i.creativecommons.org/l/by-nc/4.0/80x15.png" />

`
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["title"] = title

	err := writeTemplateToFile(filepath.Join(dir, fileName), tmpl, data)

	if err != nil {
		qerr(err)
	}

	fmt.Println(fmt.Sprintf("file \"%s\" created successfully on \"%s\"!", fileName, dir))
}

func writeTemplateToFile(filePath string, tmpl string, data interface{}) error {
	r, err := templateToReader(tmpl, data)
	if err != nil {
		return nil
	}

	err = saveReaderToDisk(filePath, r)
	if err != nil {
		return err
	}
	return nil
}

func templateToReader(tmpl string, data interface{}) (io.Reader, error) {
	t, err := template.New("tt-tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)

	return buf, err
}

func saveReaderToDisk(filePath string, r io.Reader) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func formatFileName(number, fileNname string) string {
	if strings.HasSuffix(fileName, ".md") {
		return number + "-" + fileName
	}
	return number + "-" + fileName + ".md"
}
