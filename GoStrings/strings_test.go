/*
author: zjjmlrs
date: 2018/2/4  
*/
package GoStrings

import (
	"testing"
	"strings"
	"fmt"
)

func TestReaderBytes(t *testing.T) {

	r := strings.NewReader("test中文")
	r.Seek(1, 0)
	fmt.Println(r.Len(), r.Size())

	b1 := make([]byte, 12)
	r.Read(b1)
	fmt.Println(string(b1))
	fmt.Println(r.Read(b1)) // 0 EOF

	fmt.Println(r.ReadAt(b1, 1)) // 9 EOF
	fmt.Println(string(b1))

}

func TestReaderRune(t *testing.T) {
	r := strings.NewReader("test中文")
	ch, size, _ := r.ReadRune()
	fmt.Println(string(ch), size)

	// 读取·中·
	r.Seek(4, 0)
	ch, size, _ = r.ReadRune()
	fmt.Println(string(ch), size)

	// 回退
	r.UnreadRune()
	ch, _, _ = r.ReadRune()
	fmt.Println(string(ch))

	// 继续读
	ch, _, _ = r.ReadRune()
	fmt.Println(string(ch))

	// 定位错误
	r.Seek(5, 0)
	ch, size, _ = r.ReadRune()
	fmt.Println(string(ch), size)

}

func TestReplacer(t *testing.T) {
	r := strings.NewReplacer("ab", "ac", "cd", "11")
	fmt.Println(r.Replace("abcdeabcf"))

	r = strings.NewReplacer("a", "c", "c", "11")
	fmt.Println(r.Replace("abcdeabcf"))
}
