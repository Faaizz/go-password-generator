package router

import (
	"encoding/json"
	"net/http"

	"github.com/faaizz/go-password-generator/business"
	"github.com/julienschmidt/httprouter"
)

func GeneratePwd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, err := business.ParseParams(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	pwds, err := business.GetPwds(p.MinLength, p.SpecialCharsCount, p.NumbersCount, p.PwdsToCreate)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(pwds)
}

func Start() error {
	router := httprouter.New()
	router.POST("/", GeneratePwd)

	return http.ListenAndServe(":8080", router)
}
