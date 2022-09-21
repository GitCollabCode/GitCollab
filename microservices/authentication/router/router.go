package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("auth stuff"))
	})

	r.Get("/testing", func(w http.ResponseWriter, r *http.Request) {
		// TODO: proxy will be redirected to this, create jwt and send to frontend,
		// 		 do db stuff for new or existing user
		// create new JWT, extract github id, access token from code
		//SECRET := os.Getenv("SECRET")
		//tokenAuth := jwtauth.New("HS256", []byte(SECRET), nil)
		//_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		w.Write([]byte("cream cheese"))
	})

	return r
}
