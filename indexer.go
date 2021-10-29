package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

type Document struct {
	ID   string
	Freq int64
}

func FindFreq(words []string) *map[string]int64 {
	freq := map[string]int64{}
	for _, i := range words {
		freq[i]++
	}
	return &freq
}

var (
	directory = flag.String("directory", "/temp", "Directory Bilgi Mesajı")
	query     = flag.String("query", "fox", "Query Bilgi Mesajı")
)

func main() {
	flag.Parse()
	indexer := map[string][]Document{}
	file, err := filepath.Glob(*directory + "/*")
	if err != nil {
		log.Fatal("Directory Bulanamadı", *directory)
	}
	for _, f := range file {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal("Files Bulanamadı")
		}
		words := strings.Split(string(content), " ")
		freqs := FindFreq(words)
		for word, number := range *freqs {
			if _, ok := indexer[word]; ok {
				indexer[word] = append(indexer[word], Document{ID: f, Freq: number})
			} else {
				arr := []Document{Document{ID: f, Freq: number}}
				indexer[word] = arr
			}
		}
	}
	result := indexer[*query]
	sort.Slice(result, func(i, j int) bool {
		return result[i].Freq > result[j].Freq
	})
	fmt.Println("Results for word query:", *query)
	for _, d := range result {
		fmt.Printf("%s =>%d\n", d.ID, d.Freq)
	}
}
