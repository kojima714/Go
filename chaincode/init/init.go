package point

import (
	"fmt"
)

// Init は当面最初のクエリを作るために仕様
func Init(args []string) string {
	if len(args) != 3 {
		fmt.Println("no key and value")
		return "error"
	}
	return "succss"
}
