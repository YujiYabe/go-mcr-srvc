package database

import "app/domain"


type UserRepository struct {
    SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
    result, err := repo.Execute(
        "INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
    )
    if err != nil {
        return
    }
    id64, err := result.LastInsertId()
    if err != nil {
        return
    }
    id = int(id64)
    return
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
    if err = repo.Find(&user, id).Error; err != nil {
        return
    }
    return
}
func (repo *UserRepository) FindAll() (users domain.Users, err error) {
    rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
    defer rows.Close()
    if err != nil {
        return
    }
    for rows.Next() {
        var id int
        var firstName string
        var lastName string
        if err := rows.Scan(&id, &firstName, &lastName); err != nil {
            continue
        }
        user := domain.User{
            ID:        id,
            FirstName: firstName,
            LastName:  lastName,
        }
        users = append(users, user)
    }
    return
}
