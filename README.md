# 1.Python版

Python  版 开源项目 https://github.com/ahmedkhlief/muddyc3-Revived







# 2.golang版

golang 版 开源项目 https://github.com/demonsec666/muddyc3_golang

这是根据泄漏的ahmedkhlief  MuddyC3 C2的工作POC。现在包括以下功能：

作者: WBGIII 

   

蹭饭系列作者: Demon666

# 

## 1.代理重新连接

![enter description here][1]



## 2. 加载模块

![enter description here][2]

## 3. 发送命令和接收结果



## 4. 创建Powershell负载  ----（已实现一半，还未增加到服务端）目前只有get.ps1  代码复制到客户端，进行测试



## 5.已实现Download 功能，但待修复upload功能（有部分bug）

![enter description here][3]

## 已更新增加
 ```
   1. help
   2. info
   3. exit

   
   
 ```



[Demo](https://youtu.be/gD93kX_Eq_Y)





## 待增加功能版：

 ```
   1. 需修复upload bug
   2. 需增加help功能菜单
   3. 需增加 list-info、show command、use ID等相关功能上菜单
   4. 需增加 Tab键，展示命令
   5. 需增加back 功能
 
   
   
 ```

   

## 用法：

 ``` 
   1. go get github.com/axgle/mahonia  go get github.com/olekukonko/tablewriter
   2. go run main   （PORT:9090）
   3. 复制get.ps1 代码到客户端 （并修改ip和端口）
   4. 在控制可直接输入命令
   5. 或者在控制输入  load xxxx.ps1  or Download  serverfile  clientfile  
   6. load  即加载 moudle文件中的文件
   7. Donwload 需创建file文件，再使用命令Download  serverfile(指的是file下的文件 相对路径)  clientfile  （客户端的绝对路径）
 ```

   

# 敬告(Notice):
+  仅供学习参考，做测试

+  不合理使用此项目所提供的功能而造成的任何直接或者间接的后果及损失， 均由使用者本人负责，即刻安全以及创作者不为此承担任何责任。


+   For reference only, for testing 

+   Any direct or indirect consequences and losses arising from the abuse of the featuress provided by those  project  are due to the user himself, secist and the author does not accept any responsibility.


[1]: https://demonsec666.oss-cn-qingdao.aliyuncs.com/CA7F0BB98761EF4426EB1D7FA7E223CD.jpg
[2]: https://demonsec666.oss-cn-qingdao.aliyuncs.com/9311DF125870D1C86BF186D5AA8C532C.jpg
[3]: https://demonsec666.oss-cn-qingdao.aliyuncs.com/2CA777D7D57FFD6C177C9261523B601E.jpg



