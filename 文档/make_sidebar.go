package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

func main() {
	// sidebar, err := os.Open("./_sidebar.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sidebar.Close()

	filenames := make([]string, 0, 500)
	filepath.Walk("./文档", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" || filepath.Ext(path) == "" {
			// fmt.Println(path)
			filenames = append(filenames, path)
		}
		return nil
	})

	// for i := range filenames {
	// 	fmt.Println(filenames[i])
	// }

	f1, _ := os.Create("./文档/_sidebar.md")
	w := bufio.NewWriter(f1)
	for _, pathname := range filenames {
		fmt.Println(pathname)
		name := filepath.Base(pathname)
		// fmt.Println(name)
		if filepath.Ext(name) == ".md" {
			w.WriteString(fmt.Sprintf("\t* [%s](%s)\n", name[:len(name)-3], pathname))
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
