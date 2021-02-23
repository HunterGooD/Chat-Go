package app

import (
	"crypto/rand"
	"net/http"

	"github.com/HunterGooD/Chat-Go/internal/utils"
	"github.com/HunterGooD/Chat-Go/pkg/crypto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func helloWorld(c *gin.Context) {
	type Response struct {
		Name string `json:"name"`
		Key  string `json:"key"`
		Docs string `json:"docs"`
	}
	obj := Response{
		Name: "Hello",
		Key:  "World",
		Docs: "/dev/api_docs",
	}
	c.JSON(http.StatusOK, obj)
}

// GenerateToken генерация токена
func GenerateToken() string {
	token := make([]byte, 24)
	rand.Read(token)
	return crypto.CreateHash(token)
}

func (a *App) signUp(c *gin.Context) {
	user := SignUpUser{}
	if err := c.Bind(&user); err != nil {
		errorJSON := &ErrorRequest{
			Error:   err.Error(),
			Message: "не удачная обработка данных",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		panic(err)
	} else if (user == SignUpUser{}) {
		errorJSON := &ErrorRequest{
			Error:   "Пустые значения полей",
			Message: "Заполните все поля",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		panic(err)
	}
	if !utils.ValidateUserData(map[string]string{
		"login":    user.Login,
		"name":     user.Name,
		"password": user.Password,
	}) {
		errorJSON := &ErrorRequest{
			Error:   "Ошибка",
			Message: "Значения не соответсвуют валидации",
			Errors: map[string]interface{}{
				"login":    utils.ValidLogin(user.Login),
				"name":     utils.ValidName(user.Name),
				"password": utils.ValidPassword(user.Password),
			},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		return
	}
	if err := a.db.Debug().Model(&UserDB{}).First(&UserDB{}).Where(map[string]string{
		"login": user.Login,
		"name":  user.Name,
	}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			Token := GenerateToken()
			pass, err := crypto.HashPassword(user.Password)
			if err != nil {
				errorJSON := &ErrorRequest{
					Error:   err.Error(),
					Message: "не удачная обработка данных",
					Errors:  nil,
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
				panic(err)
			}
			var userCreate *UserDB = &UserDB{
				Admin:    false,
				Login:    user.Login,
				Name:     user.Name,
				Password: pass,
				Token:    Token,
			}
			if err := a.db.Debug().Create(userCreate).Error; err != nil {
				errorJSON := &ErrorRequest{
					Error:   err.Error(),
					Message: "не удачная обработка данных",
					Errors:  nil,
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
				panic(err)
			}
			c.JSON(http.StatusCreated, userAuthResponse{
				Token: Token,
			})
		} else {
			errorJSON := &ErrorRequest{
				Error:   err.Error(),
				Message: "не удачная обработка данных",
				Errors:  nil,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
			panic(err)
		}
	} else {
		errorJSON := &ErrorRequest{
			Error:   "Пользователь должен быть уникальным",
			Message: "Аккаунт с таким логином уже существует",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
	}
}

func (a *App) signIn(c *gin.Context) {
	user := SignInUser{}
	if err := c.Bind(&user); err != nil {
		errorJSON := &ErrorRequest{
			Error:   err.Error(),
			Message: "не удачная обработка данных",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		panic(err)
	} else if (user == SignInUser{}) {
		errorJSON := &ErrorRequest{
			Error:   "Пустые значения полей",
			Message: "Заполните все поля",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		panic(err)
	}
	userLogIn := &UserDB{}
	if err := a.db.Debug().Model(&UserDB{}).First(userLogIn).Where(map[string]string{
		"login": user.Login,
	}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errorJSON := &ErrorRequest{
				Error:   err.Error(),
				Message: "Аккаунта с таким логином не существует",
				Errors:  nil,
			}
			c.AbortWithStatusJSON(http.StatusNotFound, errorJSON)
			panic(err)
		} else {
			panic(err)
		}
	}
	if !crypto.CheckPasswordHash(user.Password, userLogIn.Password) {
		errorJSON := &ErrorRequest{
			Error:   "Ошибка",
			Message: "Не правильно введен пароль",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		return
	}
	Token := GenerateToken()
	if err := a.db.Debug().Update("token", Token).Where(map[string]string{
		"login": user.Login,
	}).Error; err != nil {
		errorJSON := &ErrorRequest{
			Error:   err.Error(),
			Message: "не удачная обработка данных",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
		panic(err)
	}
	c.JSON(http.StatusOK, userAuthResponse{
		Token: Token,
	})
}
