// // package main

// /*
// It imports the testing package.
// */
// import (
// 	"fmt"
// 	"os"
// 	"testing"
// )

// func Test1(t *testing.T) {
// 	got, e := os.ReadFile("output_test.txt")
// 	// got := AsciArt("kaan")

// 	want, e := os.ReadFile("test1.txt")

// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}

// 	swant := string(want)
// 	sgot := string(got)
// 	fmt.Println(swant)

// 	if sgot != swant {
// 		t.Errorf(`got
// %v, wanted
// %v`, sgot, swant)
// 	}

// }

// func Test2(t *testing.T) {
// 	got, e := os.ReadFile("output_test.txt")
// 	// got := AsciArt("kaan")

// 	want, e := os.ReadFile("test2.txt")

// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}

// 	swant := string(want)
// 	sgot := string(got)
// 	fmt.Println(swant)

// 	if sgot != swant {
// 		t.Errorf(`got
// %v, wanted
// %v`, sgot, swant)
// 	}

// }

// func Test3(t *testing.T) {
// 	got, e := os.ReadFile("output_test.txt")
// 	// got := AsciArt("kaan")

// 	want, e := os.ReadFile("test3.txt")

// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}

// 	swant := string(want)
// 	sgot := string(got)
// 	fmt.Println(swant)

// 	if sgot != swant {
// 		t.Errorf(`got
// %v, wanted
// %v`, sgot, swant)
// 	}

// }

// func Test4(t *testing.T) {
// 	got, e := os.ReadFile("output_test.txt")
// 	// got := AsciArt("kaan")

// 	want, e := os.ReadFile("test4.txt")

// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}

// 	swant := string(want)
// 	sgot := string(got)
// 	fmt.Println(swant)

// 	if sgot != swant {
// 		t.Errorf(`got
// %v, wanted
// %v`, sgot, swant)
// 	}

// }
