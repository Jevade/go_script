package ctrl

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"../util"
)

var FILE_PATH = "./mnt/"
var URL_PATH = "/mnt/"

//attach/upload
func init() {
	os.MkdirAll(FILE_PATH, os.ModePerm)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocalFile(w, r)

}

func UploadLocalFile(w http.ResponseWriter, r *http.Request) {
	srcfile, head, err := r.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		util.RespFail(w, err.Error())
		return
	}
	suffix := ".png"
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	firetype := r.FormValue("filetype")
	if len(firetype) > 0 {
		suffix = firetype
	}

	filename := fmt.Sprintf("%d%d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstfile, err := os.Create(FILE_PATH + filename)
	if err != nil {
		log.Println(err.Error())
		util.RespFail(w, err.Error())
		return
	}

	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		log.Println(err.Error())
		util.RespFail(w, err.Error())
		return
	}
	url := URL_PATH + filename
	data := make(map[string]interface{})
	data["url"] = url
	util.RespOK(w, data)

}
