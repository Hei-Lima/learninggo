package main

import (
	"golang.org/x/tour/wc";
	"strings";
)

func WordCount(s string) map[string]int {
	var mapa = make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		mapa[word]++;
	}
	
	return mapa
}

func main() {
	wc.Test(WordCount)
}
