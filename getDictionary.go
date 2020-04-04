package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func downloadDict(dictName, dictAddress string) {
	resp, err := http.Get(dictAddress)
	if err != nil {
		fmt.Errorf("http get failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	err = ioutil.WriteFile(dictName, body, 0644)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func dictToArray(dictName string) ([]string, error) {
	var words []string
	file, err := os.Open(dictName)
	if err != nil {
		return words, fmt.Errorf("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}
