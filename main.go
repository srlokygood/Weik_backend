package main

import (
	"net/http"
	"strconv"
	debug "weik/debug"
	estruct "weik/estructuras"
	petic "weik/peticiones"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//no encontrado code 1
var code1 = map[string]int{
	"code": 1,
}

//no guardado code 2
var code2 = map[string]int{
	"code": 2,
}

//Guardado exitoso code 3
var code3 = map[string]int{
	"code": 3,
}

//Permisos invalidos code 4
var code4 = map[string]int{
	"code": 4,
}

//ya existente code 5
var code5 = map[string]int{
	"code": 5,
}

type jwtCustomClaims struct {
	Pass string `json:"pass"`
	jwt.StandardClaims
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8100"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	r := e.Group("")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	//e.POST("/acces", acces)
	e.POST("/celular", gcelular)
	e.POST("/login", glogin)
	e.POST("/register", gregister)
	e.GET("/negocios", gNegocios)
	e.GET("/negocio/:i", gNegocio)
	e.GET("/tipserv", gTipServ)
	e.GET("/servicios", gServicios)
	e.GET("/servicios2", gServicios2)
	e.GET("/servicio/:i", gServicio)
	e.GET("/menuop", gMenuop)
	e.GET("/pyr/:i", gpyr)
	e.GET("/pyr2/:i", gpyr2)
	e.GET("/rating/:i", grating)
	e.GET("/codpas", gcodpas)
	e.GET("/busca/:i", gbusca)

	e.Logger.Fatal(e.Start(":8080"))
}

/*func acces(c echo.Context) error {
	u := &estruct.Usuarios{}
	if err := c.Bind(u); err != nil {
		debug.Menserr(err)
	}
	usuariosP, err := petic.UsuarioGet(*u)
	if err != nil {
		debug.Menserr(err)
	}
	if len(usuariosP) == 0 {
		return c.JSON(http.StatusOK, code1)
	}
	claims := jwtCustomClaims{
		usuariosP[0].Password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		debug.Menserr(err)
	}
	usuariosP[0].Token = t
	return c.JSON(http.StatusOK, usuariosP)
}*/

func gcelular(c echo.Context) error {
	u := &estruct.Usuarios{}
	if err := c.Bind(u); err != nil {
		debug.Menserr(err)
	}
	usuariosP, err := petic.CelularGet(*u)
	if err != nil {
		debug.Menserr(err)
	}
	if len(usuariosP) < 1 {
		return c.JSON(http.StatusOK, code1)
	}
	return c.JSON(http.StatusOK, usuariosP)
}

func glogin(c echo.Context) error {
	u := &estruct.Usuarios{}
	if err := c.Bind(u); err != nil {
		debug.Menserr(err)
	}
	usuariosP, err := petic.UsuarioGet(*u)
	if err != nil {
		debug.Menserr(err)
	}
	if len(usuariosP) < 1 {
		return c.JSON(http.StatusOK, code1)
	}
	return c.JSON(http.StatusOK, usuariosP)
}

func gNegocios(c echo.Context) error {
	p, _ := strconv.Atoi(c.FormValue("pag"))
	negociosP, err := petic.NegociosGet(p)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, negociosP)
}

func gNegocio(c echo.Context) error {
	p, _ := strconv.Atoi(c.Param("i"))
	negocioP, err := petic.NegocioGet(p)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, negocioP)
}

func gTipServ(c echo.Context) error {
	tipservP, err := petic.TipservGet()
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, tipservP)
}

func gServicios(c echo.Context) error {
	i, _ := strconv.Atoi(c.FormValue("i"))
	p, _ := strconv.Atoi(c.FormValue("pag"))
	ServicioP, err := petic.ServiciosGet(i, p)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, ServicioP)
}

func gServicios2(c echo.Context) error {
	ServicioP, err := petic.Servicios2Get()
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, ServicioP)
}

func gServicio(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("i"))
	ServicioP, err := petic.ServicioGet(i)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, ServicioP)
}

func gMenuop(c echo.Context) error {
	menuopP, err := petic.MenuOpGet()
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, menuopP)
}

func gpyr(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("i"))
	pyrP, err := petic.PyR(i)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, pyrP)
}

func gpyr2(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("i"))
	a, _ := strconv.Atoi(c.FormValue("a"))
	pyrP, err := petic.PyR2(i, a)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, pyrP)
}

func grating(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("i"))
	ratingP, err := petic.Rating(i)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, ratingP)
}

func gcodpas(c echo.Context) error {
	codpasP, err := petic.CodigosPaises()
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, codpasP)
}

func gbusca(c echo.Context) error {
	i := c.Param("i")
	codpasP, err := petic.Servicios3Get(i)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, codpasP)
}

func gregister(c echo.Context) error {
	u := &estruct.Usuarios{}
	registerP, err := petic.RegisterGet(*u)
	if err != nil {
		debug.Menserr(err)
	}
	return c.JSON(http.StatusOK, registerP)
}
