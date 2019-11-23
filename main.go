package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func getAllFiles() []string {
	dirname := "vocab"

	f, err := os.Open(dirname)

	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var filelist []string
	i := 0
	for _, file := range files {
		filelist = append(filelist, string(file.Name()))
		i++
	}
	return filelist
}

func readFile(name string) string {
	var str string
	filename := "vocab/" + name
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		str += s.Text()
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	str = strings.Replace(str, "<font size=\"3\">", "", -1)
	str = strings.Replace(str, "<font size=\"8\">", "", -1)
	str = strings.Replace(str, "<font size=\"12\">", "", -1)
	str = strings.Replace(str, "</font>", "", -1)

	str = strings.Replace(str, "<b>", "", -1)
	str = strings.Replace(str, "</b>", "", -1)

	str = strings.Replace(str, "&nbsp;", "", -1)

	str = strings.Replace(str, "<content>", "", -1)
	str = strings.Replace(str, "</content>", "", -1)

	str = strings.Replace(str, "<p>", "", -1)
	str = strings.Replace(str, "</p>", "", -1)

	str = strings.Replace(str, "<", "\n<", -1)
	str = strings.Replace(str, ">", ">\n", -1)

	return str
}

func getWord(content string) string {
	var str string
	var res1 []string
	res1 = strings.SplitAfter(content, "</title>")
	str = res1[0]
	res1 = strings.Split(str, "<title>\n")
	str = res1[1]
	res1 = strings.Split(str, "\n</title>")
	str = res1[0]
	return str
}
func getMeaning(content string) string {
	var str string
	var res1 []string
	var res2 []string
	_ = res2

	res1 = strings.SplitAfter(content, "</title>")
	content = res1[1]

	res1 = strings.SplitAfter(content, "</i>")
	for _, i := range res1 {
		res2 = strings.Split(i, "<i>")
		str += res2[0]
	}

	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n\n", "\n", -1)
	str = strings.Replace(str, "\n", "\t", -1)

	return str

}

func csvExport(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	filelist := getAllFiles()

	var filecontent []string
	for _, i := range filelist {
		filecontent = append(filecontent, readFile(i))
	}

	data := [][]string{}
	for _, i := range filecontent {
		data = append(data, []string{getWord(i), getMeaning(i)})
	}

	csvExport(data)

}
