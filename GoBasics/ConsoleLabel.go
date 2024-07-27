package GoBasics

import "fmt"

func PrintSection(label string) {
	fmt.Printf("\n\n=======================================================================================\n%s\n=======================================================================================\n", label)
}

func Line() {
	fmt.Println("----------------------------------------------------------")
}

func Label(label string) {
	fmt.Printf("\n--[%s]--\n", label)
}
