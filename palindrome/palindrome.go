package palindrome

import "fmt"

func isPalindrome(s string) bool {
	strLenth := len(s)
	for i := 0; i < strLenth/2; i++ {
		if s[i] != s[strLenth-1-i] {
			return false
		}
	}
	return true
}

func main() {

	var input string

	fmt.Println("Enter a string:")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome")
	}

}
