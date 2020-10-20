package main

import (
	"fmt"
	"log"

	"github.com/Lokideos/think_go/task-2-search-engine/pkg/spider"
)

func main() {
	const url = "https://www.nix.ru"
	data, err := spider.Scan(url, 2)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	for k, v := range data {
		fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
	}
}
