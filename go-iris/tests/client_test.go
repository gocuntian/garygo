package gotest

import (
	"testing"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/httptest"
)

type Response struct {
	Text string `json:"text"`
}

func TestBasicJwt(t *testing.T) {
	var (
		api             = iris.New()
		myJwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("Y!w@aUjgGmXb@T$syd^j6];"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
	)

	api.Adapt(httprouter.New())
	securedPingHandler := func(ctx *iris.Context) {
		userToken := myJwtMiddleware.Get(ctx)
		var claimTestedValue string
		if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
			claimTestedValue = claims["foo"].(string)
		} else {
			claimTestedValue = "Claims Failed"
		}
		fmt.Println(claimTestedValue)
		response := Response{"Iauthenticated" + claimTestedValue}
		// get the *jwt.Token which contains user information using:
		// user:= myJwtMiddleware.Get(ctx) or context.Get("jwt").(*jwt.Token)

		ctx.JSON(iris.StatusOK, response)
	}

	api.Get("/secured/ping", myJwtMiddleware.Serve, securedPingHandler)


	e := httptest.New(api, t)

	e.GET("/secured/ping").Expect().Status(iris.StatusUnauthorized)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("Y!w@aUjgGmXb@T$syd^j6];"))

	e.GET("/secured/ping").WithHeader("Authorization", "Bearer "+tokenString).

    Expect().Status(iris.StatusOK).Body().Contains("Iauthenticated").Contains("bar")

}