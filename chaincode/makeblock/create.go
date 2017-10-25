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
func (b *Block) Init(args []string) bool {
	if len(args) != 3 {
		fmt.Println("no key and value")
		return true
	}
	n, _ := strconv.Atoi(args[2])
	b.create(args[1], n)
	return true
}

func (b *Block) create(key string, value int) error {
	fmt.Println("Begin creating a new block")
	*b = Block{key, value}
	fmt.Println("Finish creating a new block")
	return nil
}

func main() {
	if block.Init(os.Args) == false {
		fmt.Println("error")
		return
	}
	fmt.Println("main func finishes")
	return
}
