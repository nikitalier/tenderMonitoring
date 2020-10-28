package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

func (app *Application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !app.svc.Auth(user.Login, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user = app.svc.FindUserByLogin(user.Login)
	tokenStruct, err := app.svc.GenerateJWT(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenJSON, _ := json.Marshal(tokenStruct)
	w.Write(tokenJSON)
}

func (app *Application) findUserHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	user := app.svc.FindUserByLogin(login)
	j, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := app.svc.GetAllUsers()

	j, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) getAllKeyWords(w http.ResponseWriter, r *http.Request) {
	keywords := app.svc.GetAllKeyWords()

	j, err := json.Marshal(keywords)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) deleteKeyword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := app.svc.DeleteKeywords(string(body))

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func (app *Application) addKeywordHandler(w http.ResponseWriter, r *http.Request) {
// 	var keyword models.Keyword

// 	err := json.NewDecoder(r.Body).Decode(&keyword)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// log.Println(keyword)

// 	result := app.svc.AddKeyword(keyword)

// 	if !result {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

func (app *Application) addKeywordHandler(w http.ResponseWriter, r *http.Request) {
	var keyword models.Keyword

	err := json.NewDecoder(r.Body).Decode(&keyword)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// log.Println(keyword)

	id, result := app.svc.AddKeyword(keyword)

	// log.Println(id)
	// log.Println(result)

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// log.Println(id)

	j, _ := json.Marshal(id)
	// log.Println(j)
	w.Write(j)
	// w.Write([]byte(fmt.Sprint(id)))
}

func (app *Application) getAllTenders(w http.ResponseWriter, r *http.Request) {
	tenders := app.svc.GetAllTenders()

	// log.Println(tenders)

	j, err := json.Marshal(tenders)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) getTenderHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tender := app.svc.GetTender(idInt)
	j, err := json.Marshal(tender)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) updateFavorite(w http.ResponseWriter, r *http.Request) {
	var f models.Favorite
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Println(f)

	app.svc.UpdateFavorite(f)
}

func (app *Application) getFavorite(w http.ResponseWriter, r *http.Request) {
	var f models.Favorite

	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newF := app.svc.GetFavoriteStatus(f)

	j, err := json.Marshal(newF)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) getAllComments(w http.ResponseWriter, r *http.Request) {
	tenderID := r.URL.Query().Get("tenderid")
	tenderIDInt, err := strconv.Atoi(tenderID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comments := app.svc.GetAllComments(tenderIDInt)

	j, err := json.Marshal(comments)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) addNewComment(w http.ResponseWriter, r *http.Request) {
	var comments models.Comment
	err := json.NewDecoder(r.Body).Decode(&comments)
	// log.Println(comments)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.svc.AddNewComment(comments)
	// w.WriteHeader(http.StatusBadRequest)
}

func (app *Application) createTenderStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.svc.CreateTenderStatus(idInt)
}

func (app *Application) updateTenderStatus(w http.ResponseWriter, r *http.Request) {
	var status models.TenderStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Println(f)

	app.svc.UpdateTenderStatus(status)
}

func (app *Application) getTenderStatus(w http.ResponseWriter, r *http.Request) {

	tenderID := r.URL.Query().Get("id")
	tenderIDInt, err := strconv.Atoi(tenderID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status := app.svc.GetTenderStatus(tenderIDInt)

	j, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) getSummary(w http.ResponseWriter, r *http.Request) {
	summary := app.svc.GetSummary()

	j, err := json.Marshal(summary)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(j)
}

func (app *Application) testHandler(w http.ResponseWriter, r *http.Request) {
	app.svc.Test()
}
