package main

import (
	"fmt"
	"os"
	"strconv"
)

type Block struct {
	key string
	v   int
}

var block Block

// Init は当面最初のクエリを作るために仕様
func (b *Block) Init(args []string) string {
	if len(args) != 3 {
		fmt.Println("no key and value")
		return ""
	}
	b.create(args)
	return "succss"
}

func (b *Block) create(args []string) int {
	fmt.Println("Begin creating a new block")
	fmt.Println(args[1], args[2])
	n, _ := strconv.Atoi(args[2])
	*b = Block{args[1], n}
	fmt.Println(b)
	fmt.Println("Finish creating a new block")
	return -1
}

func main() {
	if block.Init(os.Args) == "" {
		fmt.Println("error")
		return
	}
	fmt.Println("main func finishes")
	return
}
