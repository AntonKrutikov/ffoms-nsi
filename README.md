Пакет позволяет загружать справочники ФФОМС с сайта [http://nsi.ffoms.ru/](http://nsi.ffoms.ru/) в формате XML

    go get github.com/antonkrutikov/ffoms-nsi

В комплекте есть простая консольная утилита `cmd\ffoms-nsi`,  позволяющая загрузить все справочники в виде файлов.

    # go run cmd/ffoms-nsi/main.go

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