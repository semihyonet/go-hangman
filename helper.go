package main

func stringToArr(s string) []string {
	newString := make([]string, len(s))

	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		newString[i] = string(chars[i])
	}
	return newString
}
