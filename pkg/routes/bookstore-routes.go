package routes

import (
	// "./controllers/"
	"github.com/gorilla/mux"
	"github.com/firmananshariadjie/book-store-golang/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.GetBook).Methodes("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methodes("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methodes("PUT")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methodes("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methodes("DELETE")
	
	// fmt.Printf("Starter server at port 8080\n")
	// log.Fatal(http.ListenAndServe(":8080",r))
}
