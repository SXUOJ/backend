package logic

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"web_app/pkg/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		Put(w, r)
		return
	}
	if m == http.MethodGet {
		Get(w, r)
		return
	}
	if m == http.MethodDelete {
		Delete(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func Put(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	contentLen := r.ContentLength

	fmt.Printf("upload content-type:%s,content-length:%d", contentType, contentLen)
	if !strings.Contains(contentType, "multipart/form-data") {
		w.Write([]byte("content-type must be multipart/form-data"))
		return
	}
	if contentLen >= 50*1024*1024 { // 50 MB
		w.Write([]byte("file to large,limit 50MB"))
		return
	}
	err := r.ParseMultipartForm(50 * 1024 * 1024)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("ParseMultipartForm error:" + err.Error()))
		return
	}

	if len(r.MultipartForm.File) == 0 {
		w.Write([]byte("not have any file"))
		return
	}
	id, _ := uuid.Getuuid()
	exten, _ := getContentType(strings.Split(r.URL.EscapedPath(), "/")[5])
	fmt.Println(id, exten)
	//log.Println(r.URL.EscapedPath())
	//C:\Users\Administrator\go\src\awesomeProject\test_file
	f, _ := os.OpenFile("./file"+"/img/"+id+"."+exten, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	for name, files := range r.MultipartForm.File {
		fmt.Printf("req.MultipartForm.File,name=%s", name)

		if len(files) != 1 {
			w.Write([]byte("too many files"))
			return
		}
		if name == "" {
			w.Write([]byte("is not FileData"))
			return
		}
		for _, ff := range files {
			handle, err := ff.Open()
			if err != nil {
				w.Write([]byte(fmt.Sprintf("unknown error,fileName=%s,fileSize=%d,err:%s", ff.Filename, ff.Size, err.Error())))
				return
			}

			io.Copy(f, handle)
			fmt.Printf("successful uploaded,fileName=%s,fileSize=%.2f MB,savePath=%s \n", id+"."+exten, float64(contentLen)/1024/1024, f.Name())

			w.Write([]byte("http://" + viper.GetString("app.server") + ":" + viper.GetString("app.port") + "/api/admin/upload/image/" + id + "." + exten))
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

	f, e := os.Open("./file" + "/img/" + strings.Split(r.URL.EscapedPath(), "/")[5])

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func getContentType(fileName string) (extension, contentType string) {
	arr := strings.Split(fileName, ".")

	// see: https://tool.oschina.net/commons/
	if len(arr) >= 2 {
		extension = arr[len(arr)-1]
		switch extension {
		case "jpeg", "jpe", "jpg":
			contentType = "image/jpeg"
		case "png":
			contentType = "image/png"
		case "gif":
			contentType = "image/gif"
		case "mp4":
			contentType = "video/mpeg4"
		case "mp3":
			contentType = "audio/mp3"
		case "wav":
			contentType = "audio/wav"
		case "pdf":
			contentType = "application/pdf"
		case "doc", "":
			contentType = "application/msword"
		}
	}
	// .*（ 二进制流，不知道下载文件类型）
	contentType = "application/octet-stream"
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	e := os.Remove("./file" + "/img/" + strings.Split(r.URL.EscapedPath(), "/")[5])
	if e != nil {
		log.Println(e)
	}
	w.Write([]byte{'o', 'k'})
}
