package main

import (
	// "fmt"
	"log"
	// "net/url"
	// "sync"

	"github.com/leenzstra/teacher_parser/pkg/config"
	"github.com/leenzstra/teacher_parser/pkg/db"
	"github.com/leenzstra/teacher_parser/pkg/models"
	"github.com/leenzstra/teacher_parser/pkg/parser"
)

// var alph = []string{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ы", "Э", "Ю", "Я"}

func main() {
	c, err := config.LoadConfig()

	if c.StartParse == 0 {
		log.Fatalln("Parsing = 0, exit...")
	}

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := db.Init(c.PostgresURL)

	// teachersChan := make(chan []models.Teacher, len(alph))
	// teachersAll := []models.Teacher{}
	// wg := sync.WaitGroup{}

	// for _, letter := range alph {
	// 	fmt.Println(letter)
	// 	wg.Add(1)

	// 	go func(char string) {
	// 		defer wg.Done()
	// 		parser := parser.New()
	// 		u := "https://pstu.ru/basic/glossary/staff/?p=1&sign="+url.QueryEscape(char)
			
	// 		teachers := parser.Start(u)

	// 		teachersChan <- teachers
	// 		fmt.Println("Ended", char,  len(teachers))
	// 	}(letter)
	// }

	// wg.Wait()

	// for v := range teachersChan {
	// 	fmt.Println("Count", len(v))
	// 	teachersAll = append(teachersAll, v...)
	// }

	parser := parser.New()
	teachersAll := parser.Start("https://pstu.ru/basic/glossary/staff/?sign=none")

	db.Unscoped().Where("1 = 1").Delete(&models.Teacher{})
	err = db.Save(teachersAll).Error
	if err != nil {
		log.Fatalln("Failed to save", err)
	}

}
