package main

import (
	"fmt"
	"errors"
)


// json 的 7 种数据类型
const (
	NULL = iota
	FALSE
	TRUE
	NUMBER
	STRING
	ARRAY
	OBJECT

)

const (
	PARSE_OK = iota

	// 空白 json str
	PARSE_EXPECT_VALUE
	// 不能被识别的值
	PARSE_INVALID_VALUE
	// ???
	PARSE_ROOT_NOT_SINGULAR
)

// 返回类型的枚举 这里应该可以有优化成 具体的类型
func parse() int {
	return 0
}

// 识别 json 字符串中的空白字符
func parseWhiteSpace(s string) (string, error) {
	for index, c := range s {
		if c != " " && c != "\t" && c != "\n" && c != "\r" {
			return s[:index], nil
		}
	}
	return s, errors.New("json str is empty")
}

func parseNull(s string) int {
	if s[:4] != "null" {
		return s, errors.New("null error")
	}
	return 0
}

func parseValue(s string) int {
	return 0
}

func main() {
	nullJSON := "   null"
	fmt.Println(nullJSON)
}
