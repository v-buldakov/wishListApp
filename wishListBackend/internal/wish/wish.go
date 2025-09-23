package wish

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Wish struct {
	Id         int64  `json:"wishId"`
	Name       string `json:"name"`
	Url        string `json:"wishUrl"`
	PicUrl     string `json:"wishPicUrl"`
	WishListId int    `json:"wishListId"`
	Comleted   bool   `json:"comleted"`
}

func GetWishById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id, _ = strconv.ParseInt(r.PathValue("wishId"), 10, 32)
	var newTestWish = Wish{id, "testName", "https://github.com/moficodes/bookdata-api/blob/master/routes.go", "", 0, false}
	w.WriteHeader(http.StatusOK)
	var j, _ = json.Marshal(newTestWish)
	log.Println(newTestWish)
	w.Write(j)
}

func CreateWish(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var newWish Wish
	var errMarsh = json.Unmarshal(body, &newWish)
	if errMarsh != nil {
		http.Error(w, errMarsh.Error(), http.StatusInternalServerError)
		return
	}

	newWish.Comleted = true
	newWish.Id = 42

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var j, _ = json.Marshal(newWish)
	log.Println(newWish)
	w.Write(j)
}

func UpdateWish(w http.ResponseWriter, r *http.Request) {

}

func DeleteWish(w http.ResponseWriter, r *http.Request) {

}
