package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 5000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "test-dir")

type FileInfo struct {
	FilePath  string
	Content   []byte
	Sum       string
	IsRenamed bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Print("start")
	start := time.Now()
	os.RemoveAll(tempPath)
	err := os.Mkdir(tempPath, os.ModePerm)

	// test := randomString(100)

	// log.Println(test)
	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file_ke- %d.txt", i))
		content := randomString(contentLength)
		errIo := ioutil.WriteFile(fileName, []byte(content), fs.ModePerm)
		if errIo != nil {
			log.Print(errIo.Error())
		}

		if i%100 == 0 && i > 1 {
			log.Println(i, "Files Created")
		}
	}

	if err != nil {
		log.Print(err.Error())
	}

	chanRec := readFiles()

	time.Sleep(50 * time.Second)

	testRecChan := getSum(chanRec)
	testRecChan1 := getSum(chanRec)
	testRecChan2 := getSum(chanRec)
	testRecChan3 := getSum(chanRec)
	testRecChan4 := getSum(chanRec)

	mergedChan := mergeChanFileInfo(testRecChan, testRecChan1, testRecChan2, testRecChan3, testRecChan4)

	chanRename1 := rename(mergedChan)
	chanRename2 := rename(mergedChan)
	chanRename3 := rename(mergedChan)
	chanrename4 := rename(mergedChan)

	mergedChanRen := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanrename4)
	// proceed()

	counterRenamed := 0
	counterTotal := 0

	for fileInfo := range mergedChanRen {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d %d files renamed", counterRenamed, counterTotal)
	log.Println(time.Since(start))

}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		defer close(chanOut)
		err := filepath.Walk(tempPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return err
			}

			buf, err := ioutil.ReadFile(path)

			if err != nil {
				log.Println(err.Error())
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil
		})

		if err != nil {
			log.Println("Error ", err.Error())
		}

	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanFileInfo(chainInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chainInMany))

	for _, eachChan := range chainInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}

		close(chanOut)
	}()

	return chanOut
}

// func generateFiles() {
// 	os.RemoveAll(tempPath)
// 	os.Mkdir(tempPath, os.ModePerm)

// 	for i := 0; i < totalFile; i++ {
// 		filename := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", i))
// 	}
// }
