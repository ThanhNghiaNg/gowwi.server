package api

import (
	"os"
	"owwi/pkg/models"
	"owwi/pkg/repositories"
	"owwi/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description User login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.UserLogin true "Login data"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /login [post]
func login(c *gin.Context) {
	var loginData models.UserLogin
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user, err = repositories.UserRepository.GetUserByUsername(loginData.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	var (
		token *jwt.Token
		s     string
	)

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"role":     utils.If(user.IsAdmin, "admin", "user"),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	s, err = token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// gen jwt token
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   s,
	})
}


// @Summary Register
// @Description User register with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body models.UserRegister true "register data"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /register [post]
func register(c *gin.Context) {
	var register models.UserRegister
	if err := c.BindJSON(&register); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var existUser, err = repositories.UserRepository.GetUserByUsername(register.Username)
	if err == nil && existUser != nil {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid bcrypt cost"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), cost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	var userData = models.CreateUser{
		Username: register.Username,
		Password: string(hashedPassword),
		Email:    register.Email,
		IsAdmin:  false,
	}

	userId, err := repositories.UserRepository.CreateUser(userData)

	if err := repositories.TypeRepository.CreateType(models.Type{
		User:        userId,
		Name:        "Income",
		Description: "Default income type",
	}); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create default income type"})
	}

	if err := repositories.TypeRepository.CreateType(models.Type{
		User:        userId,
		Name:        "Outcome",
		Description: "Default outcome type",
	}); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create default outcome type"})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	var (
		token *jwt.Token
		s     string
	)

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": register.Username,
		"user_id":  userId.Hex(),
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	s, err = token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// gen jwt token
	c.JSON(200, gin.H{
		"message": "Register successful",
		"token":   s,
	})
}

var AuthApi = struct {
	Login    func(c *gin.Context)
	Register func(c *gin.Context)
}{
	Login:    login,
	Register: register,
}
