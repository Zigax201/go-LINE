package api

import (
	"fmt"
	"go-line/model"
	"go-line/util"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rs/xid"
)

func FileUpload(r *http.Request) (string, error) {
	file, _, err := r.FormFile("imageUpload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("upload", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}

func (a *App) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	r.ParseMultipartForm(32 << 20)

	imgPath, err := FileUpload(r)
	if err != nil {
		util.RespondError(w, 400, "Error when uploading file")
	}

	guid := xid.New()
	p.ID = guid.String()
	p.Caption = r.FormValue("caption")
	author, _ := strconv.Atoi(r.FormValue("author"))
	p.Author = author //User ID from middleware
	p.ImagePath = imgPath

	a.db.Create(&p)

	util.RespondJSON(w, 200, p)
}

func (a *App) GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	posts := []model.Post{}
	a.db.Find(&posts)

	util.RespondJSON(w, 200, posts)
}

func (a *App) GetBasedOnIDHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	vars := mux.Vars(r)

	id := vars["id"]
	result := a.db.Where("id = ?", id).First(&p)
	if result.RowsAffected == 0 {
		util.RespondError(w, 400, "Post Not Found")
	}

	util.RespondJSON(w, 200, p)
}

func (a *App) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	vars := mux.Vars(r)

	id := vars["id"]
	result := a.db.Where("id = ?", id).First(&p)
	if result.RowsAffected == 0 {
		util.RespondError(w, 400, "Post Not Found")
	}
	if err := a.db.Delete(&p).Error; err != nil {
		util.RespondError(w, 500, "Error when deleting")
	}

	util.RespondJSON(w, 200, map[string]string{"message": "Post Deleted"})
}

func (a *App) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	vars := mux.Vars(r)

	id := vars["id"]
	result := a.db.Where("id = ?", id).First(&p)
	if result.RowsAffected == 0 {
		util.RespondError(w, 400, "Post Not Found")
	}
	r.ParseMultipartForm(32 << 20)

	if r.FormValue("caption") != "" {
		p.Caption = r.FormValue("caption")
	}
	if r.FormValue("author") != "" {
		author, _ := strconv.Atoi(r.FormValue("author"))
		p.Author = author //User ID from middleware
	}
	imgPath, err := FileUpload(r)
	if err != nil {
		p.ImagePath = imgPath
	}

	util.RespondJSON(w, 200, p)
}
