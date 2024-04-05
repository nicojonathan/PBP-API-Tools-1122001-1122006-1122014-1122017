package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Bebasapasaja123!")
var tokenName = "token"

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, username string) {
	tokenExpiryTime := time.Now().Add(15 * time.Minute)

	// Create claims with user data
	claims := &Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	// Encrypt claim to jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})

	userId := strconv.Itoa(claims.ID)
	values := []string{signedToken, userId}
	
	valuesJSON,_ := json.Marshal(values)

	// Save token to Redis
	InitializeRedisClient()
	ctx := context.Background()
	expirationDuration := time.Until(tokenExpiryTime)
	err = client.Set(ctx, username, valuesJSON, expirationDuration).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r)
		if !isValidToken {
			sendResponse(w, http.StatusUnauthorized, "Unauthorized")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(r *http.Request) bool {
	isAccessTokenValid, id, email := validateTokenFromCookies(r)
	fmt.Print(id, email, isAccessTokenValid)

	return false
}

func validateTokenFromCookies(r *http.Request) (bool, int, string) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Username
		}
	}
	return false, -1, ""
}
