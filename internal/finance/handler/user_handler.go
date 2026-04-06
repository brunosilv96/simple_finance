package handler

import (
	"net/http"
	"simple_finance/internal/finance/dto"
	usecase "simple_finance/internal/finance/usecase/user"
)

type UserHandler struct {
	RegisterUserUC usecase.RegisterUser
}

func NewUserHandler(repository usecase.UserRepository) *UserHandler {
	registerUserUseCase := usecase.NewRegisterUser(repository)

	return &UserHandler{
		RegisterUserUC: *registerUserUseCase,
	}
}

func (handler *UserHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/user", handler.Create)
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
	
	userResponse := &dto.RegisterUserResponse{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	userResponse.Render(w)
}
