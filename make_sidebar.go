package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

func main() {
	sidebar, err := os.Open("./_sidebar.md")
	if err != nil {
		log.Fatal(err)
	}
	defer sidebar.Close()

	filenames := make([]string, 0, 500)
	filepath.Walk("./docs", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" || filepath.Ext(path) == "" {
			filenames = append(filenames, path)
		}
		return nil
	})

	// for i := range filenames {
	// 	fmt.Println(filenames[i])
	// }

	f1, _ := os.Create("_sidebar.md")
	w := bufio.NewWriter(f1)
	for _, pathname := range filenames[1:] {
		// fmt.Println(pathname)
		name := filepath.Base(pathname)
		// fmt.Println(name)
		if filepath.Ext(name) == ".md" {
			w.WriteString(fmt.Sprintf("\t* [%s](%s)\n", filepath.Base(name), pathname))
		} else {
			w.WriteString(fmt.Sprintf("* [%s](%s)\n", name, pathname))
		}

		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
	w.Flush()
	fmt.Println("write success!")
	// r := bufio.NewReader(sidebar)
	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 		break
	// 	}
	// 	fmt.Printf(line)
	// }
}
