package main

import (
	"bufio"
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"strconv"
	"strings"
)

type KeywordPair struct {
	keyword     string
	replacement string
}

func main() {
	fmt.Println("Hello World.")

	fmt.Println("How many documents you want to generate today?")

	var documentCount int
	_, err := fmt.Scanln(&documentCount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(documentCount)

	fmt.Println("Now, select your document")

	filepath, _ := dialog.File().Title("Select your document").Load()

	in := bufio.NewReader(os.Stdin)

	file, _ := docx.ReadDocxFile(filepath)

	for i := 1; i <= documentCount; i++ {
		var kps []KeywordPair
		fmt.Printf("Insert the keywords and replacements for file number %s (type /q to quit) \n", strconv.Itoa(i))
		newFile := file.Editable()
		for b := false; !b; {
			fmt.Println("Type your keyword")
			kw, _ := in.ReadString('\n')
			kw = strings.TrimSpace(strings.TrimSuffix(kw, "\n"))
			if kw == "/q" {
				b = true
				break
			}
			fmt.Println("Type your replacement")
			rp, _ := in.ReadString('\n')
			rp = strings.TrimSpace(strings.TrimSuffix(rp, "\n"))
			if rp == "/q" {
				b = true
				break
			}
			kps = append(kps, KeywordPair{keyword: kw, replacement: rp})
		}

		for _, pair := range kps {
			_ = newFile.Replace(pair.keyword, pair.replacement, -1)
			fmt.Printf("Chave: %s, Valor: %s\n", pair.keyword, pair.replacement)
		}
		_ = newFile.WriteToFile("./doc" + strconv.Itoa(i) + ".docx")
	}

	fmt.Println(filepath)

}
