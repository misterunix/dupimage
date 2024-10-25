Sometimes you need a bit of help because of time or your own code is getting messy.  
Here are some of the places I got advice from.

# logging
https://www.jajaldoang.com/post/go-write-log-to-file/

# recursive decent
My code was so ugly. This is AI generated and it works great.  
```
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "." // Replace with the directory you want to search

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
```