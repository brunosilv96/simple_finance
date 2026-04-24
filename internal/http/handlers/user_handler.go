package http

import (
	"errors"
	"net/http"

	domain "github.com/brunosilv96/simple_finance_api/internal/finance/user"
	usecase "github.com/brunosilv96/simple_finance_api/internal/finance/user/usecase"
	"github.com/brunosilv96/simple_finance_api/internal/http/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	RegisterUserUC usecase.RegisterUser
	FindUserByIdUC usecase.FindUserByID
}

func NewUserHandler(repository domain.UserRepository) *UserHandler {
	registerUserUseCase := usecase.NewRegisterUser(repository)
	findUserUseCase := usecase.NewFindUserByID(repository)

	return &UserHandler{
		RegisterUserUC: *registerUserUseCase,
		FindUserByIdUC: *findUserUseCase,
	}
}

func (handler *UserHandler) Create(c *gin.Context) {
	var input dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	user, err := handler.RegisterUserUC.Execute(input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userResponse := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (handler *UserHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	user, err := handler.FindUserByIdUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.UserNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return

		}
	}

	userResponse := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}
