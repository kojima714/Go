package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// とりあえずブロック作成の関数を想定してのコーディング
// このファイルは基本機能の内ブロック作成に特価して作成する
type cmd interface {
	Init([]string) error
}
type Block struct {
	key string
	v   int
}
type Asset struct {
	block map[string]int
}

// Init は当面最初のクエリを作るために仕様
func (t *Asset) Init(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("no key and value")
	}
	t.block = map[string]int{}
	n, _ := strconv.Atoi(args[2])
	t.create(args[1], n)
	return nil
}

func (t *Asset) Invoke(args []string) error {
	l := len(args)
	if l <= 1 {
		return fmt.Errorf("nokey and value 'invoke'")
	}
	t.block = map[string]int{}
	n, _ := strconv.Atoi(args[2])

	switch l {
	case 3:
		if args[1] == "set" {
			t.set(args[1], n)
		}
	default:
		return fmt.Errorf("nokey and value 'invoke'")
	}
	return nil
}

func (t *Asset) set(key string, value int) {
	t.create(key, value)
}
func (t *Asset) create(key string, value int) {
	t.block[key] = value
	// TODO
	// ブロックチェーンに書き込む処理を書く
}

func main() {
	var block Asset
	if err := block.Init(os.Args); err != nil {
		log.Fatal(err)
	} else {
		for i := 1; i < 10; i++ {
			block.create("test"+strconv.Itoa(i), i)
		}
		log.Println(block.block)
		log.Println("main func finishes")
	}
}
