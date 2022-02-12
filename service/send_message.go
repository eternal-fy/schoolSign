package service

import (
	"github.com/robfig/cron"
	"log"
	"net/http"
	"net/url"
)

func SignService() {
	c := cron.New()
	defer c.Stop()
	spec := "0 0 10/1 * * ? " //crom  s m h d M date
	c.AddFunc(spec, func() {
		fetch := FetchString("http://180.76.172.177/18nodone")
		classNames := []string{"软件工程1804"}
		data, err := GetData(classNames, fetch)
		if err != nil || data == "" {
			return
		}
		sendMessage(data)
	})
	c.Start()

	select {}
}

func sendMessage(message string) {
	urlValues := url.Values{}
	//群号1055664821
	urlValues.Add("group_id", "1055664821")
	urlValues.Add("message", message)
	_, err := http.PostForm("http://106.12.174.72:8081/send_group_msg", urlValues)
	if err != nil {
		log.Println(err)
	}
}
