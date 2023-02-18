package controller

import (
	"app/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (int, error) {
	id, err := c.store.User.Create(req)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Controller) GetByIdUser(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User.GetById(req)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (c *Controller) UpdateUser(req *models.UpdateUser, id int) error {
	_, err := c.store.User.GetById(&models.UserPrimaryKey{Id: id})
	if err != nil {
		return err
	}
	err = c.store.User.Update(req, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteUser(req *models.UserPrimaryKey) error {
	err := c.store.User.Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetListUser(req *models.GetListRequest) (models.GetListResponse, error) {
	users, err := c.store.User.GetList(req)
	if err != nil {
		return models.GetListResponse{}, err
	}
	return users, nil
}
