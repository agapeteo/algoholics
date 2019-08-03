package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"nika/files"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var currentFolder string

const version = "2019-05-15-B"
const digit = "[0-9]+"

var isSilent bool
var baseTmpl = new(baseTemplate)
var digitPattern = regexp.MustCompile(digit)

type baseTemplate struct {
	baseTemplate *template.Template
	templates    map[string]string
}

func fromTemplate(w io.Writer, path string) {

	contentData, ok := baseTmpl.templates[path]
	if !ok {
		log.Panicf("can't find content for path %s", path)
	}

	tmpl, err := baseTmpl.baseTemplate.Clone()
	if err != nil {
		log.Panicf("can't create clone from base template: %v", err)
	}

	template.Must(tmpl.Parse(contentData))

	executionErr := tmpl.ExecuteTemplate(w, "base", path)
	if executionErr != nil {
		panic(executionErr)
	}
}

func prune(str string) string {
	var result string
	result = strings.Replace(str, "-", " ", -1)
	result = strings.Replace(result, "_", "", -1)
	result = digitPattern.ReplaceAllString(result, "")
	return result
}

func prepareTemplates() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		prepareLayouts()

		var err error
		baseTmpl.templates, err = files.ReadFiles(filepath.Join(currentFolder, "templates", "content"))
		if err != nil {
			panic(err)
		}

		wg.Done()
	}()

	wg.Wait()
	logIt("web site prepared.")
}

func prepareLayouts() {
	layouts, err := files.ReadFiles(filepath.Join(currentFolder, "templates", "layout"))
	if err != nil {
		panic(err)
	}

	baseTmpl.baseTemplate = template.New("main")

	baseTmpl.baseTemplate.Funcs(
		template.FuncMap{
			"prune": prune,
		})

	for l := range layouts {
		if l == "projects_content.tmpl.html" {
			continue
		}
		template.Must(baseTmpl.baseTemplate.Parse(layouts[l]))
		logItf("layout %v parsed", l)
	}
}

func validateCurrentFolder(folder string) {
	info, err := os.Stat(folder)
	if err != nil {
		log.Panic(err)
	}

	if !info.IsDir() {
		log.Panic(folder + " is not directory")
	}

	info, err = os.Stat(folder + "/static")
	if err != nil {
		log.Panicf("can't find static folder at %s, error: %v", folder+"/static", err)
	}

	if !info.IsDir() {
		log.Panic(folder + " is not directory")
	}
}

func serve(port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		logItf("requested %v ", path)

		if path == "/" {
			path = "index.html"
		}
		fromTemplate(writer, strings.Replace(path, "/", "", 1))
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(currentFolder+"/static/"))))

	logItf("web server is ready on port %v", port)
	panic(http.ListenAndServe(port, nil))
}

func generate() {

	targetFolder := "docs"
	targetFolderPath := filepath.Join(currentFolder, targetFolder)

	err := os.RemoveAll(targetFolderPath)

	err = os.Mkdir(targetFolderPath, 0755)
	if err != nil {
		logItf("can't create dir %v, error: %v \n", targetFolderPath, err)
	}

	err = files.CopyDir(
		filepath.Join(currentFolder, "static"),
		filepath.Join(currentFolder, targetFolder, "static"))

	if err != nil {
		logItf("can't copy folders: %v \n", err)
	}

	tmplFiles, filesErr := ioutil.ReadDir(filepath.Join(currentFolder, "templates", "content"))
	if filesErr != nil {
		logItf("can't read tmplFiles in directory: %v \n", filesErr)
	}

	var wg sync.WaitGroup
	for _, file := range tmplFiles {
		wg.Add(1)
		go func(f, t string) {
			copyFile(f, t)
			wg.Done()
		}(file.Name(), targetFolder)
	}
	wg.Wait()

	copyCnameFile(targetFolder)

	logIt("Completed.")
}

func copyCnameFile(targetFolder string) {
	cFile := filepath.Join(currentFolder, "CNAME")
	cBytes, err := ioutil.ReadFile(cFile)
	if err != nil {
		log.Panicf("can't read CNAME file at path %v  error: %v", cFile, err)
	}

	fileErr := ioutil.WriteFile(filepath.Join(currentFolder, targetFolder, "CNAME"), cBytes, 0755)
	if fileErr != nil {
		log.Panicf("can't create file: %v", fileErr)
	}
	logIt("CNAME file copied")
}

func copyFile(fileName, targetFolder string) {
	pageBuf := bytes.NewBufferString("")
	fromTemplate(pageBuf, fileName)

	fileErr := ioutil.WriteFile(filepath.Join(currentFolder, targetFolder, fileName), pageBuf.Bytes(), 0755)
	if fileErr != nil {
		logItf("can't create file: %v", fileErr)
	}

	logItf("%v -> ready \n", fileName)
}

func logItf(f string, v ...interface{}) {
	logIt(fmt.Sprintf(f, v))
}

func logIt(msg string) {
	if isSilent {
		return
	}
	log.Print(msg)
}

func main() {
	isGenerate := flag.Bool("generate", false, "generate file, otherwise run local web server")
	port := flag.Int("port", 8080, "web port")
	folder := flag.String("folder", "", "source folder")
	flag.BoolVar(&isSilent, "s", false, "enable silent mode")

	flag.Parse()

	logItf("version: %s", version)

	if *folder == "" {
		var err error
		currentFolder, err = files.LocalFolder()
		if err != nil {
			log.Panicf("can't find local folder: %v", err)
		}
	} else {
		currentFolder = *folder
	}
	validateCurrentFolder(currentFolder)
	logItf("using folder: %v", currentFolder)

	prepareTemplates()

	if *isGenerate {
		generate()
		return
	}

	serve(":" + strconv.Itoa(*port))
}
