package edd_curl
import (
	"net/url"
	"net/http"
	"strings"
	"log"
	"encoding/base64"
	"io/ioutil"
)

type Request struct {
	url string //url地址
	req *http.Request //请求实例
	cli *url.Values
	headers map[string]string//请求头
	postData map[string]string//post提交参数
}

//构造request实例对象
func NewRequst(url string)*Request  {
	if url=="" {
		log.Fatalln("Lack of request url")
	}
	return &Request{
		url:url,
	}
}
//设定post提交参数
func (this *Request)SetPostData(postData map[string]string)*Request{
	this.postData=postData
	return this
}

//将用户post参数处理
func (this *Request)setPostData() *strings.Reader {
	if this.postData == nil {
		log.Fatalln("Lack of request params")
	}

	for k,v:=range this.postData{
		this.cli.Add(k,v)
	}
	return strings.NewReader(this.cli.Encode())

}

//post请求
func (this *Request)Post()string  {
	return this.send(http.MethodPost)
}

//get请求
func (this *Request)Get()string  {
	return this.send(http.MethodGet)
}

//设定headers
func (this *Request)SetHeaders(headers map[string]string) *Request {
	this.headers=headers
	return this
}

//将用户自定义请求头添加到http.Request实例
func (this *Request)setHeaders(){
	if this.headers==nil {
		log.Fatalln("Lack of request headers")
	}
	for k, v := range this.headers {
		this.req.Header.Set(k,v)
	}
}

//发送请求
func (this *Request)send(method string) string{
	this.cli=&url.Values{}
	req,err:=http.NewRequest(method,this.url,this.setPostData())
	if err !=nil{
		log.Fatalln(err)
	}

	this.req=req
	this.setHeaders()

	/**************处理响应数据***************/
	resp,err := http.DefaultClient.Do(req)
	if err!=nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}


func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic "+base64.StdEncoding.EncodeToString([]byte(auth))
}