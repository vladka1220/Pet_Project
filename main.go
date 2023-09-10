package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Обработка входа пользователя
	})

	http.HandleFunc("product.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "product.html")
	})

	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			referencia := r.FormValue("referencia")
			productDescription := r.FormValue("product-description")
			productQuantity := r.FormValue("product-quantity")
			loteProduto := r.FormValue("lote-produto")
			maquina := r.FormValue("maquina")

			// действия при добавлении товара
			fmt.Println("Название товара:", referencia)
			fmt.Println("Описание товара:", productDescription)
			fmt.Println("Количество:", productQuantity)
			fmt.Println("Lote do produto:", loteProduto)
			fmt.Println("Máquina:", maquina)

			// Здесь сохранить данные в базе данных или выполнять другие операции
		}
		// Перенаправление после обработки
		http.Redirect(w, r, "product.html", http.StatusFound)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}

