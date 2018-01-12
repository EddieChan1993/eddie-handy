package edd_curl
import (
	"net/url"
	"net/http"
	"strings"
	"log"
	"encoding/base64"
	"io/ioutil"
	"sync"
)

type Request struct {
	url    string //url地址
	req    *http.Request //请求实例
	cli    *url.Values
	header map[string]string //请求头
	param  map[string]string //post提交参数
	sync.RWMutex
}

//构造request实例对象
func NewRequst(url string,headers map[string]string)*Request  {
	if url=="" {
		log.Fatalln("Lack of request url")
	}
	return &Request{
		url:url,
		header:headers,
	}
}
//传入请求参数，POST/GET
func (this *Request) SetParms(postData map[string]string)*Request{
	this.param =postData
	return this
}

//将参数加入请求中
func (this *Request) initParams() *strings.Reader {
	if this.param == nil {
		log.Fatalln("Lack of request params")
	}

	for k,v:=range this.param {
		this.cli.Add(k,v)
	}
	return strings.NewReader(this.cli.Encode())

}

//post请求
func (this *Request)Post()([]byte,error)  {
	return this.send(http.MethodPost)
}

//get请求
func (this *Request)Get()([]byte,error)  {
	return this.send(http.MethodGet)
}

//将用户自定义请求头添加到http.Request实例
func (this *Request) initHeaders(){
	if this.header ==nil {
		log.Fatalln("Lack of request header")
	}
	for k, v := range this.header {
		this.req.Header.Set(k,v)
	}
}

//发送请求
func (this *Request)send(method string) ([]byte,error){
	this.Lock()
	defer this.Unlock()

	this.cli=&url.Values{}
	req,err:=http.NewRequest(method,this.url,this.initParams())
	if err !=nil{
		return nil,err
	}

	this.req=req
	this.initHeaders()

	/**************处理响应数据***************/
	resp,err := http.DefaultClient.Do(req)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}


func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic "+base64.StdEncoding.EncodeToString([]byte(auth))
}