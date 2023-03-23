package main

import (
	"fmt"
	"sync"
)

func main() {
	var arrNum []int32

	arrNum = append(arrNum, 1)
	arrNum = append(arrNum, 7)
	arrNum = append(arrNum, 2)
	arrNum = append(arrNum, 4)

	ch1 := make(chan bool)
	ch2 := make(chan []int32)
	var wg sync.WaitGroup

	wg.Add(1)
	go devide(arrNum, ch1, ch2, &wg)

	var isDone bool = <-ch1
	var arrNum1 []int32 = <-ch2

	for {
		if isDone {
			break
		}
		wg.Add(1)
		go devide(arrNum1, ch1, ch2, &wg)
	}
	wg.Wait()
	fmt.Println(arrNum1)
}

func devide(arrNum []int32, cha1 chan bool, cha2 chan []int32, wg *sync.WaitGroup) {
	defer wg.Done()
	var isDone bool = true
	var arrNum1 []int32
	arrNum1 = append(arrNum1, arrNum...)

	for i := 0; i < len(arrNum1); i++ {
		arrNum1[i] = 0
	}

	for j := 0; j < len(arrNum1); j++ {
		for k := j + 1; k < len(arrNum1); k++ {
			if !(arrNum[j]%3 == 0 && arrNum[k]%3 == 0) {
				sum := arrNum[j] + arrNum[k]
				if sum%3 == 0 {
					isDone = false
					arrNum1[j] = arrNum1[j] + 1
					arrNum1[k] = arrNum1[k] + 1
				}
			}
		}
	}

	highest := 0
	indexHighest := 0
	for i := 0; i < len(arrNum1); i++ {
		if highest < int(arrNum1[i]) {
			highest = int(arrNum1[i])
			indexHighest = i
		}
	}

	cha1 <- isDone
	cha2 <- append(arrNum[:indexHighest], arrNum[indexHighest+1:]...)
}
