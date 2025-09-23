package main

import (
	"log"
	"net/http"
	"wishlistApi/internal/database"
	"wishlistApi/internal/wish"
	"wishlistApi/internal/wishList"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contentType", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "test root call"}`))
}

func main() {
	var m = http.NewServeMux()
	var s = &server{}

	database.Migrate()

	m.HandleFunc("/", s.ServeHTTP)

	m.HandleFunc("GET /api/v1/wish/{wishId}", wish.GetWishById)
	m.HandleFunc("DELETE /api/v1/wish/{wishId}", wish.DeleteWish)
	m.HandleFunc("PUT /api/v1/wish/{wishId}", wish.UpdateWish)
	m.HandleFunc("POST /api/v1/wish/", wish.CreateWish)

	m.HandleFunc("GET /api/v1/wishList/{wishListId}", wishList.GetWishListById)
	m.HandleFunc("DELETE /api/v1/wishList/{wishListId}", wishList.DeleteWishList)
	m.HandleFunc("PUT /api/v1/wishList/{wishListId}", wishList.UpdateWishList)
	m.HandleFunc("POST /api/v1/wishList/", wishList.CreateWishList)

	m.HandleFunc("GET /api/v1/user/{userId}/wishList/{wishListId}", wishList.GetWishListById)

	log.Println("Run serve")
	log.Fatal(http.ListenAndServe(":9000", m))
}
