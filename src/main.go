package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	auth "backend/src/auth"
	dev "backend/src/device"
	user "backend/src/user"
)

type AuthRequest struct {
	ApiKey string `json:"api_key"`
	Email  string `json:"email"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/auth", AuthHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/signin", SignInHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/signup", SignUpHandler).Methods("POST", "OPTIONS")
	r.Use(CorsMiddleware)

	sb := r.NewRoute().Subrouter()
	sb.HandleFunc("/api/device/info", InfoHandler).Methods("POST", "OPTIONS")
	sb.HandleFunc("/api/device/info/list", ListInfoHandler).Methods("GET", "OPTIONS")
	sb.Use(CorsMiddleware)
	sb.Use(AuthMiddleware)

	http.Handle("/", r)

	log.Print("Server gonna work at 9090 port!")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var a AuthRequest
	var ar AuthResponse

	// Read body of request and try convert into AuthRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	// Try authenticate with information
	token, err := auth.Auth(a.ApiKey, a.Email)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	ar = AuthResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ar)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var a SignInRequest
	var ar AuthResponse

	// Read body of request and try convert into AuthRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	// Try authenticate with information
	token, err := auth.AuthWithPassword(a.Email, a.Password)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	ar = AuthResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ar)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var a SignInRequest

	// Read body of request and try convert into AuthRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	// Try authenticate with information
	_, err = user.CreateNewUser(a.Email, a.Password)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var i dev.Info

	user_id := r.Context().Value("user_id")

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		JSONError(w, errors.New("error converting data"), http.StatusBadRequest)
		return
	}

	err = dev.CreateNewInfo(user_id.(int), i)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListInfoHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)

	params := mux.Vars(r)
	deviceName := params["name"]

	info, err := dev.ListDeviceInfo(user_id, deviceName)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// Check if token is valid
		isAuth, userId := auth.IsAuthenticated(token)

		if isAuth {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_id", userId)
			rcopy := r.WithContext(ctx)

			next.ServeHTTP(w, rcopy)
		} else {
			JSONError(w, errors.New("forbidden"), http.StatusForbidden)
		}
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	er := ErrorResponse{
		Error: err.Error(),
	}

	json.NewEncoder(w).Encode(er)
}
