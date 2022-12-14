package handler_error

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

type HttpError struct {
}

func NewHttpError() *HttpError {
	return &HttpError{}
}

func (h *HttpError) NewError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
