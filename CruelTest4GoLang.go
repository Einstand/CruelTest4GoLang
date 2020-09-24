package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return wd
}

func findAllGoFiles(path string) ([]string, error) {
	var matches []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			var sName = file.Name()
			var sExt = filepath.Ext(sName)
			if sExt == ".go" {
				var sNameNoExt = sName[0 : len(sName)-3]
				matches = append(matches, sNameNoExt)
			}
		}
	}
	return matches, nil
}

func split_test_and_not(source []string) ([]string, []string) {
	var aryNotTest []string
	var aryIsTest []string
	for _, sName := range source {
		if len(sName) > 5 {

			var sRight5 = sName[len(sName)-5:]
			if sRight5 == "_test" {
				aryIsTest = append(aryIsTest, sName)
			} else {
				aryNotTest = append(aryNotTest, sName)
			}
		} else {
			aryNotTest = append(aryNotTest, sName)
		}
	}
	sort.Strings(aryNotTest)
	sort.Strings(aryIsTest)
	return aryNotTest, aryIsTest
}

func check_test_exist(aryNotTest []string, aryIsTest []string) bool {
	var bOK = true
	for _, sName := range aryNotTest {
		var sTestName = sName + "_test"
		var nArrayLen = len(aryIsTest)
		nFind := sort.SearchStrings(aryIsTest, sTestName)
		if nFind >= nArrayLen {
			fmt.Println(sTestName + ".go not exist")
			bOK = false
		}
	}
	return bOK
}

func main() {
	fmt.Println("Cruel Test for GoLang")
	wd := getWorkingDirectory()
	fmt.Println("Working Directory is " + wd)

	files, err := findAllGoFiles(wd)
	if err != nil {
		log.Fatal(err)
	}

	aryNotTest, aryIsTest := split_test_and_not(files)

	if check_test_exist(aryNotTest, aryIsTest) {
		fmt.Println(aryNotTest)
		fmt.Println(aryIsTest)
	} else {
		os.Exit(1) // return non zero when error
	}

	//files, err := ioutil.ReadDir(wd)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, file := range files {
	//	if !file.IsDir() {
	//		var sName = file.Name()
	//		var sExt = filepath.Ext(sName)
	//		if sExt == ".go" {
	//			var sNameNoExt = sName[0 : len(sName)-3]
	//			fmt.Println(sNameNoExt)
	//		}
	//	}
	//}
}
