/*
author: ZengJJ
date: 2018/3/6  
*/
package GoGoGo

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	//创建一个新的list
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	prtList(l)                   //输出list的值,01234
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出尾部元素的值,4
	fmt.Println()

	// 插入元素
	l.InsertAfter(6, l.Front())               //首部元素之后插入一个值为6的元素
	prtList(l)                                //输出list的值,061234
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	prtList(l)                                //输出list的值,601234
	l.MoveToFront(l.Back())                   //将尾部元素移动到首部
	prtList(l)                                //输出list的值,460123
	fmt.Println()

	// 新建另一个list
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	prtList(l2)        //输出l2的值,460123
	// 删除list中对应的元素
	for e := l2.Front(); e != nil; e = e.Next() {
		if e.Value == 6 {
			l2.Remove(e)
		}
	}
	prtList(l2)
	fmt.Println()

	// 清空list
	l.Init()           //清空l
	fmt.Print(l.Len()) //0
	prtList(l)         //输出list的值,无内容

}

func prtList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func TestDeleteElemByIndex(t *testing.T) {
	index := 2
	test1 := []int{1, 2, 3, 4, 5}
	test1 = append(test1[:index], test1[index+1:]...)
	fmt.Println(test1)
}
