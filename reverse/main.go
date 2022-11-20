package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	usage := `Usage: go run . [OPTION]
	
EX: go run . --reverse=<fileName>`

	if len(os.Args) == 1 {
		fmt.Println(usage)
		// fmt.Println("Usage: go run . [OPTION]")
		// fmt.Println("EX: go run . --reverse=<fileName>")
		return
	}

	template := "fonts/standard.txt"
	filename := ""
	args := os.Args[1:]
	for _, v := range args {
		if len(v) > 10 && v[:10] == "--reverse=" {
			filename = v[10:]
		} else {
			fmt.Println("Usage: go run . [OPTION]")
			fmt.Println("EX: go run . --reverse=<fileName>")
			return
		}
	}

	f, e := os.Open(template)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()
	content := bufio.NewReader(f)

	f.Seek(0, 0)
	filename = "output_test.txt"
	j, e := os.Open("audit/" + filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}

	defer j.Close()

	j.Seek(0, 0)
	inputcontent := bufio.NewReader(j)

	var dstr []string
	/*
		It creates a 2D array of strings.
	*/

	inputbase := make([][]string, 8)

	database := ""
	input := ""

	num := 1000
	for i := 0; i < num; i++ {
		database, _ = content.ReadString('\n')
		dstr = append(dstr, database)
	}
	b := 0

	for i := 0; i < 8; i++ {
		input, _ = inputcontent.ReadString('\n')
		a := len(input)
		b = a
		inputbase[i] = make([]string, a)
		for j, v := range input {
			inputbase[i][j] = string(v)
		}
	}

	memory := 0
	result := ""

	for j := 0; j < b; j++ {
		i := 0
		if (inputbase[i][j] == " ") && (inputbase[i+1][j] == " ") && (inputbase[i+2][j] == " ") && (inputbase[i+3][j] == " ") && (inputbase[i+4][j] == " ") && (inputbase[i+5][j] == " ") && (inputbase[i+6][j] == " ") && (inputbase[i+7][j] == " ") {
			var letter [][]string

			letter = arraymaker(j+1, inputbase, memory)

			memory = j + 1

			match := 0
			line0 := ""
			line1 := ""
			line2 := ""
			line3 := ""
			line4 := ""
			line5 := ""
			line6 := ""
			line7 := ""

			for ln, arr := range letter {
				for _, v := range arr {
					if ln == 0 {
						line0 = line0 + string(v)
					}
					if ln == 1 {
						line1 = line1 + string(v)
					}
					if ln == 2 {
						line2 = line2 + string(v)
					}
					if ln == 3 {
						line3 = line3 + string(v)
					}
					if ln == 4 {
						line4 = line4 + string(v)
					}
					if ln == 5 {
						line5 = line5 + string(v)
					}
					if ln == 6 {
						line6 = line6 + string(v)
					}
					if ln == 7 {
						line7 = line7 + string(v)
					}
				}
			}

			for index := 0; index < num; index++ {
				if (strings.TrimRight(dstr[index], "\n") == line0) && (strings.TrimRight(dstr[index+1], "\n") == line1) && (strings.TrimRight(dstr[index+2], "\n") == line2) && (strings.TrimRight(dstr[index+3], "\n") == line3) && (strings.TrimRight(dstr[index+4], "\n") == line4) && (strings.TrimRight(dstr[index+5], "\n") == line5) && (strings.TrimRight(dstr[index+6], "\n") == line6) {
					match = index
					// fmt.Println(match)
					// break
				}
			}

			result = result + string(((match-1)/9)+32)

			line0 = ""
			line1 = ""
			line2 = ""
			line3 = ""
			line4 = ""
			line5 = ""
			line6 = ""
			line7 = ""

			continue

		}

	}
	resultarray := split_white_spaces(result)
	str := ""
	for _, word := range resultarray {
		str = str + word + " "
	}
	str = remove_spaces(str)
	fmt.Println(str)

}

func arraymaker(y int, s [][]string, x int) [][]string {
	var result = make([][]string, 8)
	for i := 0; i < 8; i++ {
		result[i] = make([]string, y)
		for j := x; j < y; j++ {
			result[i][j] = s[i][j]
		}
	}
	return result
}

func lenfinder(s []string, t int) int {
	return len(s[t])
}

func remove_spaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return remove_spaces(s[:len])
	}
	return s[:len]
}

func split_white_spaces(s string) []string {
	var str []string
	var word string
	l := len(s) - 1

	for i, v := range s {
		if i == l {
			word = word + string(v)
			str = append(str, word)
		} else if v == 32 || v == 15 || v == 10 {
			if s[i+1] == ' ' || s[i+1] == '	' || s[i+1] == 10 {
			} else {
				str = append(str, word)
				word = ""
			}
		} else {
			word = word + string(v)
		}
	}
	return str
}
