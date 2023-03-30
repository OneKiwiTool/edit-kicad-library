package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func FileReplaceText(path string, old string, new string) error {

	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	newContents := strings.Replace(string(read), old, new, -1)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	path := "/home/vanson/onekiwi-kicad-library/footprint.pretty/"
	old := "solder_mask_margin"
	new := "solder_paste_margin"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		pathfile := path + file.Name()
		FileReplaceText(pathfile, old, new)
	}
}
