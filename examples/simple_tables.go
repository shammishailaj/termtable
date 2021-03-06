package main

import (
	"fmt"
	"github.com/stevedomin/termtable"
)

func main() {
	fmt.Println("Simple table:")

	t := termtable.NewTable(nil, nil)
	t.SetHeader([]string{"LOWERCASE", "UPPERCASE", "NUMBERS"})
	t.AddRow([]string{"abc", "ABCD", "12345"})
	t.AddRow([]string{"defg", "EFGHI", "678"})
	t.AddRow([]string{"hijkl", "JKL", "9000"})
	fmt.Println(t.Render())

	fmt.Println("\nSimple table, alternative syntax:")

	rows := [][]string{
		[]string{"abc", "ABCD", "12345"},
		[]string{"defg", "EFGHI", "678"},
		[]string{"hijkl", "JKL", "9000"},
	}
	t = termtable.NewTable(rows, nil)
	t.SetHeader([]string{"LOWERCASE", "UPPERCASE", "NUMBERS"})
	fmt.Println(t.Render())

	fmt.Println("\nSimple table w/ separators and custom padding:")

	t = termtable.NewTable(nil, &termtable.TableOptions{
		Padding:      3,
		UseSeparator: true,
	})
	t.SetHeader([]string{"LOWERCASE", "UPPERCASE", "NUMBERS"})
	t.AddRow([]string{"abc", "ABCD", "12345"})
	t.AddRow([]string{"defg", "EFGHI", "678"})
	t.AddRow([]string{"hijkl", "JKL", "9000"})
	fmt.Println(t.Render())
}
