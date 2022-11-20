package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// if os.Args[1] == "" {
	// 	return
	// }
	if len(os.Args) == 1 {
		AsciArt("hello")
		AsciArt("HELLO")
		AsciArt("HeLlo HuMaN")
		AsciArt("1Hello 2There")
		AsciArt("Hello\\nThere")
		AsciArt("Hello\\n\\nThere")
		AsciArt("{Hello & There #}")
		AsciArt("hello There 1 to 2\\!")
		AsciArt("MaD3IrA&LiSboN")
		AsciArt("1a\"#FdwHywR&/()=")
		AsciArt("{|}~")
		AsciArt("[\\]^_ 'a")
		AsciArt("RGB")
		AsciArt(":;<=>?@")
		AsciArt("\\!\" #$%&'()*+,-./")
		AsciArt("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		AsciArt("abcdefghijklmnopqrstuvwxyz")
		return
	}
	str := os.Args[1]
	for _, v := range os.Args[2:] {
		str += " " + v
	}
	AsciArt(str)
}

func AsciArt(s string) string {

	f, e := os.Create("output_test.txt")
	if e != nil {
		os.Exit(0)
	}
	res := ""

	severallines := false

	var sarray []string
	for _, v := range s {
		sarray = append(sarray, string(v))
	}

	end := len(sarray) - 1

	for i, _ := range sarray {
		if i != 0 {
			if i == end && sarray[i] == "!" && sarray[i-1] == "\\" {
				sarray = remove(sarray, i-1)
				i = i - 1
			}
			if sarray[i] == "n" && sarray[i-1] == "\\" {
				// if i < 3 {
				// 	severallines = false
				// } else {
				severallines = true
				// }
			}
		}
	}

	s = strings.Join(sarray, "")

	if severallines {
		args := strings.Split(s, "\\n")
		single := len(args) - 1
		for _, word := range args {
			if single == 1 && word == "" {
				fmt.Println()
				return (".")
			}
			if word == "" {
				fmt.Println()
			} else {
				for i := 0; i < 8; i++ {
					for _, letter := range word {
						res += GetLine(2 + int(letter-' ')*9 + i)
					}
					fmt.Println(res)
					fmt.Fprintln(f, res)
					res = ""
				}
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range s {

				res += GetLine(2 + int(letter-' ')*9 + i)
			}
			fmt.Println(res)
			fmt.Fprintln(f, res)
			res = ""
		}
	}
	f.Close()
	return s
}

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
