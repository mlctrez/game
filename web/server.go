package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/mlctrez/wasmexec"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	// TODO: remove and use embedded FS
	files := http.FileServer(http.Dir("web"))
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/favicon.ico":
			w.WriteHeader(404)
		case "/app.js":
			writeAppJs(w)
		case "/app.wasm":
			err := buildWasm()
			if err != nil {
				w.WriteHeader(500)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
			files.ServeHTTP(w, r)
		default:
			files.ServeHTTP(w, r)
		}
	}
	mux.HandleFunc("/", handler)

	_ = http.ListenAndServe(":8080", mux)
}

//go:embed *.js *.html *.css
var fs embed.FS

func writeAppJs(w http.ResponseWriter) {

	var err error
	var content []byte

	content, err = wasmexec.Current()
	if err != nil {
		w.WriteHeader(500)
		log.Printf("error getting current wasmexec: %v", err)
		return
	}

	var launch []byte
	launch, err = fs.ReadFile("app_launch.js")
	if err != nil {
		w.WriteHeader(500)
		log.Printf("error reading app_launch.js: %v", err)
		return
	}

	_, err = w.Write(append(append(content, []byte("\n\n")...), launch...))
	if err != nil {
		log.Printf("write error: %v", err)
		return
	}

}

func buildWasm() (err error) {
	command := exec.Command("go", "build", "-o", "web/app.wasm", ".")
	command.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	var output []byte
	if output, err = command.CombinedOutput(); err != nil {
		return fmt.Errorf("%s \n\n %s", string(output), err.Error())
	}
	var stat os.FileInfo
	if stat, err = os.Stat("web/app.wasm"); err != nil {
		return err
	}
	fmt.Println(stat.Size())
	return nil
}
