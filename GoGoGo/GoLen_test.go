/*
author: zjjmlrs
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

	/*
	在Go语言中支持两个字符类型，
	一个是 byte (uint8)，代表UTF-8字符串的单个字节的值；
	另一个是 rune (int32)，代表单个Unicode字符。出于简化语言的考虑，Go语言的多数API都假设字符串为UTF-8编码。
	*/
}

func TestAppend(t *testing.T) {
	data := make([]int, 5, 10)
	data[0] = 1
	data[1] = 2

	app1 := []int{1, 2, 3, 4, 5} //len <=5 则 	result[0] = 99 会 影响源Slice
	result1 := append(data, app1...)
	result1[0] = 99 // 修改到了data[0]
	result1[7] = 98

	fmt.Println("length:", len(app1), ":", app1) // 无论如何app1 不会变
	fmt.Println("length:", len(data), ":", data)
	fmt.Println("length:", len(result1), ":", result1)

	app2 := []int{1, 2, 3, 4, 5, 6} //len > 5 则 	result[0] = 991 不会 影响源Slice
	result2 := append(data, app2...)
	result2[0] = 991 // 修改到了data[0]
	result2[7] = 981

	fmt.Println("length:", len(app2), ":", app2) // 无论如何app1 不会变
	fmt.Println("length:", len(data), ":", data)
	fmt.Println("length:", len(result2), ":", result2)
	fmt.Println("length:", len(result1), ":", result1)

	// 如果此时app2的长度还是小于5，比如 app2 := []int{1, 2, 3, 4, 5}
	// 则修改app2的元素，会同时修改data和app1的元素，包括app1[7]也会被修改

	/*
	内置函数append可以向一个切片后追加一个或多个同类型的其他值。
	如果追加的元素数量超过了原切片容量，那么最后返回的是一个全新数组中的全新切片。
	如果没有超过，那么最后返回的是原数组中的全新切片。
	*/
}

func TestSlice(t *testing.T) {
	// 切片是对同一个底层数组的引用, 所以其中一个改变了，其他的都会变
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:]
	s[4] = 99
	fmt.Println(s, s1)
}

func TestCopy(t *testing.T) {
	/*
	将内容从一个数组切片复制到另一个数组切片。
	如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
	*/
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)

	slice2 = []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)

	// 互不影响, 显然，slice1和slice2都是已经初始化了的，所以copy的时候只会copy值
	// 而且也不会改变容量
	slice1[0] = 99
	fmt.Println(slice1, slice2)
}
