package logic

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		SPut(w, r)
		return
	}
	if m == http.MethodGet {
		SGet(w, r)
		return
	}
	if m == http.MethodDelete {
		SDelete(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func SPut(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	contentLen := r.ContentLength

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

	exten := strings.Split(r.URL.EscapedPath(), "/")[5]
	//log.Println(r.URL.EscapedPath())
	//C:\Users\Administrator\go\src\awesomeProject\test_file
	err = os.Mkdir("./file"+"/sample/"+exten, 0777)
	err = os.Mkdir("./file/sample/"+exten+"/sample", 0777)
	if err != nil {
		w.Write([]byte("can not create sample"))
		fmt.Println(err)
		return
	}
	f, _ := os.OpenFile("./file"+"/sample/"+exten+"/"+exten+".zip", os.O_RDWR|os.O_CREATE, 0755)
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

			w.Write([]byte("http://" + viper.GetString("app.server") + "/api/admin/upload/sample/" + exten))

		}

	}
	dst := "./file/sample/" + exten + "/sample"
	dst2 := "./file/sample/" + exten
	archive, err := zip.OpenReader(dst2 + "/" + exten + ".zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	for _, f := range archive.File {
		filePath := dst

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}

}

func SGet(w http.ResponseWriter, r *http.Request) {

	f, e := os.Open("./file" + "/sample/" + strings.Split(r.URL.EscapedPath(), "/")[5] + "/" + strings.Split(r.URL.EscapedPath(), "/")[5] + ".zip")

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func SDelete(w http.ResponseWriter, r *http.Request) {
	e := os.RemoveAll("./file" + "/sample/" + strings.Split(r.URL.EscapedPath(), "/")[5])
	if e != nil {
		log.Println(e)
	}
	w.Write([]byte{'o', 'k'})
}
