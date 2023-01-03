package logic

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strings"
	"web_app/pkg/uuid"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		SPut(w, r)
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

func SPut(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	contentLen := r.ContentLength

	fmt.Printf("upload content-type:%s,content-length:%d", contentType, contentLen)
	if !strings.Contains(contentType, "multipart/form-data") {
		w.Write([]byte("content-type must be multipart/form-data"))
		return
	}
	if contentLen >= 300*1024*1024 { // 300 MB
		w.Write([]byte("file to large,limit 300MB"))
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
