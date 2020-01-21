package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/axgle/mahonia"
)

var AGENTS map[string]string

//var COMMAND
//TIME := list.New()

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := mahonia.NewDecoder("gbk").ConvertString(string(r.Form.Get("data")))
	//路由正则
	url_info, _ := regexp.Compile("/info/*")
	//url_md,_:=regexp.Compile("/md/*")
	url_cm, _ := regexp.Compile("/cm/*")

	//url path
	fmt.Println("URL_path", r.URL.Path)

	//路由/info/*
	if url_info.MatchString(r.URL.Path) {
		//匹配

		//输出data

		fmt.Println("Form", data)
		AGENTS = make(map[string]string)
		url_path, _ := regexp.Compile(`[A-Z]+`)
		id := url_path.FindString(r.URL.Path)
		fmt.Println(id)
		AGENTS[id] = "id"
		fmt.Println(AGENTS[id])

	} else if url_cm.MatchString(r.URL.Path) {
		//r.ParseForm()

		//命令实现需要配合输入
		url_path, _ := regexp.Compile(`[A-Z]+`)
		var id = url_path.FindString(r.URL.Path)
		value, ok := AGENTS[id]
		if ok {
			fmt.Println(value)
		} else {
			fmt.Fprintf(w, "REGISTER")
		}

	} else {
		//全都不匹配输出请求详细
		//先强制断开连接
		fmt.Println(r.Close)
		//自动关闭服务器

		fmt.Println("Request解析")
		//HTTP方法
		fmt.Println("method", r.Method)
		// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
		fmt.Println("RequestURI", r.RequestURI)
		//URL类型,下方分别列出URL的各成员
		fmt.Println("URL_scheme", r.URL.Scheme)
		fmt.Println("URL_opaque", r.URL.Opaque)
		fmt.Println("URL_user", r.URL.User.String())
		fmt.Println("URL_host", r.URL.Host)
		fmt.Println("URL_path", r.URL.Path)
		fmt.Println("URL_RawQuery", r.URL.RawQuery)
		fmt.Println("URL_Fragment", r.URL.Fragment)
		//协议版本
		fmt.Println("proto", r.Proto)
		fmt.Println("protomajor", r.ProtoMajor)
		fmt.Println("protominor", r.ProtoMinor)

		//打印全部头信息
		for k, v := range r.Header {
			// fmt.Println("Header key:" + k)
			for _, vv := range v {
				fmt.Println("header key:" + k + "  value:" + vv)
			}
		}

		//解析body
		//r.ParseMultipartForm(128)
		//fmt.Println("解析方式:ParseMultipartForm")
		r.ParseForm()
		fmt.Println("解析方式:ParseForm")

		//body内容长度
		fmt.Println("ContentLength", r.ContentLength)

		//打印全部内容
		fmt.Println("Form", r.Form)

		//该请求的来源地址
		fmt.Println("RemoteAddr", r.RemoteAddr)

		///data:=r.RemoteAddr
		//发送邮件通知
		//SendMail("Danger notice ！！！！",data)
		os.Exit(0)
	}

	//fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

//func SendMail(subject string, body string ) error {
//    //定义邮箱服务器连接信息
//    mailConn := map[string]string {
//        "user": "",
//        "pass": "",
//        "host": "",
//        "port": "",
//    }
//
//    port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
//
//    m := gomail.NewMessage()
//    m.SetHeader("Subject", subject)  //设置邮件主题
//    m.SetBody("text/html", body)     //设置邮件正文
//
//    d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
//
//    err := d.DialAndSend(m)
//    return err
//
//}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
