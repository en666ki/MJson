package server

import (
    "log"
    "fmt"
    "net/http"
    "path/filepath"
    "os"
    "time"
    "io"

    "github.com/en666ki/MJson/pkg/json_reader"

    "github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./web/spa/index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("upload")
	r.Body = http.MaxBytesReader(w, r.Body, 1024 * 1024)
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}
    file, fileHeader, err := r.FormFile("file")
	if err != nil {
        fmt.Println("no file")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

    fmt.Println("File passed checks")

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

    fmt.Println("dir checked")

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    var result bool
    result, err = json_reader.Parse(dst.Name())
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result {
        fmt.Fprintf(w, "Json is correct")
    } else {
        fmt.Fprintf(w, "Json is incorrect")
    }

}

func Main() {
    fmt.Println("Init server")
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeHandler)
    router.HandleFunc("/upload", uploadHandler).Methods("POST")
    srv := &http.Server{
        Handler:    router,
        Addr:    ":3000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Start server")
    log.Fatal(srv.ListenAndServe())
}
