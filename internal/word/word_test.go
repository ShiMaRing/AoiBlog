package word

import (
	"fmt"
	"testing"
)

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	var s = "hello_code_kkk"
	camelCase := UnderscoreToUpperCamelCase(s)
	fmt.Println(camelCase)
}

func TestCamelCaseUnderscore(t *testing.T) {
	var s = "helloWorldHaa&&&"
	underscore := CamelCaseUnderscore(s)
	fmt.Println(underscore)
}
