package ffoms

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//DictionaryGroup - принадлежность справочника
type DictionaryGroup struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
}

//DictionaryInfo ...
type DictionaryInfo struct {
	Code     string `json:"code"`
	Mnemonic string `json:"mnemonic"`
}

//DictionaryOrganization - организация, ответсвенная за ведение справочника
type DictionaryOrganization struct {
	ID   int    `json:"int"`
	Name string `json:"name"`
}

type dictionaryCollection struct {
	List []Dictionary `json:"list"`
}

//Dictionary - информация о справочнике ФФОМС
type Dictionary struct {
	UserVersion   string                 `json:"user_version"`
	Group         DictionaryGroup        `json:"d_group"`
	FullName      string                 `json:"fullName"`
	Info          DictionaryInfo         `json:"d"`
	ProviderParam string                 `json:"providerParam"`
	Organization  DictionaryOrganization `json:"respOrganization"`
	LastUpdate    string                 `json:"last_update"`
	ID            int                    `json:"id"`
	ShortName     string                 `json:"shortName"`
}

//GetFile загружает указанную версию справочника в формате XML
func (d *Dictionary) GetFile() ([]byte, error) {
	if d.ProviderParam == "" {
		return nil, fmt.Errorf("ошибка запроса файла: пустое поле ProviderParam")
	}

	s := strings.Split(d.ProviderParam, "v")
	if len(s) != 2 {
		return nil, fmt.Errorf("ошибка запроса файла: неправильное значение поля ProviderParam (%s)", d.ProviderParam)
	}

	id := s[0]
	version := s[1]

	url := fmt.Sprintf("http://nsi.ffoms.ru/refbook?type=XML&id=%s&version=%s", id, version)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка формирования GET запроса: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать ответ: %v", err)
	}

	zip, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть zip архив: %v", err)
	}

	file, err := zip.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить файл из zip архива: %v", err)
	}

	result, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать полученный файл справочника: %v", err)
	}

	return result, nil
}

//GetDictionaryList получает список справочников с сайта ФФОМС
func GetDictionaryList() ([]Dictionary, error) {
	url := "http://nsi.ffoms.ru/data?pageId=refbookList&containerId=refbookList&size=110"

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка формирования GET запроса: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать ответ: %v", err)
	}

	var collection dictionaryCollection

	err = json.Unmarshal(body, &collection)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить структуру из JSON ответа: %v", err)
	}

	return collection.List, nil
}

//FindDictionary получает последнюю версию справочника по его коду
func FindDictionary(code string) (*Dictionary, error) {
	collection, err := GetDictionaryList()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить списрок справочников")
	}

	for _, v := range collection {
		if v.Info.Code == code {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("справочник с указанным кодом не найден в списке справочников")
}
