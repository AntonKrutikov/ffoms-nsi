package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/antonkrutikov/ffoms-nsi"
)

var (
	dir      string
	listFlag bool
	allFlag  bool
	getFlag  string
	verFlag  bool
)

func init() {
	flag.StringVar(&dir, "dir", "nsi", "Директория, куда будут сохраняться XML файлы справочников")
	flag.BoolVar(&verFlag, "ver", false, "Добавлять номер версии к названию файла")
	flag.BoolVar(&listFlag, "list", false, "Отобразить список справочников доступных на сайте ФФОМС")
	flag.BoolVar(&allFlag, "all", false, "Загружает все последние версии справочников в директорию, указанную в -dir")
	flag.StringVar(&getFlag, "get", "", "Код справочника для загрузки")
}

func main() {
	flag.Parse()

	if listFlag {
		list, err := ffoms.GetDictionaryList()
		if err != nil {
			log.Fatal(err)
		}

		for _, d := range list {
			fmt.Printf("%s\t%s\t%s\t%s\n", d.Info.Code, d.UserVersion, d.LastUpdate, d.ShortName)
		}
		return
	}

	if allFlag || getFlag != "" {
		_ = os.Mkdir(dir, os.ModePerm)
	}

	if getFlag != "" {
		dic, err := ffoms.FindDictionary(getFlag)
		if err != nil {
			log.Fatal(err)
		}
		if err := saveFile(dic); err != nil {
			log.Fatal(err)
		}
	}

	if allFlag {
		list, err := ffoms.GetDictionaryList()
		if err != nil {
			log.Fatal(err)
		}

		for _, d := range list {
			if err := saveFile(&d); err != nil {
				log.Fatal(err)
			}
		}
	}

	flag.PrintDefaults()
}

func saveFile(d *ffoms.Dictionary) error {

	file, err := d.GetFile()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.xml", d.Info.Code)
	if verFlag {
		filename = fmt.Sprintf("%s_%s.xml", d.Info.Code, d.UserVersion)
	}

	return ioutil.WriteFile(dir+"/"+filename, file, 0666)
}
