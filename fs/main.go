package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLine(num int) string {
	s := ""
	f, e := os.Open("../fonts/standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)

	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		/*
			The above code reads a line from the file and stores it in the string s.
		*/
		s, _ = content.ReadString('\n')
	}

	s = strings.TrimSuffix(s, "\n")
	return s
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main (s string) string {
	if len(os.Args) < 2 {
		return
	}

	font := "standard.txt"
	text := ""

	// 1. Gathering info (input and flags)
	for _, v := range os.Args[1:] {
		if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			text += v
		}
	}

	var sarray []string
	for _, v := range s {
		sarray = append(sarray, string(v))
	}

	for i, _ := range sarray {
		// if i != 0 {
		if i == 32 {
			sarray = remove(sarray, i-1)
			i = i - 1
		}
		// }
	} fmt.Println(s)
}

