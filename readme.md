## ireact-命令行生成react组件
## 用法
- 将ireact.exe和templates文件夹复制到项目根目录
- 编辑package.json
```json
 "scripts": {
    "ir":"./ireact.exe"
  },
```
## 请手动在src下创建components或者views文件夹，否则会报错
## 生成组件-执行
```
yarn ir -cn Hello
```
运行结果

>在src/components/中生成Hello.js文件
## 生成页面-执行
```
yarn ir -vn Hello
```
运行结果

>在src/views/中生成Hello.js文件

## 生成多个
```shell
yarn ir -vn Hello Helloa Hellob -cn Foo
```
运行结果
```shell
$ go run main.go -vn Hello Helloa Hellob -cn Foo
生成Hello页面成功
生成Helloa页面成功
生成Hellob页面成功
生成Foo组件成功
```