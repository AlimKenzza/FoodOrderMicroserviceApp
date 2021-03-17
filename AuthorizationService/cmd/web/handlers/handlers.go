package handlers

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/foundation/rest_errors"
	"gitlab.com/AlimKenzza/authorization/pkg/data"
	"gitlab.com/AlimKenzza/authorization/utils"
	"net/http"
	"os"
	"time"
)

var jwtKey = []byte("secret")

//Claims jwt claims struct
type Claims struct {
	data.User
	jwt.StandardClaims
}

func Create(c *gin.Context) {
	var user data.Register
	c.Bind(&user)
	exists := checkUserExists(user)

	valErr := utils.ValidateUser(user, rest_errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "email already exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	data.HashPassword(&user)
	_, err := data.DB.Query(data.CreateUserQuery, user.Username, user.Password, user.Email)
	rest_errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created succesfully"})
}

func checkUserExists(user data.Register) bool {
	rows, err := data.DB.Query(data.CheckUserExists, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

func Login(c *gin.Context) {
	var user data.Login
	c.Bind(&user)

	row := data.DB.QueryRow(data.LoginQuery, user.Username)

	var id int64
	var name, email, password string

	err := row.Scan(&id, &name, &password, &email)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	match := data.CheckPasswordHash(user.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{

		User: data.User{
			Id: id, Username: name, Email: email,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, err := token.SignedString(jwtKey)
	rest_errors.HandleErr(c, err)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "logged in succesfully", "user": claims.User, "token": tokenString})
}

func LoginUser(c *gin.Context) {
	var userr data.Login
	c.Bind(&userr)

	row := data.DB.QueryRow(data.LoginQuery, userr.Username)

	var id int
	var name, email, password string

	err := row.Scan(&id, &name, &password, &email)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	match := data.CheckPasswordHash(userr.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	token, err := CreateToken(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	c.JSON(http.StatusOK, id)
}

func CreateToken(id int) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
