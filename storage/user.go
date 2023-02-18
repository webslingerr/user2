package storage

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// change userRepo
// change fileName

type userRepo struct {
	fileName string
}

func NewUserRepo(fileName string) *userRepo {
	return &userRepo{
		fileName: fileName,
	}
}

// READ A FILE
func (u *userRepo) Read() ([]models.User, error) {
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return []models.User{}, err
	}
	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// CREATE
func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {
	users, err := u.Read()
	if err != nil {
		return 0, err
	}

	id = len(users) + 1
	users = append(users, models.User{
		Id:      id,
		Name:    req.Name,
		Surname: req.Surname,
	})

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GETBYID
func (u *userRepo) GetById(req *models.UserPrimaryKey) (models.User, error) {
	users, err := u.Read()
	if err != nil {
		return models.User{}, err
	}

	for _, val := range users {
		if val.Id == req.Id {
			return val, nil
		}
	}
	return models.User{}, errors.New("There is no user with this id")
}

// UPDATE
func (u *userRepo) Update(req *models.UpdateUser, id int) error {
	users, err := u.Read()
	if err != nil {
		return err
	}

	for i, v := range users {
		if v.Id == id {
			users[i].Name = req.Name
			users[i].Surname = req.Surname
		}
	}

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// DELETE
func (u *userRepo) Delete(req *models.UserPrimaryKey) error {
	users, err := u.Read()
	if err != nil {
		return err
	}

	for i, v := range users {
		if v.Id == req.Id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return nil
}

// GETLIST
func (u *userRepo) GetList(req *models.GetListRequest) (models.GetListResponse, error) {
	users, err := u.Read()
	if err != nil {
		return models.GetListResponse{}, err
	}

	newUsers := []models.User{}
	if req.Limit+req.Offset > len(users) {
		return models.GetListResponse{}, err
	}
	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		newUsers = append(newUsers, users[i])
	}
	return models.GetListResponse{
		Count: len(newUsers),
		Users: newUsers,
	}, nil
}
