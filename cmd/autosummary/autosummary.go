package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var SideBar string

const readme = "readme.md"
const output  = ""

// @author yugj
func main() {

	var base = "/Users/yugj/mdev/yugj/blog"

	summary := manager{}
	ignoredFiles := summary.ignoredFiles()

	dirs, err := ioutil.ReadDir(base)
	if err != nil {
		log.Fatal(err)
	}

	SideBar += "- BLOG\n"

	for _, dir := range dirs {

		_, ok := ignoredFiles[dir.Name()]
		if ok {
			continue
		}

		firstCategory := summary.buildCategory(base, dir.Name(), readme)
		summary.appendLine(" ", firstCategory)

		children, err := ioutil.ReadDir(base + "/" + dir.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, child := range children {

			if child.Name() == readme {
				continue
			}
			secondCategory := summary.buildCategory(base, dir.Name(), child.Name())
			summary.appendLine("    ",secondCategory)
		}
	}

	fmt.Println(SideBar)
}

type manager struct {

}

func (*manager) buildCategory(basePath, dirname, filename string) string{
	readme := basePath + "/" + dirname + "/" + filename
	title := getMdTitle(readme)

	if title == "" {
		title = dirname
	}

	return "- [" + title + "]" + "(/" + dirname + "/" + filename +   ")"
}

func getMdTitle(filePath string) string {

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}

	fileContent := string(fileBytes)
	firstLine := strings.Split(fileContent, "\n")[0]

	return strings.ReplaceAll(firstLine, "# ", "")
}

func (*manager) appendLine(tab string, line string) {
	SideBar += tab + line + "\n"
}

func (*manager) ignoredFiles() map[string]string {
	ignoredFiles := make(map[string]string)
	ignoredFiles[".git"] = "1"
	ignoredFiles[".gitignore"] = "1"
	ignoredFiles[".idea"] = "1"
	ignoredFiles["readme.md"] = "1"
	ignoredFiles["_404.md"] = "1"
	ignoredFiles["_coverpage.md"] = "1"
	ignoredFiles["_sidebar.md"] = "1"
	ignoredFiles["index.html"] = "1"
	return ignoredFiles
}
