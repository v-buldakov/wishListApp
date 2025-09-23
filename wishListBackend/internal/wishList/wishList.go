package wishList

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wishlistApi/internal/wish"
)

type WishList struct {
	Id     int64
	Name   string
	Wishes []wish.Wish
	UserId int64
}

func GetWishListById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var userId, _ = strconv.Atoi(r.PathValue("userId"))
	var id, _ = strconv.ParseInt(r.PathValue("wishListId"), 10, 32)
	var newTestWishList = WishList{id, "testName", nil, 0}
	w.WriteHeader(http.StatusOK)
	var j, _ = json.Marshal(newTestWishList)
	log.Println(newTestWishList)
	w.Write(j)
}

func CreateWishList(w http.ResponseWriter, r *http.Request) {

}

func UpdateWishList(w http.ResponseWriter, r *http.Request) {

}

func DeleteWishList(w http.ResponseWriter, r *http.Request) {

}
