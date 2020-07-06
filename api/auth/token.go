package auth

import "net/http"

func CreateToken(uid uint32) (string, error) {

}

func TokenValid(r *http.Request) error {

}

func ExtractToken(r *http.Request) string {

}

func ExtractTokenID(r *http.Request) (uint32, error) {

}

func Pretty(data interface{}) {

}
