package controller

import (
	middlewares "ExpencesManagment/Middlewares"
	models "ExpencesManagment/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserAuth struct {
	Username string
	Password string
}

// CreateExpenses godoc
// @Summary Create user
// @Description Create an new user
// @Tags user
// @Accept json
// @Produce json
// @Param body body UserAuth true "New user"
// @Router /auth/register [post]
func RegisterUser(c *gin.Context) {
	var input UserAuth
	var username []models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := models.User{
		UserName: input.Username,
		Password: string(HashedPassword),
	}

	if err := db.Where("user_name = ?", input.Username).First(&username).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username is not avalable"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to create user"})
		return
	}

	c.JSON(http.StatusOK, "User succesfully Registed")
}

// authenticate godoc
// @Summary authenticate the user
// @Description authenticate the user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body UserAuth true "Login user"
// @Router /auth/Login [post]
func LoginUser(c *gin.Context) {
	var input UserAuth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := db.Where("user_name = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err :=middlewares.GenerateJWT(user.UserID,user.UserName)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"user login succesfully",
	"token": token })
}

