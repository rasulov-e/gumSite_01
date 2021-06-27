package server

import (
	"database/sql"
	"gumsite_01/storage"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	headers handlers.CORSOption
	methods handlers.CORSOption
	origins handlers.CORSOption
	store   storage.Storage
}

func NewServer() *Server {
	s := &Server{
		router:  mux.NewRouter(),
		headers: handlers.AllowedHeaders([]string{"X-Requsted-With", "Content-Type", "Authorization"}),
		methods: handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"}),
		origins: handlers.AllowedOrigins([]string{"*"}),
	}
	return s
}

func (s *Server) Start(port string, dbConfig string) error {
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.store = storage.New(db)
	defer db.Close()

	s.handleRequests()
	return http.ListenAndServe(port, handlers.CORS(s.headers, s.methods, s.origins)(s.router))
}

func (s *Server) handleRequests() {
	s.router.HandleFunc("/selectAllCategories", s.SelectAllCategories()).Methods("GET")
	s.router.HandleFunc("/selectCategoryByName", s.SelectCategoryByName()).Methods("GET")
	s.router.HandleFunc("/insertCategory", s.InsertCategory()).Methods("POST")
	s.router.HandleFunc("/deleteCategory", s.DeleteCategory()).Methods("DELETE")

	// handles all functions for product assortiment management
	s.router.HandleFunc("/selectAllProducts", s.SelectProducts()).Methods("GET")
	s.router.HandleFunc("/selectProductByName", s.SelectProductByName()).Methods("GET")
	s.router.HandleFunc("/insertProduct", s.AddProduct()).Methods("POST")
	s.router.HandleFunc("/deleteProduct", s.DeleteProduct()).Methods("DELETE")

	// handles all functions for order management
	s.router.HandleFunc("/selectAllOrders", s.SelectAllOrders()).Methods("GET")
	s.router.HandleFunc("/selectOrderByName", s.SelectOrderByName()).Methods("GET")
	s.router.HandleFunc("/insertOrder", s.InsertOrder()).Methods("POST")
	s.router.HandleFunc("/deleteOrder", s.DeleteOrder()).Methods("DELETE")

}
