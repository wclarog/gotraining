package library

import (
	"context"
	"errors"

	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/dgrijalva/jwt-go"
	kitTransport "github.com/go-kit/kit/transport/http"
)

type (
	AuthInfo struct {
		Role   Role
		UserId string
	}
	Role int8
	User struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  Role   `json:"role"`
	}
)

const (
	authHeader              = "Authorization"
	JWTTokenContextKey      = "JWT_TOKEN"
	CtxKey                  = "USER_CTX_KEY"
	RoleUserAdmin      Role = iota + 1
	RoleUserClient
)

var (
	ErrTokenInvalid = errors.New(" invalid token")
)

func AccessControl(name string, role Role) bool {
	access := make(map[string][]Role)
	access["AddMaterialEndpoint"] = []Role{RoleUserAdmin}
	access["UpdateMaterialEndpoint"] = []Role{RoleUserAdmin, RoleUserClient}

	if len(access[name]) > 0 {
		for _, v := range access[name] {
			if v == role {
				return true
			}
		}
	}
	return false

}

func NewAuthMiddleware(e Endpoints) Endpoints {
	return Endpoints{
		AddMaterialEndpoint:        newAuthMiddleware("AddMaterialEndpoint")(e.AddMaterialEndpoint),
		UpdateMaterialEndpoint:     newAuthMiddleware("UpdateMaterialEndpoint")(e.UpdateMaterialEndpoint),
		DeleteMaterialEndpoint:     e.DeleteMaterialEndpoint,
		GetMaterialsEndpoint:       e.GetMaterialsEndpoint,
		GetMaterialByCodeEndpoint:  e.GetMagazineByCodeEndpoint,
		GetBooksEndpoint:           e.GetBooksEndpoint,
		GetBookByCodeEndpoint:      e.GetBookByCodeEndpoint,
		GetMagazinesEndpoint:       e.GetMagazinesEndpoint,
		GetMagazineByCodeEndpoint:  e.GetMagazineByCodeEndpoint,
		GetNewspapersEndpoint:      e.GetNewspapersEndpoint,
		GetNewspaperByCodeEndpoint: e.GetNewspaperByCodeEndpoint,
	}
}

func newAuthMiddleware(name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			token, ok := ctx.Value(JWTTokenContextKey).(string)
			if !ok {
				return nil, ErrTokenInvalid
			}

			ai, err := GetAuthInfoFromJwtToken(token)
			if err != nil {
				return nil, err
			}

			// user service get User
			if !AccessControl(name, ai.Role) {
				return nil, ErrInvalidAccess
			}

			ctx = ToContext(ctx, User{
				Role: ai.Role,
				Id:   ai.UserId,
			})

			return next(ctx, request)
		}
	}
}

func InsertJwtIntoContext() kitTransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		token := r.Header.Get(authHeader)
		if token == "" {
			return ctx
		}
		return context.WithValue(ctx, JWTTokenContextKey, token)
	}
}

func GetAuthInfoFromJwtToken(jwtToken string) (AuthInfo, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return AuthInfo{}, ErrTokenInvalid
	}

	claims := token.Claims.(jwt.MapClaims)

	role, ok := claims["userRole"]
	if !ok {
		return AuthInfo{}, ErrTokenInvalid
	}

	userId, ok := claims["userId"]
	if !ok {
		return AuthInfo{}, ErrTokenInvalid
	}

	return AuthInfo{
		Role:   role.(Role),
		UserId: fmt.Sprintf("%v", userId),
	}, nil
}

func ToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, CtxKey, val)
}

func FromContext(ctx context.Context) (User, error) {
	usr, ok := ctx.Value(CtxKey).(User)
	if !ok {
		return User{}, ErrTokenInvalid
	}
	return usr, nil
}
