package main

import (
    "fmt"
    "time"
    "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var JWTSecretKey = []byte("Y!w@aUjgGmXb@T$syd^j6];")

type requestResponse struct{
    Status int
    Message string
    Token string
}

func myHandler(ctx *iris.Context){
    //user:=ctx.Get("user").(*jwt.Token)
    ctx.Writef("This is an authenticated request\n")
    ctx.Writef("Claim content:\n")
   // ctx.Writef("%s",user.Signature)
}

func authUser(ctx *iris.Context){
         response  :=  requestResponse{}
         token:=jwt.New(jwt.SigningMethodHS256)
         claims:=make(jwt.MapClaims)
         claims["userId"] = 101
         claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*360)).Unix()
         claims["iat"] = time.Now().Unix()
         token.Claims = claims
         tokenString, err := token.SignedString(JWTSecretKey)
		if err != nil{
			ctx.JSON(iris.StatusInternalServerError,"error in generating token")
			return
		}
		response.Token = tokenString
		response.Status = 1
		response.Message = "success"
		ctx.JSON(iris.StatusOK,response)
}

type Response struct {
	Text string `json:"text"`
}

func PingHandler(ctx *iris.Context) {
	response := Response{"All good. You don't need to be authenticated to call this"}
	ctx.JSON(iris.StatusOK, response)
}

func SecuredPingHandler(ctx *iris.Context) {
	response := Response{"All good. You only get this message if you're authenticated"}
	// get the *jwt.Token which contains user information using:
	// user:= myJwtMiddleware.Get(ctx) or context.Get("jwt").(*jwt.Token)
    fmt.Println(ctx)
	ctx.JSON(iris.StatusOK, response)
}

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    
    jwtHandler:=jwtmiddleware.New(jwtmiddleware.Config{
        ValidationKeyGetter:func(token *jwt.Token)(interface{},error){
            return []byte("Y!w@aUjgGmXb@T$syd^j6];"),nil
        },
        SigningMethod:jwt.SigningMethodHS256,
    })

     app.Get("/auth/genToken",authUser)
     app.Get("/ping", PingHandler)
 
	app.Get("/secured/ping", jwtHandler.Serve, SecuredPingHandler)

    app.Use(jwtHandler)
    app.Get("/pings",myHandler)
    app.Listen(":8088")
}