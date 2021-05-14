package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"archive/zip"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./zipped [filename] [zipname] [relative path]")
		return
	}

	fileName := os.Args[1]
	zipName := os.Args[2]
	filePath := os.Args[3]

	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	newZip, err := os.Create(zipName)
	if err != nil {
		log.Fatal(err)
	}
	defer newZip.Close()

	archive := zip.NewWriter(newZip)

	f, err := archive.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	_, err = f.Write([]byte(fileBytes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[*] Saved %s to path \"%s\" in %s\n", fileName, filePath, zipName)
}
