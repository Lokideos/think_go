package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Lokideos/think_go/task-3-interfaces-mocks/pkg/spider"
)

// Scanner interface
type Scanner interface {
	Scan(string, int) (map[string]string, error)
}

func main() {
	const url = "http://info.cern.ch/"
	fmt.Println("Данная программа предназначена для поиска информации на заданном сайте.")
	fmt.Printf("На данный момент в качестве сайта выбра %s\n", url)
	fmt.Println("Производится сканирование сайта. Пожалуйста, подождите...")
	crawler := new(spider.Spider)
	si, err := buildSearchIndex(crawler, url, 2)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
		return
	}

	fmt.Println("Можно приступать к поиску.")
	fmt.Println("Для выхода из программы введите 'exit'")
	fmt.Println("---------------------")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите интересующее вас слово: ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("exit", text) == 0 {
			fmt.Println("Завершение работы программы...")
			return
		}

		wc := 0
		keys := []string{}

		for k, v := range si {
			if strings.Contains(v, text) {
				keys = append(keys, k)
				wc++
			}
		}

		if wc == 0 {
			fmt.Println("Введенное слово не найдено на выбранном сайте.")
			continue
		}

		fmt.Printf("Введенное слово %s было найдено на следующих страницах сайта %s:\n", text, url)
		for _, k := range keys {
			fmt.Println(k)
		}

		fmt.Print("\n")
	}
}

func buildSearchIndex(s Scanner, url string, depth int) (map[string]string, error) {
	fmt.Println("Производится сканирование сайта. Пожалуйста, подождите...")
	data, err := s.Scan(url, depth)

	if err != nil {
		return nil, err
	}

	fmt.Println("Сканирование сайта завершено.")

	return data, err
}
