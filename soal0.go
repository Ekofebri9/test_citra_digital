package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(elements []string) []string {
	hurufSama := map[string]bool{}
	hasil := []string{}
	for i := range elements {
		if hurufSama[elements[i]] == false {
			hurufSama[elements[i]] = true
			hasil = append(hasil, elements[i])
		}
	}
	return hasil
}

func main() {
	inputan := "omama"
	arr := strings.Split(inputan, "")
	vocal := "aAiIuUeEoO"
	hurufHidup := []string{}
	hurufMati := []string{}
	for i := 0; i < len(arr); i++ {
		if strings.Contains(vocal, arr[i]) {
			hurufHidup = append(hurufHidup, arr[i])
		} else {
			hurufMati = append(hurufMati, arr[i])
		}
	}
	hurufHidup = removeDuplicates(hurufHidup)
	fmt.Println("Huruf Hidup : ", len(hurufHidup))
	fmt.Println("Huruf Mati : ", len(hurufMati))
}
