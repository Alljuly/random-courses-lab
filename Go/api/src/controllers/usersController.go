package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

//Cria usuario
func CreateUser(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}

	var user models.UserModel
	if err = json.Unmarshal(body, &user); err != nil{
		log.Fatal(err)
	}


	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.CreateUserRepository(db)
	createdID, err := repository.Create(user)
	if err != nil {
			log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Usuario com ID %d Criado", createdID)))
}
//Busca usuario por id
func GetUserByID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("get user"))

}
//Busca usuarios
func GetUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("get all users"))

}
//Altera os dados de um usuario
func UpdateUserByID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("update user"))
}
//Remove usuario do banco
func DeleteUserByID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("delete user"))

}