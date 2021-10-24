package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deletescape/jobtracker"
)

type RecurseContext struct {
	BaseDir string
}

func RecurseWorker(jt *jobtracker.JobTracker, path string, context jobtracker.Context) {
	c := context.(RecurseContext)

	files, err := ioutil.ReadDir(filepath.Join(c.BaseDir, path))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			jt.AddJob(filepath.Join(path, f.Name()))
		} else {
			fmt.Println(filepath.Join(path, f.Name()))
		}
	}
}
