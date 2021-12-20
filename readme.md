## ireact-命令行生成react组件
## 用法
- 将ireact.exe和templates文件夹复制到项目根目录
- 编辑package.json
```json
 "scripts": {
  "ir":"./ireact.exe"
  },
```
执行
```
yarn ir -n Hello
```
运行结果

>在src/components/中生成Hello.js文件