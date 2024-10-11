package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto/errors"
	"github.com/gin-gonic/gin"
)

func RequestTimeout(timeout time.Duration, errTimeout *errors.AppError) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		finished := make(chan struct{})

		go func() {
			c.Next()
			close(finished)
		}()

		select {
		case <-ctx.Done():
			c.Abort()
			c.AbortWithStatusJSON(errTimeout.Code, errTimeout)
		case <-finished:
			fmt.Println("FINISHED!")
		}

	}
}