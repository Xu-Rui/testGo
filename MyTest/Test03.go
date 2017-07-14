package main

import (
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func sort(head *node){
	if head!=nil {
		sort(head.left)
		fmt.Printf("%d", head.data)
		sort(head.right)
	}
}

func find(head *node,data int){

	var temp *node

	if head.data < data{
		if head.left!=nil{
			temp = head.left
		}else{
			head.left = &node{data:data,left:nil,right:nil}
		}
	}else{
		if(head.right!=nil) {
			temp = head.right
		}else{
			head.right = &node{data:data,left:nil,right:nil}
		}
	}

	if temp!=nil{
		find(temp,data)
	}

}
func main(){
	var count int
	fmt.Scanf("%d",&count)
	data:=make([]int,count)

	for i:=0;i<count;i++{
		fmt.Scanf("%d",&data[i])
	}

	head := &node{data:data[0],left:nil,right:nil}

	for i:=1;i<count;i++{
		find(head,data[i])
	}

	sort(head)
}
