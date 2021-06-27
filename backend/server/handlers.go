package server

import (
	"encoding/json"
	"fmt"
	"gumsite_01/storage"
	"net/http"
	"strconv"
)

// SelctProducts return a variable of type http.HandleFunct that
// Selects all products from database by calling storage.selectAllProducts
// it returns all products in json format
func (s *Server) SelectProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := storage.SelectAllProducts(s.store)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, p)
	}
}

// SelctProductsByName return a variable of type http.HandleFunct that
// takes a (json string: name) and searches for a product with the same name in db
// by calling a storage.FindProductByName
// it returns that product in json format
func (s *Server) SelectProductByName() http.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		p, err := storage.FindProductByName(s.store, req.Name)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *Server) DeleteProduct() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		storage.DeleteProduct(s.store, req.Name)
		s.respond(w, r, http.StatusOK, "done")
	}
}

func (s *Server) AddProduct() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Cost        string `json:"cost"`
		Image       string `json:"image"`
		CategoryID  string `json:"category_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cost, err := strconv.Atoi(req.Cost)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		categoryId, err := strconv.Atoi(req.CategoryID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		p := storage.Product{
			Name:        req.Name,
			Description: req.Description,
			Cost:        cost,
			Image:       req.Image,
			CategoryID:  categoryId,
		}
		if err := storage.InserProduct(s.store, p); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusCreated, p)
	}
}

func (s *Server) SelectAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := storage.SelectAllOrders(s.store)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		fmt.Fprint(w, o)
		s.respond(w, r, http.StatusOK, o)
	}
}

func (s *Server) SelectOrderByName() http.HandlerFunc {
	type Request struct {
		Customer string `json:"customer_name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		c, err := storage.SelectOrderByName(s.store, req.Customer)
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
			return
		}
		s.respond(w, r, http.StatusCreated, c)
	}
}

// InsertOrder returns a valiable of type http.HandleFunc
// that inserts a new order into the db table gumsite_orders
// by calling storage.InsertOrder
func (s *Server) InsertOrder() http.HandlerFunc {
	type Request struct {
		Customer          string `json:"customer"`
		Address           string `json:"address"`
		Orders            []int  `json:"orders"`
		Cost              int    `json:"cost"`
		Payment           int    `json:"payment"`
		OrderDate         string `json:"order_date"`
		OrderShippingDate string `json:"order_shipping_date"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, &json.UnmarshalTypeError{})
			return
		}
		o := storage.Order{
			Customer:          req.Customer,
			Address:           req.Address,
			Orders:            req.Orders,
			Cost:              req.Cost,
			Payment:           req.Payment,
			OrderDate:         req.OrderDate,
			OrderShippingDate: req.OrderShippingDate,
		}
		if err := storage.InsertOrder(s.store, o); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, "done")
	}
}

func (s *Server) DeleteOrder() http.HandlerFunc {
	type Request struct {
		Customer string `json:"customer_name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := storage.DeleteOrderByName(s.store, req.Customer); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
	}
}

func (s *Server) InsertCategory() http.HandlerFunc {
	type Request struct {
		Name     string `json:"name"`
		Qunataty int    `json:"quantity"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		c := storage.Category{
			CategoryName:     req.Name,
			CategoryQuantity: req.Qunataty,
		}
		if err := storage.InsertCategory(s.store, c); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, "done")
	}
}

func (s *Server) SelectAllCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := storage.SelectAllCategories(s.store)
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
			return
		}
		s.respond(w, r, http.StatusCreated, c)
	}
}

func (s *Server) SelectCategoryByName() http.HandlerFunc {
	type Request struct {
		CategoryName string `json:"category_name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		c, err := storage.SelectCategoryByName(s.store, req.CategoryName)
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
			return
		}
		s.respond(w, r, http.StatusCreated, c)
	}
}

func (s *Server) DeleteCategory() http.HandlerFunc {
	type Request struct {
		CategoryName string `json:"category_name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := storage.DeleteCategory(s.store, req.CategoryName); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
	}
}

// error is an analog of respond
// but used exucivly for sending errors back to the client
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{
		"error": err.Error(),
	})
}

// respond sends a json data
// usualy used to send json back to the client
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
