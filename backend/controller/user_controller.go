package controller

import (
	"dict/config"
	"dict/entity"
	"dict/helper"
	"dict/model"
	"dict/workflow"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService workflow.UserService
	authService config.AuthService
}

func NewUserController(userService workflow.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}

type UserFormatter struct {
	UserID uint   `json:"id"`
	Email  string `json:"email"`
	//Phone  string `json:"phone"`
	Token string `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter { //Token akan didapatkan dari JWT
	formater := UserFormatter{
		UserID: user.ID,
		Token:  token,
	}

	return formater
}

func (h *userController) Login(c *gin.Context) {
	var input entity.LoginEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responsError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	token, err := h.authService.GenerateTokenUser(loggedInUser.ID)
	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "success", FormatUser(loggedInUser, token))

	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateProfile(c *gin.Context) {
	var input entity.DataUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// fmt.Println(err.Error())
		// return
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Create Account Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	input.ID = c.MustGet("currentUser").(model.User).ID
	updateUser, err := h.userService.UpdateProfile(input)
	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", updateUser)
	c.JSON(http.StatusOK, response)
}
