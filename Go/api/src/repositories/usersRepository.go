package repositories

import ("database/sql"
	"api/src/models"
)

type UserRepository struct{
	db *sql.DB
}


func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) Create(user models.UserModel) (uint, error){
	statement, err := repo.db.Prepare("Insert into usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if err != nil{
		return 0, err 
	}

	defer statement.Close()

	res, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(lastID), nil
}