package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in auth middleware")
		//set ctx timeout
		ctx, cancle := context.WithTimeout(r.Context(), time.Millisecond * 1)

		defer cancle()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
