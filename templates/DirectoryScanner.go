package templates

import (
    "os"
    "path/filepath"
    "io/ioutil"
)

type DirectoryScanner interface {
    Templates() []Template
}

type DirectoryScannerImpl struct {
    RootPath  string
    templates []Template
}

func (s DirectoryScannerImpl) Templates() []Template {
    s.templates = make([]Template, 0)
    filepath.Walk(s.RootPath, s.scanFile)
    return s.templates
}

func (s *DirectoryScannerImpl) scanFile(path string, f os.FileInfo, err error) error {
    if f.IsDir() {
        return nil
    }

    fileContents, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }
    filename := f.Name()
    filename = filename[:len(filename) - 4]
    contents := string(fileContents)
    template := AssemblePage(filename, contents)
    s.templates = append(s.templates, template)
    return nil
}
