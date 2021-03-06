package main

import "fmt"
import "net/http" // package untuk web server
import "html/template" // package untuk template engine

// contoh penggunaan sistem template/layout 


func contoh(res http.ResponseWriter, req *http.Request) {
  
  // tentukan terlebih dahulu html mana saja yang akan dipakai
  var layout = template.Must(template.ParseFiles(
       "layout_html/bagian_head.html",
       "layout_html/bagian_header.html",
       "layout_html/halaman_contoh_kondisi.html",
       "layout_html/bagian_footer.html",
   ))
   
  // cek file layout_html/halaman_contoh.html untuk mengecek nama templatenya dibagian DEFINE
  var err = layout.ExecuteTemplate(res, "halaman_contoh_kondisi", nil)
  
  // cek error
  if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func main() {

  // contoh route
  http.HandleFunc("/", contoh) 

  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
}
