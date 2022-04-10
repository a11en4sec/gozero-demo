package middleware

import (
	"fmt"
	"net/http"
)

type GlobalMiddleware struct {

}

func NewGlobalMiddleware() *GlobalMiddleware {
	return &GlobalMiddleware{}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// before
		fmt.Println("[+] before.")
		
		next(w,r)
		
		// after
		fmt.Println("[+] after.")
	}
}