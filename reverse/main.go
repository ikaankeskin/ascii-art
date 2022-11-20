package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// if len(os.Args) == 1 {
	// 	return
	// }

	f, e := os.Open("../fonts/standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()
	content := bufio.NewReader(f)

	f.Seek(0, 0)

	j, e := os.Open("alphabet.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}

	defer j.Close()

	j.Seek(0, 0)
	inputcontent := bufio.NewReader(j)

	// scanner := bufio.NewScanner(inputcontent)
	// scan or read the bytes of text line by line
	// for scanner.Scan() {
	// 	fmt.Println(scanner)

	// 	// }
	// 	// println(match)
	// }

	// fmt.Println(scanner.Text())
	var dstr []string
	/*
		It creates a 2D array of strings.
	*/

	inputbase := make([][]string, 8)

	database := ""
	input := ""

	// fmt.Println(input)
	num := 1000
	for i := 0; i < num; i++ {
		database, _ = content.ReadString('\n')
		dstr = append(dstr, database)
	}
	b := 0
	// input, _ = inputcontent.ReadString('\n')
	// input = strings.TrimSuffix(input, "\n")
	// a := len(input)
	// fmt.Println(a)

	for i := 0; i < 8; i++ {
		input, _ = inputcontent.ReadString('\n')
		a := len(input)
		b = a
		inputbase[i] = make([]string, a)
		for j, v := range input {
			inputbase[i][j] = string(v)
		}
	}

	// fmt.Println(inputbase[5][6])
	// fmt.Println(inputbase[5][7])
	memory := 0

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
			// fmt.Println(line0)
			// fmt.Println(line1)
			// fmt.Println(line2)
			// fmt.Println(line3)
			// fmt.Println(line4)
			// fmt.Println(line5)
			// fmt.Println(line6)
			// fmt.Println(line7)
			for index := 0; index < num; index++ {
				if (strings.TrimRight(dstr[index], "\n") == line0) && (strings.TrimRight(dstr[index+1], "\n") == line1) && (strings.TrimRight(dstr[index+2], "\n") == line2) && (strings.TrimRight(dstr[index+3], "\n") == line3) && (strings.TrimRight(dstr[index+4], "\n") == line4) && (strings.TrimRight(dstr[index+5], "\n") == line5) {
					match = index
					// fmt.Println(match)
					// break
				}
			}

			fmt.Printf(string(((match - 1) / 9) + 32))
			line0 = ""
			line1 = ""
			line2 = ""
			line3 = ""
			line4 = ""
			line5 = ""
			line6 = ""
			line7 = ""
			// match = 0

			continue

		}

	}
	fmt.Println()

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

// fmt.Println(dstr)

// for i, v := range dstr {
// 	fmt.Println(i)
// 	fmt.Println(v)
// }

// fmt.Println(strings.Contains(dstr, input))

// database = strings.TrimSuffix(database, "\n")

// s = strings.TrimSuffix(s, "\n")
// if s == "" {
// 	continue
// }
// for iindex := 0; iindex <= 8; iindex++ {
// 	input, _ = inputcontent.ReadString('\n')
// input = strings.TrimSuffix(input, "\n")
// 	// if input == "" {
// 	// 	continue
// 	// }
// 	// if len(input) != len(s) {
// 	// 	continue
// 	// }
// 	if s == input && s != "" {
// 		fmt.Println(input)
// 	}

// 	// for _, v := range input {
// 	// 	fmt.Printf(string(v))
// 	// }
// 	// fmt.Println()
// 	// input = ""
// }

// if s == input {
// 	// fmt.Println(len(s) + len(input))
// 	// fmt.Println(len(input))
// 	// fmt.Println("there is something")

// 	fmt.Printf("matched:")
// 	fmt.Println(s)
// 	fmt.Printf("at:")
// 	fmt.Println(i)
// }
// fmt.Println(s)
// s = ""
// fmt.Printf(s)
// i = i + 8

// fmt.Println()
/*
	The above code reads a line from the file and stores it in the variable s.
*/
// s, _ = content.ReadString('\n')
// fmt.Println(s)
// s = strings.TrimSuffix(s, "\n")
// fmt.Println(s)

// func AsciArt(s string) string {

// 	f, e := os.Create("output_test.txt")
// 	if e != nil {
// 		os.Exit(0)
// 	}
// 	res := ""

// 	severallines := false

// 	var sarray []string
// 	for _, v := range s {
// 		sarray = append(sarray, string(v))
// 	}

// 	end := len(sarray) - 1

// 	for i, _ := range sarray {
// 		if i != 0 {
// 			if i == end && sarray[i] == "!" && sarray[i-1] == "\\" {
// 				sarray = remove(sarray, i-1)
// 				i = i - 1
// 			}
// 			if sarray[i] == "n" && sarray[i-1] == "\\" {
// 				// if i < 3 {
// 				// 	severallines = false
// 				// } else {
// 				severallines = true
// 				// }
// 			}
// 		}
// 	}

// 	s = strings.Join(sarray, "")

// 	if severallines {
// 		args := strings.Split(s, "\\n")
// 		single := len(args) - 1
// 		for _, word := range args {
// 			if single == 1 && word == "" {
// 				fmt.Println()
// 				return (".")
// 			}
// 			if word == "" {
// 				fmt.Println()
// 			} else {
// 				for i := 0; i < 8; i++ {
// 					for _, letter := range word {
// 						res += GetLine(2 + int(letter-' ')*9 + i)
// 					}
// 					fmt.Println(res)
// 					fmt.Fprintln(f, res)
// 					res = ""
// 				}
// 			}
// 		}

// 	} else {
// 		for i := 0; i < 8; i++ {
// 			for _, letter := range s {

// 				res += GetLine(2 + int(letter-' ')*9 + i)
// 			}
// 			fmt.Println(res)
// 			fmt.Fprintln(f, res)
// 			res = ""
// 		}
// 	}
// 	f.Close()
// 	return s
// }

// func GetLine(num int) string {
// 	s := ""
// 	f, e := os.Open("../fonts/standard.txt")
// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}
// 	defer f.Close()

// 	f.Seek(0, 0)

// 	content := bufio.NewReader(f)
// 	for i := 0; i < num; i++ {
// 		/*
// 			The above code reads a line from the file and stores it in the string s.
// 		*/
// 		s, _ = content.ReadString('\n')
// 	}

// 	s = strings.TrimSuffix(s, "\n")
// 	return s
// }

//	func remove(slice []string, s int) []string {
//		return append(slice[:s], slice[s+1:]...)
//	}
