package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var baseUrl = "https://api.namefake.com/"
var basePath = "Random User"
var baseFilteredPath = "Filtered"

type UserPayload struct {
	Email    string
	Password string
}

type RandName struct {
	Number     int    `json:"number"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	BirthData  string `json:"birth_data"`
	Generation string `json:"gen"`
	FilePath   string `json:"filepath"`
	Md5Sum     string `json:"md5sum"`
}

func main() {

	// os.RemoveAll(basePath)
	// os.Mkdir(basePath, os.ModePerm)
	// for i := 0; i < 500; i++ {
	// 	writeToFile(fetchRandName())
	// }

	start := time.Now()
	// moveAllFiles()

	recIt1 := readFiles()
	recIt2 := filterGen(recIt1)
	recIt3 := filterGen(recIt1)
	recIt4 := filterGen(recIt1)

	chanMerged := mergeChan(recIt2, recIt3, recIt4)

	recGenMd51 := generateMd5(chanMerged)
	recGenMd52 := generateMd5(chanMerged)
	recGenMd53 := generateMd5(chanMerged)

	chanGenMd5Sum := mergeChan(recGenMd51, recGenMd52, recGenMd53)

	for user := range chanGenMd5Sum {
		log.Println(user.Number)
	}

	end := time.Since(start)
	// filterGen(<-chainOut)

	time.Sleep(1 * time.Second)
	log.Println(end)

}

func readFiles() <-chan RandName {
	sendIt := make(chan RandName)

	go func() {
		defer close(sendIt)
		index := 0
		pathToRead := filepath.Join(os.Getenv("TEMP"), basePath)
		err := filepath.Walk(pathToRead, func(path string, info fs.FileInfo, err error) error {

			if info.IsDir() {
				return err
			}

			buff, errReadFile := ioutil.ReadFile(path)

			if errReadFile != nil {
				return errReadFile
			}

			var user RandName
			errDecoding := json.Unmarshal(buff, &user)

			if errDecoding != nil {
				return errDecoding
			}

			user.Number = index
			sendIt <- user
			index++
			return nil
		})

		if err != nil {
			log.Println(err.Error())
		}
	}()

	return sendIt
}

func generateMd5(randNameCh <-chan RandName) <-chan RandName {
	chanOut := make(chan RandName)
	go func() {
		defer close(chanOut)
		for randName := range randNameCh {
			content, err := ioutil.ReadFile(randName.FilePath)
			if err != nil {
				log.Println(err.Error())
			}

			md5sum := md5.Sum(content)
			randName.Md5Sum = fmt.Sprintf("%x", md5sum)
			writeToFile(randName)
			chanOut <- randName

		}
	}()

	return chanOut
}

func sendSomething() <-chan RandName {
	chanOut := make(chan RandName)

	go func() {
		defer close(chanOut)
		for i := 0; i < 10; i++ {
			user := RandName{
				Name:    "Ahmad Mardiana",
				Address: fmt.Sprint(i),
			}
			chanOut <- user
		}

	}()

	return chanOut
}

func receiveFromChan(chanIn <-chan RandName) {
	// chanOut := make(<-chan RandName)
	go func() {
		for getUser := range chanIn {
			log.Println(getUser.Name)
			log.Println(getUser.Address)
		}

	}()

}

func fetchRandName() RandName {

	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var randName1 RandName
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&randName1)
	randName1.FilePath = basePath + "/" + randName1.Name + ".json"
	return randName1
}

func writeToFile(randName RandName) {
	randNameFile, _ := json.MarshalIndent(randName, "", "  ")
	errWrtie := os.WriteFile(randName.FilePath, randNameFile, os.ModePerm)

	if errWrtie != nil {
		log.Println(errWrtie.Error())
	}
}

func filterGen(randNameCh <-chan RandName) <-chan RandName {
	chanOut := make(chan RandName)

	go func() {
		defer close(chanOut)
		for randName := range randNameCh {
			birthDate, err := time.Parse("2006-01-02", randName.BirthData)
			oldPath := basePath + "/" + randName.Name + ".json"

			if err != nil {
				log.Println(err.Error())
			}

			switch true {
			case isGenX(birthDate):
				// err := os.Mkdir(baseFilteredPath+"/Gen X", os.ModePerm)
				if err != nil {
					log.Println(err.Error())
				}
				randName.Generation = "Gen X"
				randName.FilePath = baseFilteredPath + "/" + randName.Generation + "/" + randName.Name + ".json"
				os.Rename(oldPath, randName.FilePath)
			case isGenMillenials(birthDate):
				// err := os.Mkdir(baseFilteredPath+"/Gen Millenial", os.ModePerm)
				if err != nil {
					log.Println(err.Error())
				}
				randName.Generation = "Gen Millenial"
				randName.FilePath = baseFilteredPath + "/" + randName.Generation + "/" + randName.Name + ".json"
				os.Rename(oldPath, randName.FilePath)
			case isGenZ(birthDate):
				// err := os.Mkdir(baseFilteredPath+"/Gen Z", os.ModePerm)
				if err != nil {
					log.Println(err.Error())
				}
				randName.Generation = "Gen Z"
				randName.FilePath = baseFilteredPath + "/" + randName.Generation + "/" + randName.Name + ".json"
				os.Rename(oldPath, randName.FilePath)

			default:
				// err := os.Mkdir(baseFilteredPath+"/No Generation Defined", os.ModePerm)
				if err != nil {
					log.Println(err.Error())
				}
				randName.Generation = "No Generation Defined"
				randName.FilePath = baseFilteredPath + "/" + randName.Generation + "/" + randName.Name + ".json"
				os.Rename(oldPath, randName.FilePath)

			}
			writeToFile(randName)
			os.Remove(oldPath)

			chanOut <- randName

		}
	}()

	return chanOut
}

func moveAllFiles() {
	err := filepath.Walk(baseFilteredPath+"/No Generation Defined", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		buff, errReadFile := ioutil.ReadFile(path)

		if errReadFile != nil {
			return errReadFile
		}

		var user RandName
		errDecoding := json.Unmarshal(buff, &user)

		if errDecoding != nil {
			return errDecoding
		}

		// chan1 <- user

		os.Rename(user.FilePath, "Random User/"+user.Name+".json")
		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}

	err1 := filepath.Walk(baseFilteredPath+"/Gen X", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		buff, errReadFile := ioutil.ReadFile(path)

		if errReadFile != nil {
			return errReadFile
		}

		var user RandName
		errDecoding := json.Unmarshal(buff, &user)

		if errDecoding != nil {
			return errDecoding
		}

		// chan1 <- user

		os.Rename(user.FilePath, "Random User/"+user.Name+".json")
		return nil
	})

	if err1 != nil {
		log.Println(err.Error())
	}

	err2 := filepath.Walk(baseFilteredPath+"/Gen Z", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		buff, errReadFile := ioutil.ReadFile(path)

		if errReadFile != nil {
			return errReadFile
		}

		var user RandName
		errDecoding := json.Unmarshal(buff, &user)

		if errDecoding != nil {
			return errDecoding
		}

		// chan1 <- user

		os.Rename(user.FilePath, "Random User/"+user.Name+".json")
		return nil
	})

	if err2 != nil {
		log.Println(err.Error())
	}

	err3 := filepath.Walk(baseFilteredPath+"/Gen Millenial", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		buff, errReadFile := ioutil.ReadFile(path)

		if errReadFile != nil {
			return errReadFile
		}

		var user RandName
		errDecoding := json.Unmarshal(buff, &user)

		if errDecoding != nil {
			return errDecoding
		}

		// chan1 <- user

		os.Rename(user.FilePath, "Random User/"+user.Name+".json")
		return nil
	})

	if err3 != nil {
		log.Println(err.Error())
	}

}

func mergeChan(randNameMany ...<-chan RandName) <-chan RandName {
	wg := new(sync.WaitGroup)
	wg.Add(len(randNameMany))
	chanOut := make(chan RandName)

	for _, eachChan := range randNameMany {
		go func(eachChan <-chan RandName) {
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

func isGenX(birhtDate time.Time) bool {
	if birhtDate.Year() > 1965 && birhtDate.Year() <= 1980 {
		return true
	}
	return false
}

func isGenMillenials(birhtDate time.Time) bool {
	if birhtDate.Year() > 1980 && birhtDate.Year() <= 1996 {
		return true
	}
	return false
}

func isGenZ(birthDate time.Time) bool {
	if birthDate.Year() > 1996 && birthDate.Year() <= 2012 {
		return true
	}
	return false
}
