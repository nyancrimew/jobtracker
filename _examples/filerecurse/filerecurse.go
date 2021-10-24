package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deletescape/jobtracker"
)

func main() {
	jt := jobtracker.NewJobTracker(RecurseWorker, 60, jobtracker.DefaultNapper)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			jt.AddJob(f.Name())
		} else {
			fmt.Println(f.Name())
		}
	}
	abs, _ := filepath.Abs(".")
	jt.StartAndWait(RecurseContext{BaseDir: abs})
}
