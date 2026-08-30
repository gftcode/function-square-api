package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	path string
	fileName string
}

func NewFile() *File {
	fileName := "function².txt"
	folder := "../../"

	_ = os.MkdirAll(folder, os.ModePerm)

	completePath := filepath.Join(folder, fileName)
	create(completePath)

	return &File{path: completePath, fileName: fileName}
}


func (f *File) AddContent(contexts ...any) error {	
    file, err := f.open()	
    if err != nil {		
        return err	
    }	
    defer file.Close()	
    
    for _, context := range contexts {		
        _, err = fmt.Fprint(file, context)		
        if err != nil {			
            return err		
        }	
    }		
    return nil
}

func create(completePath string) (*os.File, error) {
	file, err := os.Create(completePath)
	if err != nil {
		return nil, err
	}	

	return file, nil
}


func (f *File) open() (*os.File, error) {
	file, err := os.OpenFile(f.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
