package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, g := range dataset {
		err := dp.Parse(g)
		if err != nil {
			log.Println("ошибка парсинга:", err)
			continue
		}
		g, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка парсинга:", err)
			continue
		}
		fmt.Println(g)
	}
}
