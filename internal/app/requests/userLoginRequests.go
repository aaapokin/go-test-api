package requests

import (
	"app/internal/app"
	"app/internal/app/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	context  *gin.Context
	user     *models.User
	userInDB *models.User
}

func NewUserLoginRequest(context *gin.Context, user *models.User) *UserLoginRequest {
	return &UserLoginRequest{user: user, context: context}
}

func (request *UserLoginRequest) validate() error {
	if err := request.context.BindJSON(request.user); err != nil {
		badResponseByError(request.context, err)
		return err
	}
	var ok bool
	var err error
	request.userInDB, ok, err = app.Store.User().FindByLogin(request.user.Login)
	if err != nil {
		badResponseByError(request.context, err)
		return err
	}
	if !ok {
		err := errors.New("user not found")
		badResponseByError(request.context, err)
		return err
	}
	if request.userInDB.Password != request.user.Password {
		err := errors.New("your password is invalid")
		badResponseByError(request.context, err)
		return err
	}

	return nil
}

func (request *UserLoginRequest) GetDTO() (*models.User, error) {
	err := request.validate()
	if err != nil {
		badResponseByError(request.context, err)
		return nil, err
	}
	return request.userInDB, nil
}
