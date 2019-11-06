Пакет позволяет загружать справочники ФФОМС с сайта [http://nsi.ffoms.ru/](http://nsi.ffoms.ru/) в формате XML

    go get github.com/antonkrutikov/ffoms-nsi

## Получить список справочников:

    list, _ := ffoms.GetDictionaryList()

    for _, d := range list {
        fmt.Printf("%s\t%s\t%s\t%s\n", d.Info.Code, d.UserVersion, d.LastUpdate, d.ShortName)
    }

## Получить файл справлоника по его коду:

    dic, _ := ffoms.FindDictionary("F001")
    file, _ := dic.GetFile()
	

## Консольная утилита

В комплекте есть пример [консольной утилиты](https://github.com/AntonKrutikov/ffoms-nsi/releases/download/v1.0.0/ffoms-nsi-win64.exe), позволяющей загрузить все справочники в виде файлов.

    go run github.com/antonkrutikov/ffoms-nsi/cmd/ffoms-nsi

Использование:

    -all
        Загружает все последние версии справочников в директорию, указанную в -dir
    -dir string
        Директория, куда будут сохраняться XML файлы справочников (default "nsi")
    -get string
        Код справочника для загрузки
    -list
        Отобразить список справочников доступных на сайте ФФОМС
    -ver
        Добавлять номер версии к названию файла