package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pockethealth/internchallenge/pkg/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var LISTENING_PORT = "80" //Q2: modify the server to run on port 80 from 8080
var ALLOWED_ORIGINS = []string{"http://localhost:4200", "https://localhost:4200", "https://localhost", "https://127.0.0.1:4200", "http://127.0.0.1:4200", "https://127.0.0.1"}
var ALLOWED_HEADERS = []string{"Accept", "Content-type"}
var ALLOWED_METHODS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
var EXPOSED_HEADERS = []string{"Content-disposition", "cache-control", "content-length", "expires", "pragma"}

// Q3 & Q1 Issues, built a registration template code in Golang
type User struct {
	UserName      string `json:"username" bson:"username"`
	Email         string `json:"email" bson:"email"`
	FavoriteColor string `json:"favorite_color" bson:"favorite_color"` //Part2 Q1: Store user's favourate color
}

// Question 3: Return the ID
// UserApiController handles HTTP requests for the User API
type UserApiController struct {
	userService UserApiService
}

// NewUserApiController returns a new instance of the UserApiController
func NewUserApiController(userService UserApiService) *UserApiController {
	return &UserApiController{
		userService: userService,
	}
}

// RegisterUserRequest represents the request to register a new user
type RegisterUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

// RegisterUserResponse represents the response from registering a new user
type RegisterUserResponse struct {
	UserID string `json:"user_id"`
}

// registerUserHandler is the handler for the POST /register route
func (c *UserApiController) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// register new user
	userID, err := c.userService.RegisterUser(req.UserName, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response
	res := RegisterUserResponse{
		UserID: userID,
	}

	// encode response as JSON
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Routes returns the HTTP routes for the User API
func (c *UserApiController) Routes() []*Route {
	return []*Route{
		{
			Name:        "RegisterUser",
			Method:      "POST",
			Pattern:     "/register",
			HandlerFunc: c.registerUserHandler,
		},
	}
}

// Route represents an HTTP route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// UserApiService represents a service that interacts with the User API
type UserApiService interface {
	RegisterUser(userName string, email string) (string, error)
}

// UserService represents an implementation of the UserApiService interface
type UserService struct {
	// we can add fields here to store dependencies such as a database connection
}

// NewUserApiService returns a new instance of the UserService
func NewUserApiService() *UserService {
	return &UserService{}
}

// RegisterUser registers a new user with the given name and email
func (s *UserService) RegisterUser(userName string, email string) (string, error) {
	// generate user ID
	userID := "123456" // replace this with your own code to generate a unique user ID
	// return user ID
	return userID, nil
}

// mux router
func createRouter() *mux.Router {
	// configure auth controller and service to handle requests
	UserApiService := user.NewUserApiService()
	UserApiController := user.NewUserApiController(UserApiService)

	// initialize routes
	mRouter := mux.NewRouter().StrictSlash(false)
	//add paths to subrouter
	for _, route := range UserApiController.Routes() {
		mRouter.
			Methods(route.Method, http.MethodOptions).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// Q1: allowing anyone to register for access to PocketHealth
	// add registration route

	//anyone can register for access to the web application by
	//sending a POST request to /register with the required user information
	mRouter.
		Methods(http.MethodPost).
		Path("/register").
		HandlerFunc(mRouter.ServeHTTP) //Controller: user-who interact with system provide information
		// add registration route

	mRouter.
		Methods(http.MethodPost).
		Path("/register").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Controller: user-who interact with system provide information
			name := r.FormValue("username")
			//Email := r.FormValue("email")
			userID := r.FormValue("userid")

			// redirect to home page
			http.Redirect(w, r, fmt.Sprintf("/home?name=%s&userid=%s", name, userID), http.StatusSeeOther)
		})

	// Question 4: add home route
	mRouter.
		Methods(http.MethodGet).
		Path("/home").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get query parameters
			name := r.FormValue("name")
			userID := r.FormValue("userid")

			// write response
			fmt.Fprintf(w, "Welcome to PocketHealth %s. Your User ID is: %s", name, userID)
		})

	return mRouter
}

func main() {
	// configure router
	router := createRouter()

	// configure request handler
	n := negroni.New()
	n.UseHandler(router)

	// CORS config
	handler := handlers.CORS(
		handlers.AllowedOrigins(ALLOWED_ORIGINS),
		handlers.AllowedHeaders(ALLOWED_HEADERS),
		handlers.AllowedMethods(ALLOWED_METHODS),
		handlers.ExposedHeaders(EXPOSED_HEADERS),
	)(n)

	// setup and start the server
	server := &http.Server{Addr: ":" + LISTENING_PORT, Handler: handler}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
