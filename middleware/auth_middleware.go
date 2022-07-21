package middleware

import (
	"context"
	"fmt"
	"net/http"
	"trackingApp/graph/models"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

var wCtxKey = &wContextKey{"w"}

type wContextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("token")

		//Adminov token
		//header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiIxIiwiZXhwaXJlc19BdCI6IjIwMjItMDMtMjlUMTE6MzU6NDUuMjM1MTc5OTQ1KzAyOjAwIiwiaXNzdWVkQXQiOjE2NDg0NjAxNDV9.GtDC_YjzUUoEQuk72_UAErwNM8xWsvnG9Lw44D8PWw8"

		//token(sezonac)
		//header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiIyIiwiZXhwaXJlc19BdCI6IjIwMjItMDMtMzBUMTA6MTY6NTguMDkxMTA1MDE4KzAyOjAwIiwiaXNzdWVkQXQiOjE2NDg1NDE4MTh9.9l58YEKdX-PzWqbBbgKZWhVceOVKOuqRXoo_-boJx8s"

		//Neprijavljen korisnik
		//header := ""

		if header == "" {
			fmt.Println("[INFO] Unauthenticated user")
			next.ServeHTTP(w, r)
			return
		}

		id, err := ParseToken(header)

		if err != nil {
			fmt.Println("cannot parse token")

			return
		}
		wCtx := context.WithValue(r.Context(), wCtxKey, &w)
		r = r.WithContext(wCtx)

		// create user and check if user exists in db
		userId := id
		user := models.User{ID: userId}

		// put it in context
		ctx := context.WithValue(r.Context(), userCtxKey, &user)

		fmt.Println("[INFO] Authenticated user")

		// and call the next with our new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {

	raw, _ := ctx.Value(userCtxKey).(*models.User)

	return raw
}

func WForContext(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(wCtxKey).(*http.ResponseWriter)
	return raw
}
