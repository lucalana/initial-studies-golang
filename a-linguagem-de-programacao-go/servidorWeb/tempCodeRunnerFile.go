http.HandleFunc("/imagem", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})