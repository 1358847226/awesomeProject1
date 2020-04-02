package api

import (
	"awesomeProject1/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	_ "io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func GetAccessToken() string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp2,_ := client.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxf835047e3970b804&secret=cc3b3e36d7475099bc452be494208095",)
	defer resp2.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp2.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)
	comma := strings.Index(resultData2, ":")
	d := strings.Index(resultData2,",")
	res2 := resultData2[comma+2 :d-1]
	return res2
}

func SubscribeMessage(res2 string, data models.SubscribeMessage) string{
	client := &http.Client{Timeout: 5 * time.Second}
	b ,_ := json.Marshal(data)

	log.Println("b",data)
	body := bytes.NewBuffer(b)
	log.Println("body",body)
	resp, _ := client.Post("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="+res2,"content-type:application/json;charset=UTF-8",
		body)
	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()


	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)

	return resultData2
}

func MediaCheckAsync(token string, url string) string{
	client := &http.Client{Timeout: 5 * time.Second}
	type Data struct {
		//Access_token string `json:"access_token"`
		Media_url string `json:"media_url"`
		Media_type int `json:"media_type"`
	}
	var data Data
	data = Data{
		//Access_token: token,
		Media_url:    url,
		Media_type:   1,
	}
	b ,_ := json.Marshal(data)
	log.Println("b",data)
	body := bytes.NewBuffer(b)
	log.Println("body",body)
	resp, _ := client.Post("https://api.weixin.qq.com/wxa/media_check_async?access_token="+token,"content-type:application/json;charset=UTF-8",
		body)
	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()

	log.Println("res213",resultData2)

	return resultData2
}

func ImgSecCheck(token string, file *multipart.FileHeader) string{
	client := &http.Client{Timeout: 5 * time.Second}
	type Media struct {
		//Access_token string `json:"access_token"`
		ContentType string `json:"content_type"`
		Value *multipart.FileHeader `json:"value"`
	}
	var media Media
	media = Media{
		ContentType: "image/jpeg",
		Value:       file,
	}
	log.Println(media)
	b ,_ := json.Marshal(media)
	log.Println("b",media)
	body := bytes.NewBuffer(b)

	log.Println("body",body)

	resp, _ := client.Post("https://api.weixin.qq.com/wxa/img_sec_check?access_token="+token,"application/octet-stream",
		body)

	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	type MeadiaResult struct {
		Errcode int `json:"errcode"`
		Errmsg string `json:"errmsg"`
	}
	var meadiaResult MeadiaResult
	if err := json.Unmarshal([]byte(resultData2), &meadiaResult); err == nil {
		if meadiaResult.Errcode == 0{

		}
	} else {
		fmt.Println(err)
	}
	log.Println("res213",resultData2)

	return resultData2
}

func MsgSecCheck(token string, content string) string{
	client := &http.Client{Timeout: 5 * time.Second}
	type Data struct {
		//Access_token string `json:"access_token"`
		Content string `json:"content"`
	}
	var data Data
	data = Data{Content:content}
	b ,_ := json.Marshal(data)
	log.Println("b",data)
	body := bytes.NewBuffer(b)
	log.Println("body",body)
	resp, _ := client.Post("https://api.weixin.qq.com/wxa/msg_sec_check?access_token="+token,"content-type:application/json;charset=UTF-8",
		body)
	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()

	log.Println("res213",resultData2)

	return resultData2
}