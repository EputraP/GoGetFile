package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	fs := http.FileServer(http.Dir("./image"))
	http.Handle("/", cors(fs))
	http.HandleFunc("/image", GetImageHandler)
	// construct JSON data
	data := make(map[string]interface{})
	data["image"] = "base64data"
	data["barcodeFormat"] = 234882047
	data["maxNumPerPage"] = 1
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("/image2", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server Running")
	http.ListenAndServe(":4352", nil)
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Printf("Image Handler Running")
	content, err := ioutil.ReadFile("image/wqwqwq.png")
	if err != nil {
		log.Fatal(err)
	}
	var base64Encoding string

	base64Encoding += "data:image/png;base64,"
	// // Determine the content type of the image file
	// mimeType := http.DetectContentType(content)

	// // Prepend the appropriate URI scheme header depending
	// // on the MIME type
	// switch mimeType {
	// case "image/jpeg":
	// 	base64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	base64Encoding += "data:image/png;base64,"
	// }

	// Append the base64 encoded output
	base64Encoding += toBase64(content)
	fmt.Println(base64Encoding)
	// fmt.Printf("file content: ", content)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// jsonData := []byte(`{"status":` + base64Encoding + `}`)
	jsonData := []byte(`{"status":` + base64Encoding + `}`)
	w.Write([jsonData,jsonData])

}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do your cors stuff
		// return if you do not want the FileServer handle a specific request
		enableCors(&w)
		fs.ServeHTTP(w, r)
	}
}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
