package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"github.com/iqra-shams/chi/cmd"
)

var jwtkey = []byte("secret-key")
type ResponseData struct {
	Lines         int    `json:"Number_of_Lines"`
	Words         int    `json:"Number_of_Words"`
	Puncuations   int    `json:"Number_of_Puncuations"`
	Vowels        int    `json:"Number_of_Vowels"`
	ExecutionTime string `json:"ExecutionTime"`
	Routines      int    `json:"Number_of_Routines"`
}

func HandlerPostReq(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	err := r.ParseMultipartForm(1000 << 20)
	if err != nil {
		http.Error(w, "fail to parse form", http.StatusInternalServerError)
	}
	Sroutines := r.FormValue("routines")

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "fail to get  file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	var BufFile bytes.Buffer

	_, err = io.Copy(&BufFile, file)
	if err != nil {
		http.Error(w, "fail to read file", http.StatusInternalServerError)
		return
	}

	routines, _ := strconv.Atoi(Sroutines)
	result := cmd.ProcessFile(BufFile.String(), routines)
	fmt.Println(result)

	executiontime := time.Since(start).String()

	responseData := ResponseData{
		Lines:         result.LineCount,
		Words:         result.WordsCount,
		Vowels:        result.VowelsCount,
		Puncuations:   result.PuncuationsCount,
		ExecutionTime: executiontime,
		Routines:      routines,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
	fmt.Printf("Execution time: %v \n", executiontime)
}

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Missing authorization header")
// 		return
// 	}
// 	tokenString = tokenString[len("Bearer "):]

// 	err := verifyToken(tokenString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Invalid token")
// 		return
// 	}

// 	fmt.Fprint(w, "Welcome to the the protected area")

// }

// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return jwtkey, nil
// 		})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}

// 	return nil
// }

// package api

// import (
// 	"fmt"

// 	"net/http"

// 	"github.com/golang-jwt/jwt"
// )


// var (
// 	secretKey = []byte("secret -key")
// )
// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Missing authorization header")
// 		return
// 	}
// 	tokenString = tokenString[len("Bearer "):]

// 	err := verifyToken(tokenString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Invalid token")
// 		return
// 	}
// 	fmt.Fprint(w, "Welcome to the the protected area")

// }

// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if !token.Valid {
// 		return fmt.Errorf("Invalid token")
// 	}
// 	return nil
// }