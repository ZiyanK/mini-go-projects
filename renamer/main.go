package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fileName := "birthday_001.txt"
	// newName, err := match(fileName)
	// if err != nil {
	// 	fmt.Println("no match")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir("./sample")
	if err != nil {
		panic(err)
	}
	count := 0
	type rename struct {
		fileName string
		path     string
	}
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	fmt.Println(count)
	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		fmt.Println("New file name", newFilename)
		newPath := filepath.Join(dir, newFilename)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mv %s => %s\n", origPath, newPath)
	}
	// origPath := fmt.Sprintf("%s/%s")
	// newPath := fmt.Sprintf("%s/%s")
}

//match return the new filename or an error
func match(fileName string, total int) (string, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s did not match our pattern", err)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
