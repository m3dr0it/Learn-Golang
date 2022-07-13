package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func proceed() {
	counterTotal := 0
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string,
		info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		buf, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		sum := fmt.Sprintf("%x", md5.Sum(buf))
		destintionPath := filepath.Join(tempPath, fmt.Sprintf("file ke-%s.txt", sum))
		errR := os.Rename(path, destintionPath)

		if errR != nil {
			return errR
		}

		counterRenamed++

		log.Println(sum)
		return nil

	})

	if err != nil {
		log.Println(err.Error())
	}
}
