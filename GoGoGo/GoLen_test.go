/*
author: ZengJJ
date: 2018/2/4  
*/
package core

import (
	"testing"
	"fmt"
)

func TestLen(t *testing.T) {
	fmt.Println(len("gfargr中文grf"))         // length of bytes
	fmt.Println(len([]rune("gfargr中文grf"))) // length of unicode char

	fmt.Println(len(string(rune('编'))))
}

func TestRune(t *testing.T) {
	s := "中文grf"
	fmt.Println(string(s[0]))         // get one byte
	fmt.Println(string([]rune(s)[0])) // get one unicode char

	// string 类型的 for range，是按 rune 来遍历的
	for _, c := range s {
		fmt.Println(string(c)) // get one unicode char
	}
}
