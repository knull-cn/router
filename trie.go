package router

import (
	"fmt"
	"strings"
)

func LogTrace(fmtstr string, args ...interface{}) {
	//fmt.Printf("%s\n", fmt.Sprintf(fmtstr, args...))
}

type trieNode struct {
	Key    string
	Value  interface{}
	parent *trieNode
	Childs map[string]*trieNode
	path   string
}

type TrieTree struct {
	root *trieNode
}

var HasSetValue = fmt.Errorf("Repeat set value")
var LogicError = fmt.Errorf("Logic Error!!!")
var EmptyPath = fmt.Errorf("path is empty!!!")

func (tt *TrieTree) AddPath(path string, v interface{}) error {
	path = strings.TrimSpace(path)
	if len(path) == 0 {
		return EmptyPath
	}
	//LogTrace("addpath(%s)", path)
	arrs := []string{""}
	if path != "/" {
		//如果path=="/"，那么split的到的是两个"",所以先判断下;
		arrs = strings.Split(path, "/")
	}

	if tt.root == nil {
		tt.root = &trieNode{Childs: map[string]*trieNode{}}
	}
	node := addNode(0, arrs, tt.root)
	if node.Value == nil {
		node.Value = v
		//LogTrace("set value(%+v) to(%s)", v, node.path)
	}
	return HasSetValue
}

func (tt *TrieTree) RepleasePath(path string, v interface{}) error {
	path = strings.TrimSpace(path)
	if len(path) == 0 {
		return EmptyPath
	}
	arrs := strings.Split(path, "/")
	if tt.root == nil {
		tt.root = &trieNode{Childs: map[string]*trieNode{}}
	}
	node := addNode(0, arrs, tt.root)
	node.Value = v
	return nil
}

func addNode(pos int, arrs []string, cur *trieNode) *trieNode {
	key := arrs[pos]
	fullpath := strings.Join(arrs[:pos+1], "/")
	//LogTrace("\tkey=%s;fullpath=%s;arrs=%+v;", key, fullpath, arrs[pos:])
	//
	child, ok := cur.Childs[key]
	if !ok {
		//create;
		child = &trieNode{
			Key:    key,
			Childs: map[string]*trieNode{},
			parent: cur,
			path:   fullpath,
		}
		cur.Childs[key] = child
	}
	//不是最后,那么要继续添加trietree node;
	if len(arrs) > pos+1 {
		pos++
		return addNode(pos, arrs, child)
	}
	//是叶子节点了,那么value.
	return child
}

func (tt *TrieTree) GetValue(path string) interface{} {
	path = strings.TrimSpace(path)
	if len(path) == 0 {
		return nil
	}
	arrs := strings.Split(path, "/")
	node := findNode(0, arrs, tt.root)
	if node == nil {
		return nil
	}
	return node.Value
}

func findNode(pos int, arrs []string, cur *trieNode) *trieNode {
	key := arrs[pos]
	child, ok := cur.Childs[key]
	LogTrace("\tkey(%d)=%s;arrs=%+v", pos, key, arrs)
	if !ok {
		LogTrace("\treturn_1")
		return nil
	}
	//is last;
	if len(arrs) <= pos+1 {
		LogTrace("\t\treturn_2=%s;v=%v", child.Key, child.Value)
		return child
	}
	pos++
	node := findNode(pos, arrs, child)
	//如果没找到,那么直接返回child;
	if node == nil || node.Value == nil {
		// if node != nil {
		// 	LogTrace("\t\treturn_3=%s;v=%v", child.Key, child.Value)
		// }
		return child
	}
	//如果以后都没找到,那么返回nil;
	LogTrace("\t\treturn_4")
	return node
}
