package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
"github.com/golang-jwt/jwt"
	// "bytes"
	// "io"
	// "strconv"
	// "github.com/iqra-shams/chi/cmd"
)


var (
	secretKey = []byte("secret -key")
)

var user = map[string]string{
	"user": "password",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Token string `json:"token"`
}
type ResponseData struct {
	Lines         int    `json:"Number_of_Lines"`
	Words         int    `json:"Number_of_Words"`
	Puncuations   int    `json:"Number_of_Puncuations"`
	Vowels        int    `json:"Number_of_Vowels"`
	ExecutionTime string `json:"ExecutionTime"`
	Routines      int    `json:"Number_of_Routines"`
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := user[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ExpireTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(),
		},
	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: ExpireTime,
	})

	// Send the token in the JSON response
	response := LoginResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func Restricted(w http.ResponseWriter, r *http.Request){
	cookie, err:= r.Cookie("token")
	if err != nil {
		if err== http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims:= &Claims{}
	tkn,err :=jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil{
		if err==jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
w.WriteHeader(http.StatusBadRequest)
return
	}
	if !tkn.Valid{
		w.WriteHeader(http.StatusUnauthorized)
return
	}
	w.Write([]byte(fmt.Sprintf("hello,%s",claims.Username)))
	// start := time.Now()

	// // err := r.ParseMultipartForm(1000 << 20)
	// // if err != nil {
	// // 	http.Error(w, "fail to parse form", http.StatusInternalServerError)
	// // }
	// Sroutines := r.FormValue("routines")

	// file, _, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, "fail to get  file", http.StatusInternalServerError)
	// 	return
	// }

	// defer file.Close()

	// var BufFile bytes.Buffer

	// _, err = io.Copy(&BufFile, file)
	// if err != nil {
	// 	http.Error(w, "fail to read file", http.StatusInternalServerError)
	// 	return
	// }

	// routines, _ := strconv.Atoi(Sroutines)
	// result := cmd.ProcessFile(BufFile.String(), routines)
	// fmt.Println(result)

	// executiontime := time.Since(start).String()

	// responseData := ResponseData{
	// 	Lines:         result.LineCount,
	// 	Words:         result.WordsCount,
	// 	Vowels:        result.VowelsCount,
	// 	Puncuations:   result.PuncuationsCount,
	// 	ExecutionTime: executiontime,
	// 	Routines:      routines,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(responseData)
	// fmt.Printf("Execution time: %v \n", executiontime)
	
	
}

