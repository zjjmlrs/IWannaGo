/*
author: zjjmlrs
date: 2018/3/2
*/

package GoUnsafe

import (
	"fmt"
	"testing"
	"unsafe"


)

func TestUnsafe(t *testing.T) {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	pb := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.c)))
	// 给转换后的指针赋值
	*pb = 10
	// 结构体内容跟着改变
	fmt.Println(s)

}

func TestAlignment(t *testing.T) {
	// 内存对齐
	var x struct {
		a bool
		b int16
		c []int
		d bool // sizeof(x)= 40  整个结构体对齐 Alignment=8
	}

	//通常情况下布尔和数字类型需要对齐到它们本身的大小(最多8个字节),其它的类型对齐到机器字大小.(64位的机器字大小为64位,8字节)

	fmt.Printf("%-30s%-30s%-30s%-50s\n",
		"Row", "Sizeof", "Alignof(对齐倍数)", "Offsetof(偏移量)")

	fmt.Printf("%-30s%-30d%-30d%-50s\n",
		"x", unsafe.Sizeof(x), unsafe.Alignof(x), "")
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.a", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.b", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.c", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.c", unsafe.Sizeof(x.d), unsafe.Alignof(x.d), unsafe.Offsetof(x.d))

}

