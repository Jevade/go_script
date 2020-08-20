package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/qiniu/api.v7/v7/auth/qbox"
    "github.com/qiniu/api.v7/v7/storage"
    "golang.org/x/net/context"
    "github.com/atotto/clipboard"
    "github.com/Jevade/go_script/gofile/getfile"
)
//MyPutRet  定制返回数据结构体
type MyPutRet struct {
    Key    string
    Hash   string
    Fsize  int
    Bucket string
    Name   string
}


func GetAllFile(filePath string) (error,string) {
    if getfile.IsFile(filePath) {
        return nil,filePath
    }
    return nil,"error"
}


func main(){
	if len(os.Args) != 0 {
		fmt.Println(os.Args[0])
    }else{
        fmt.Println("请输入参数")
        return
    }

    if os.Args[1]=="-h"{
        fmt.Println("Usage:")
        fmt.Println("1: ./qiniu  SrcPath")
        return
    }

    filepath := os.Args[1]
    fmt.Println("上传文件:"+filepath)
    _,file := GetAllFile(filepath)
    err,result := putFile(file)
    if err!=nil{
        fmt.Println("上传失败，退出")
        return
    }
    clipboard.WriteAll(result)
    fmt.Println("文件名义复制到剪切板")
    return
}


func putFile(localFile string)(error,string){
// 自定义返回值结构体
    bucket := "vtol"
    accessKey := "vA5IXFjRpguFwQO4_eEiunvONORSek1hN4iN7NZS"
    secretKey := "qlE6jfP3dmrlOUw-fBYmF0jBceUe-HIfevqqw3Ba"
    pathToWatch := "/Users/liu/VTOL/pic"
    pathsli := strings.Split(localFile, "/")
    key := pathsli[len(pathsli)-1]
    fmt.Println(pathToWatch)
// 使用 returnBody 自定义回复格式
    putPolicy := storage.PutPolicy{
        Scope:      bucket,
        ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
    }
    mac := qbox.NewMac(accessKey, secretKey)
    fmt.Println("初始化链接")
    upToken := putPolicy.UploadToken(mac)
    cfg := storage.Config{}
    formUploader := storage.NewFormUploader(&cfg)
    ret := MyPutRet{}
    putExtra := storage.PutExtra{
        Params: map[string]string{
            "x:name": "github logo",
        },
    }
    fmt.Println("开始上传")
    err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
    if err != nil {
        fmt.Println(err)
        return err,""
    }
    fmt.Println("结束上传")
    fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
    return nil,"http://q675hxulc.bkt.clouddn.com/"+ret.Key
}
