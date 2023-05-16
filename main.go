package main

import (
	"bufio"
	"fmt"
	"os"

	json "github.com/neilotoole/jsoncolor"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	enc.SetColors(json.DefaultColors())

	for scanner.Scan() {
		var m map[string]any
		err := json.Unmarshal(scanner.Bytes(), &m)
		if err != nil {
			fmt.Println(scanner.Text())
			continue
		} 

		err = enc.Encode(m)
		if err != nil {
			fmt.Println(scanner.Text())
		}
	}
}
