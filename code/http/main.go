package main

/* Do an HTTP request to google.com and print the body of the response to terminal
 */
import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/text/encoding/charmap"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.es")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//(A)
	//fmt.Println(resp)

	//bs :=[]byte{}

	// (B)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// (C)
	//io.Copy(os.Stdout, resp.Body)

	// (D)
	lw := logWriter{}
	//io.Copy(lw, resp.Body)

	// (E) arregla acentos enyes es y tal que salian mal por pantalla (y al redigir a fichero) pese a tener mintty que usa la consola de cygwin y todo en UTF-8
	// El problema es que por defecto html debe viajar en iso-8859-1
	defer resp.Body.Close()

	rdrBody := io.Reader(resp.Body)
	//contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	//if strings.Contains(contentType, "charset=iso-8859-1") {
	rdrBody = charmap.ISO8859_1.NewDecoder().Reader(rdrBody)
	io.Copy(lw, rdrBody)
	//}
	// if strings.Contains(contentType, "charset=utf-8") {
	// rdrBody = charmap.
	//NewDecoder().Reader(rdrBody)
	// }

	// El codigo hacia abajo mostraba solos estos primeros 254 caracteres
	// <!doctype html><html itemscope=noodp" name="r
	// body, err := ioutil.ReadAll(rdrBody)
	// if err != nil {
		// fmt.Println(err)
		// return
	// }

	// n := 256
	// if n > len(body) {
		// n = len(body)
	// }
	// fmt.Println(string(body[:n]))

}

// (D)
// Creating this function i am achieving that type logWriter implement the Writer interface.
func (logWriter) Write(bs []byte) (int, error) {
	// Crap implementation that allow us compile but does not do the right thing of "writing"
	// return 1, nil

	// A good implementation that writes  to something (p.e. to console) external to our program.
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
