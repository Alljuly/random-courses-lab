package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Cria usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Não foi possível ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	var usuario usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		http.Error(w, "Erro ao converter usuario em struct", http.StatusBadRequest)
		return
	}

	db, err := banco.Connection()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Erro ao criar statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	inserir, err := statement.Exec(usuario.Name, usuario.Email)
	if err != nil {
		http.Error(w, "Erro ao executar statement", http.StatusInternalServerError)
		return
	}

	idInserido, err := inserir.LastInsertId()
	if err != nil {
		http.Error(w, "Erro ao recuperar ID do usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

//Busca usuario por id
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, "Erro ao converter id para uint.", http.StatusExpectationFailed) 
		return
	}

	db, err := banco.Connection()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco.", http.StatusBadRequest)
		return
	}
	defer db.Close()

	row, err := db.Query("Select * from users WHERE id = ?", ID)
	if err != nil{
		http.Error(w, "Erro ao conectar ao banco.", http.StatusBadRequest)
		return
	}

	var usuario usuario
	for row.Next(){
		if err = row.Scan(&usuario.ID, &usuario.Name, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}

	}
	
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuario); err != nil{
		w.Write([]byte("Erro ao converter usuarios em json"))
		return
	}
	
}

//Busca usuario por id
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	db, err := banco.Connection()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco.", http.StatusBadRequest)
	}

	defer db.Close()

	sql := "Select * from users"

	rows, err := db.Query(sql)
	if err != nil {
		http.Error(w, "Erro ao buscar usuarios", http.StatusNotFound)
		return
	}

	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario
		if err := rows.Scan(&usuario.ID, &usuario.Name, &usuario.Email); err != nil{
			http.Error(w, "Erro ao escanear usuario", http.StatusBadRequest)
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuarios em json"))
		return
	}

}

//Remove usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil{
		w.Write([]byte("Erro ao converter id"))
		return
	}

	db, err := banco.Connection()
	if err != nil{
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("Delete from users WHERE id = ?")
	if err != nil{
		w.Write([]byte("Erro ao apagar ao usuario"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil{
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//Altera os dados de um usuario

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter parametros do ID"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler body da requisição"))
		return
	}

	var usuario usuario
	if err = json.Unmarshal(body, &usuario); err != nil{
		w.Write([]byte("Erro ao converter usuario em json"))
		return
	}

	db, err := banco.Connection()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil{
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(usuario.Name, usuario.Email, ID); err != nil{
		w.Write([]byte("Erro ao executar update"))
		return
	}


	w.WriteHeader(http.StatusNoContent)


}
