package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch exibe o conteúdo encontrado em cada URL especificada
func main() {
	url := "go.dev/"
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	dstFile, err := os.Create("saida.html")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	b, err := io.Copy(dstFile, resp.Body)
	fmt.Printf("%s", resp.Status)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("Bytes copied: %d", b)
}
