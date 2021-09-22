package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "Kiyik"
    var revStr string = ""
	var length = len(str)

	for i := length - 1; i >= 0; i -- {
		revStr = revStr + string(str[i])
	}

	if  strings.ToLower(str) == strings.ToLower((revStr)){
		fmt.Println(str, "is palindrome");
	} else {
		fmt.Println(str, "is NOT a Palindrome");
	}
}