package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	pbu "github.com/barber_shop/api-gateway/genproto/users_service"

	l "github.com/barber_shop/api-gateway/pkg/logger"
	// "github.com/barber_shop/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	IssuedAt  string `json:"issued_at"`
	ExpiredAt string `json:"expired_at"`
}

const (
	authorizationHeaderKey  = "authorization"
	authorizationPayloadKey = "authorization_payload"
)

func (h *handlerV1) AuthMiddleware(resourse, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(authorizationHeaderKey)
		fmt.Println(c.Request.URL.Path)
		if len(accessToken) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		
		payload, err := h.serviceManager.StaffAuthService().VerifyToken(context.Background(), &pbu.VerifyTokenRequest{
			AccessToken: accessToken,
			Resource:    resourse,
			Action:      action,
		})
		if err != nil {
			h.log.Error("failed to verify token", l.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		if !payload.HasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(ErrNotAllowed))
			return
		}

		c.Set(authorizationHeaderKey, Payload{
			ID:        payload.Id,
			UserID:    payload.UserId,
			Email:     payload.Email,
			Type:      payload.Type,
			IssuedAt:  payload.IssuedAt,
			ExpiredAt: payload.ExpiredAt,
		})
		c.Next()
	}

}

func (m *handlerV1) GetAuthPayload(c *gin.Context) (*Payload, error) {
	i, exists := c.Get(authorizationHeaderKey)
	if !exists {
		return nil, errors.New("-<*>_<*>-")
	}

	fmt.Println("///////////////////////////////////")
	fmt.Println(i)
	fmt.Println("///////////////////////////////////")

	payload, ok := i.(Payload)
	if !ok {
		return nil, errors.New("unknown user")
	}
	return &payload, nil
}
