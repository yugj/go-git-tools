package autosummary

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const readme = "readme.md"
const output = "_sidebar.md"

func Generate(base string) {

	summary := manager{}
	summary.basePath = base

	ignoredFiles := summary.ignoredFiles()

	containsTarget := summary.containsTarget(base)
	if !containsTarget {
		log.Println("target file not exists")
		return
	}

	dirs, _ := ioutil.ReadDir(base)

	summary.content += "- BLOG\n"

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
			summary.appendLine("    ", secondCategory)
		}
	}

	fmt.Println(summary.content)
	summary.writeTarget()

	log.Println("auto generate summary success")
}

type manager struct {
	content  string
	basePath string
}

func (s *manager) writeTarget() {

	sideBarPath := s.basePath + "/" + output
	file, _ := os.Create(sideBarPath)
	_, err := file.WriteString(s.content)
	if err != nil {
		log.Println(err)
		_ = file.Close()
	}

}

func (*manager) containsTarget(basePath string) bool {
	dirs, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.Name() == output {
			return true
		}
	}
	return false
}

func (*manager) buildCategory(basePath, dirname, filename string) string {
	readme := basePath + "/" + dirname + "/" + filename
	title := getMdTitle(readme)

	if title == "" {
		title = dirname
	}

	return "- [" + title + "]" + "(/" + dirname + "/" + filename + ")"
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

func (m *manager) appendLine(tab string, line string) {
	m.content += tab + line + "\n"
}

func (*manager) ignoredFiles() map[string]string {
	ignores := ".git,.gitignore,.idea,README.md,_404.md,_coverpage.md,_sidebar.md,index.html"
	ignoresArray := strings.Split(ignores, ",")

	ignoredFiles := make(map[string]string)
	for i := range ignoresArray {
		ignoredFiles[ignoresArray[i]] = "1"
	}

	return ignoredFiles
}
