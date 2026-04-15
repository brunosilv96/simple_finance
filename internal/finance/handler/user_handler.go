package handler

import (
	"errors"
	"net/http"

	"github.com/brunosilv96/simple_finance_api/internal/finance/dto"
	usecaseError "github.com/brunosilv96/simple_finance_api/internal/finance/usecase"
	usecase "github.com/brunosilv96/simple_finance_api/internal/finance/usecase/user"
)

type UserHandler struct {
	RegisterUserUC usecase.RegisterUser
	FindUserByIdUC usecase.FindUserByID
}

func NewUserHandler(repository usecase.UserRepository) *UserHandler {
	registerUserUseCase := usecase.NewRegisterUser(repository)
	findUserUseCase := usecase.NewFindUserByID(repository)

	return &UserHandler{
		RegisterUserUC: *registerUserUseCase,
		FindUserByIdUC: *findUserUseCase,
	}
}

func (handler *UserHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/user", handler.Create)
	mux.HandleFunc("GET /api/v1/user/{id}", handler.FindByID)
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterUserRequest

	err := input.Bind(r.Body)
	if err != nil {
		dto.ErrorReponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := handler.RegisterUserUC.Execute(input.Name)
	if err != nil {
		dto.ErrorReponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	userResponse.Render(w)
}

func (handler *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		dto.ErrorReponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	user, err := handler.FindUserByIdUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, usecaseError.UserNotFound):
			dto.ErrorReponse(w, http.StatusNotFound, err.Error())
			return
		default:
			dto.ErrorReponse(w, http.StatusInternalServerError, "Server Internal Error")
			return

		}
	}

	userResponse := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	userResponse.Render(w)
}
