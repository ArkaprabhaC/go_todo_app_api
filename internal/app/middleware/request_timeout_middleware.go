package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/model/dto/errors"
	"github.com/gin-gonic/gin"
)

func RequestTimeout(timeout time.Duration, errTimeout *errors.RequestTimeoutError) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 3 * time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		finished := make(chan struct{})

		go func() {
			c.Next()
			close(finished)
		}()

		select {
		case <-ctx.Done():
			fmt.Println("HERE!")
			c.Abort()
			c.AbortWithStatusJSON(errTimeout.Code, errTimeout)
		case <-finished:
			fmt.Println("FINISHED!")
		}

	}
}