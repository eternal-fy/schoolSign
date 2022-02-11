package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Student struct {
	Sco, Class, Name string
}

func FetchString(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("获取失败")
	}

	byte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("获取错误")
	}
	str := string(byte)
	return str

}
func GetData(classNames []string, fetch string) (string, error) {
	split := strings.Split(fetch, "<br>")
	students := []Student{}

	for _, context := range split {
		arr := strings.Fields(context)
		if !Contains(classNames, arr[1]) {
			continue
		}
		student := Student{arr[0], arr[1], arr[2]}
		students = append(students, student)
	}
	if len(students) == 0 {
		log.Println("学生数为0")
		return "", errors.New("学生数为0")
	}

	data := ""
	for _, context := range students {
		line := context.Sco + " " + context.Class + " " + context.Name + "\n"
		data += line
	}

	return data, nil
}
func Contains(classNames []string, className string) bool {
	for _, item := range classNames {
		if item == className {
			return true
		}
	}
	return false
}
