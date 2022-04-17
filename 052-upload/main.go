package main
import "fmt"
import "log"
import "html/template"
import "net/http"
import "io/ioutil"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, nil)
	} else {
		// 1. parse input
		r.ParseMultipartForm(10 << 20)
		// 2. retrieve file
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		// 3. write temporary file on our server
		tempFile, err := ioutil.TempFile("images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		// 4. return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}

/*func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Uploading File")
}
*/

func setupRoutes() {
    http.HandleFunc("/uploads", uploadFile)
	port := ":5000"
	fmt.Println("Server is running on port" + port)
    log.Fatal(http.ListenAndServe(port, nil))
}
func main() {
    setupRoutes()
}