package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//downloadDict("words_alpha.txt", "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")

	allwords, err := dictToArray("words_alpha.txt")
	if err != nil {
		fmt.Errorf("%v", err)

	}
	kwords := filter(allwords, 'k')
	//fmt.Printf("%v", kwords)
	hwords := filter(allwords, 'h')

	rand.Seed(time.Now().UnixNano())
	kword := kwords[rand.Intn(len(kwords))]
	fmt.Println(kword)
	hword := hwords[rand.Intn(len(hwords))]
	fmt.Println(hword)
}

func filter(inputArray []string, initial byte) []string {

	var filtered []string
	for _, s := range inputArray {
		if s[0] == initial {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
