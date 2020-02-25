package main

import (
	//    "time"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAllFile(pathname string, s chan<- string) error {
	rd, err := ioutil.ReadDir(pathname)
	//fmt.Println("正在查找->>", pathname)
	if err != nil {
		return err
	}
	for _, fi := range rd {
		fullPath := pathname + "/" + fi.Name()
		if IsDir(fullPath) {
			err = GetAllFile(fullPath, s)
			if err != nil {
				return err
			}
		}
		if IsFile(fullPath) {
			s <- fullPath
		} else {
			continue
		}

	}
	return nil
}

func main() {
	if len(os.Args) != 0 {
		fmt.Println(os.Args[0])
	}
    if os.Args[1]=="-h"{
        fmt.Println("Usage:")
        fmt.Println("1:      ./getfile  SrcPath ToPath fileID")
        fmt.Println("2:      SrcPath ToPath should be realpath like '/pa/pa' or './p1'")
        fmt.Println("3:      PWD should have dataInjection fold")
        fmt.Println("4:      fileID should be like a20191107 and not pure numeric string")
        fmt.Println("5:      Enjoy!")
        return
    }
	filepath := "/Volumes/SanDisk/数据管理"
	if len(os.Args) > 1 {
		filepath = os.Args[1]
	}
	topath := "/home/ubuntu/zdyf/dataInjection/nos"
	if len(os.Args) > 2 {
		topath = os.Args[2]
	}

	fileID := "a20191025"
	if len(os.Args) > 3 {
		fileID = os.Args[3]
	}

	//fmt.Println("查找文件夹是:", filepath)
	s := make(chan string, 10)

	Len := 3
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ii := 0
		for {
			fi := <-s
			fmt.Println(ii)
			fmt.Println(fi)
			ii += 1
			if fi == "done" {
				return
			}
			info, _ := AnalysePath(fi, topath,fileID, Len)
			MakeDirs(info)
			MakeFile(info)
		}
	}()
	err := GetAllFile(filepath, s)
	if err == nil {
		s <- "done"
	}
	wg.Wait()
}

func AnalysePath(spath, topath,fileID string, pathDepth int) (map[string]interface{}, error) {
	info := make(map[string]interface{})
    fmt.Println(spath)
    fmt.Println(pathDepth)
	pathsli := strings.Split(spath, "/")
    fmt.Println(pathsli)
	Len := len(pathsli)
	u1 := uuid.NewV4()
	exts := strings.Split(spath, ".")
	ext := ""
	if len(exts) > 1 {
		ext = "." + exts[len(exts)-1]
	}
	for i, v := range pathsli[:Len-1] {
		k := fmt.Sprintf("%d级目录", i)
		info[k] = v
	}
	info["fileID"] = fileID
	info["ext"] = ext
	info["filepath"] = spath
	info["filename"] = pathsli[Len-1]
	info["loadname"] = u1.String() + ext
	info["jsonname"] = u1.String() + ".json"
	info["load"] = topath + "/load/"
	info["merge"] = topath + "/merge/"
	info["meta"] = topath + "/meta/"
	return info, nil
}

func MakeDirs(info map[string]interface{}) (string, error) {
	err1 := MakeDir(info["load"].(string))
	err2 := MakeDir(info["merge"].(string))
	err3 := MakeDir(info["meta"].(string))
	return "", nil
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	if err3 != nil {
		return "", err3
	}
	return "ok", nil
}

func MakeFile(info map[string]interface{}) (string, error) {
	//WriteFile(info["load"]+info["loadname"], "123")
	w, err4 := copy(info["filepath"].(string), info["load"].(string)+info["loadname"].(string))
	if err4 != nil {
		return "", err4
	}
	//fmt.Println(w)
	info["size"] = w
	mjson, _ := json.Marshal(info)
	mString := string(mjson)
	err5 := WriteFile(info["meta"].(string)+"/"+info["jsonname"].(string), mString)
	if err5 != nil {
		return "", err5
	}
	return "ok", nil

}
func WriteFile(filename, content string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
		// 创建文件失败处理

	} else {
		_, err = f.Write([]byte(content))
		if err != nil {
			return err
			// 写入失败处理

		}
	}
	return nil
}

func MakeDir(_dir string) (err error) {
	exist, err := PathExists(_dir)
	if err != nil {
		return
	}

	if exist {
		err = errors.New("failed! Existed")
	} else {
		// 创建文件夹
		err := os.MkdirAll(_dir, os.ModePerm)
		if err != nil {
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	return
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
