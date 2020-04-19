package main

import (
	"fmt"
	"golang-practice/iterator"
	"golang-practice/makemap"
	"golang-practice/slice"
	"golang-practice/structure"
)

func main() {
	iterator.Iterator()
	iterator.IteratorSecond()
	fmt.Println(iterator.IteratorThird(3))
	structure.Structure()
	structure.StructureSecond()
	structure.StructureThird()
	slice.Slice()
	slice.SliceSecond()
	slice.SliceThird()
	slice.SliceFourth()
	slice.SliceFifth()
	makemap.MakeMap()
	makemap.MakeMapSecond()
	makemap.MakeMapThird()
}
