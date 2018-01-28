package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 { // If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}

	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}

	fileList.PushBack(fileNode)
}

func getFilesInDirRecursivelyBySize(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
	}

	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			getFilesInDirRecursivelyBySize(
				fileList,
				filepath.Join(path, dirFile.Name()),
			)
		} else if dirFile.Mode().IsRegular() {
			insertSorted(
				fileList,
				FileNode{FullPath: fullpath, Info: dirFile},
			)
		}
	}
}

func main() {
	fileList := list.New()
	getFilesInDirRecursivelyBySize(fileList, "/home")

	for element := fileList.Front(); element != nil; element =
		element.Next() {
		fmt.Printf("%d ", element.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
