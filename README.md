# digdns

## 使用dig命令循环查询递归服务器对domain.txt列表域名的QUERY TIME，并记录返回单次查询时间，平均响应时间及失败率
## 测试环境：Ubuntu


1.使用前清除get10和get20两个文件的内容  
2.在项目根目录执行’go run main.go‘命令执行项目(也可以在根目录执行’go build‘命令后手动执行生成的文件，文件名为go.mod文件中的module名)  
3.等待程序执行完成后查询结果  
