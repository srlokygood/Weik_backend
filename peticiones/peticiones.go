package peticiones

import (
	"weik/conexion"
	estruct "weik/estructuras"
)

func CelularGet(u estruct.Usuarios) ([]estruct.Usuarios, error) {
	usuariosP := []estruct.Usuarios{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT celular FROM usuarios WHERE celular=? AND codeCountry=? AND tipUser='01' LIMIT 1", u.Celular, u.CodeCountry)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Usuarios
	for filas.Next() {
		err = filas.Scan(&c.Celular)
		if err != nil {
			return nil, err
		}
		usuariosP = append(usuariosP, c)
	}
	return usuariosP, nil
}

func UsuarioGet(u estruct.Usuarios) ([]estruct.Usuarios, error) {
	usuariosP := []estruct.Usuarios{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT iduser, password, nombres, cedula, celular, Fdn, correo, tipUser, inserta FROM usuarios WHERE celular=? AND codeCountry=? AND password=md5(?) AND tipUser='01' LIMIT 1", u.Celular, u.CodeCountry, u.Password)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Usuarios
	for filas.Next() {
		err = filas.Scan(&c.Iduser, &c.Password, &c.Nombres, &c.Cedula, &c.Celular, &c.Fdn, &c.Correo, &c.TipUser, &c.Inserta)
		if err != nil {
			return nil, err
		}
		usuariosP = append(usuariosP, c)
	}
	return usuariosP, nil
}

func NegociosGet(max int) ([]estruct.Negocios, error) {
	negociosP := []estruct.Negocios{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT negocios.idnegocio,nombre,nit,direccion,image,imgPrincipal,negocios.iduser,prioridad,inserta,telefonos,celulares,round(AVG(valorizacion)) AS valoriz,ubicacion FROM negocios,rating WHERE rating.idnegocio=negocios.idnegocio AND negocios.idnegocio > ? AND negocios.valid=1 GROUP BY idnegocio ORDER BY RAND() ,prioridad DESC  LIMIT 10", max)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Negocios
	for filas.Next() {
		err = filas.Scan(&c.Idnegocio, &c.Nombre, &c.Nit, &c.Direccion, &c.Image, &c.ImagePrincipal, &c.Iduser, &c.Prioridad, &c.Inserta, &c.Telefonos, &c.Celulares, &c.Valoriz, &c.Ubicacion)
		if err != nil {
			return nil, err
		}
		negociosP = append(negociosP, c)
	}
	return negociosP, nil
}

func NegocioGet(i int) ([]estruct.Negocios, error) {
	negocioP := []estruct.Negocios{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT negocios.idnegocio,nombre,nit,direccion,image,imgPrincipal,negocios.iduser,prioridad,inserta,telefonos,celulares,round(AVG(valorizacion)) AS valoriz,ubicacion FROM negocios,rating WHERE rating.idnegocio=negocios.idnegocio AND negocios.idnegocio = ? AND negocios.valid=1 LIMIT 1", i)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Negocios
	for filas.Next() {
		err = filas.Scan(&c.Idnegocio, &c.Nombre, &c.Nit, &c.Direccion, &c.Image, &c.ImagePrincipal, &c.Iduser, &c.Prioridad, &c.Inserta, &c.Telefonos, &c.Celulares, &c.Valoriz, &c.Ubicacion)
		if err != nil {
			return nil, err
		}
		negocioP = append(negocioP, c)
	}
	return negocioP, nil
}

func TipservGet() ([]estruct.Tipservicio, error) {
	tipservP := []estruct.Tipservicio{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idtipserv,img,tipserv,estado FROM tipServicio WHERE estado = 1 ORDER BY RAND()")
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Tipservicio
	for filas.Next() {
		err = filas.Scan(&c.Idtipserv, &c.Img, &c.Tipserv, &c.Estado)
		if err != nil {
			return nil, err
		}
		tipservP = append(tipservP, c)
	}
	return tipservP, nil
}

//pedir servicios por categoria i = categoria, max = desde que id trae registros
func ServiciosGet(i, max int) ([]estruct.Servicio, error) {
	serviciosP := []estruct.Servicio{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "SELECT idservicio,servicio,valor,servicios.idnegocio,servicios.inserta,servicios.imgPrincipal,imagenes,tipserv,prioridad FROM servicios,negocios WHERE negocios.idnegocio=servicios.idnegocio AND tipserv=? AND idservicio > ? AND servicios.valid=1 AND estado=1 LIMIT 10"
	filas, err := db.Query(query, i, max)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Servicio
	for filas.Next() {
		err = filas.Scan(&c.Idservicio, &c.Servicio, &c.Valor, &c.Idnegocio, &c.Inserta, &c.ImgPrincipal, &c.Imagenes, &c.Tipserv, &c.Prioridad)
		if err != nil {
			return nil, err
		}
		serviciosP = append(serviciosP, c)
	}
	return serviciosP, nil
}

//pedir 10 servicios a azar con valor de 3 a 5
func Servicios2Get() ([]estruct.Servicio, error) {
	serviciosP := []estruct.Servicio{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idservicio,servicio,valor,servicios.idnegocio,servicios.inserta,servicios.imgPrincipal,imagenes,tipserv,negocios.nombre FROM servicios,negocios WHERE negocios.idnegocio=servicios.idnegocio AND servicios.valid=1 AND estado=1 AND  prioridad BETWEEN 3 AND 5 ORDER BY RAND() LIMIT 5")
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Servicio
	for filas.Next() {
		err = filas.Scan(&c.Idservicio, &c.Servicio, &c.Valor, &c.Idnegocio, &c.Inserta, &c.ImgPrincipal, &c.Imagenes, &c.Tipserv, &c.NegNombre)
		if err != nil {
			return nil, err
		}
		serviciosP = append(serviciosP, c)
	}
	return serviciosP, nil
}

//pedir servicios por nombre clave
func Servicios3Get(i string) ([]estruct.Servicio, error) {
	serviciosP := []estruct.Servicio{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "SELECT idservicio,servicio,valor,servicios.idnegocio,servicios.inserta,servicios.imgPrincipal,imagenes,tipServicio.tipserv,negocios.nombre FROM servicios,negocios,tipServicio WHERE negocios.idnegocio=servicios.idnegocio AND servicios.valid=1 AND servicios.estado=1 	AND servicios.idservicio>0 	AND servicios.tipserv=tipServicio.idtipserv AND negocios.idnegocio=servicios.idnegocio AND servicios.valid=1 AND servicios.estado=1 AND negocios.valid=1 AND CONCAT(negocios.nombre,servicios.servicio,tipServicio.tipserv) LIKE '%" + i + "%' ORDER BY RAND() LIMIT 10"
	filas, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Servicio
	for filas.Next() {
		err = filas.Scan(&c.Idservicio, &c.Servicio, &c.Valor, &c.Idnegocio, &c.Inserta, &c.ImgPrincipal, &c.Imagenes, &c.Tipserv, &c.NegNombre)
		if err != nil {
			return nil, err
		}
		serviciosP = append(serviciosP, c)
	}
	return serviciosP, nil
}

//detalle del servicio filtrado por el id
func ServicioGet(i int) ([]estruct.Servicio, error) {
	serviciosP := []estruct.Servicio{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT servicios.idservicio,servicio,descripcion,valor,servicios.idnegocio,servicios.inserta,servicios.imgPrincipal,imagenes,tipserv,negocios.nombre,negocios.imgPrincipal,servicios.domicilio,negocios.direccion,tipCobro.cobro,ROUND(AVG(valorizacion)) AS valoriz FROM servicios,negocios,tipCobro,rating WHERE negocios.idnegocio=servicios.idnegocio AND servicios.idservicio = ? AND servicios.valid=1 AND estado=1 AND servicios.tipCobro=tipCobro.tipCobro AND rating.idnegocio=negocios.idnegocio LIMIT 1", i)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Servicio
	for filas.Next() {
		err = filas.Scan(&c.Idservicio, &c.Servicio, &c.Descripcion, &c.Valor, &c.Idnegocio, &c.Inserta, &c.ImgPrincipal, &c.Imagenes, &c.Tipserv, &c.NegNombre, &c.ImgNegPrinc, &c.Domicilio, &c.Direccion, &c.TipCobro, &c.Valoriz)
		if err != nil {
			return nil, err
		}
		serviciosP = append(serviciosP, c)
	}
	return serviciosP, nil
}

func MenuOpGet() ([]estruct.Menu_options, error) {
	menuopP := []estruct.Menu_options{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT opcion,direct,icon,sesion FROM menu_options")
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Menu_options
	for filas.Next() {
		err = filas.Scan(&c.Option, &c.Direct, &c.Icon, &c.Sesion)
		if err != nil {
			return nil, err
		}
		menuopP = append(menuopP, c)
	}
	return menuopP, nil
}

func PyR(i int) ([]estruct.PyR, error) {
	pyrP := []estruct.PyR{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT pregunta,respuesta,iduser,idservicio,inserta FROM pyr WHERE idservicio = ? ORDER BY RAND(),inserta LIMIT 10", i)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.PyR
	for filas.Next() {
		err = filas.Scan(&c.Pregunta, &c.Respuesta, &c.Iduser, &c.Idservicio, &c.Inserta)
		if err != nil {
			return nil, err
		}
		pyrP = append(pyrP, c)
	}
	return pyrP, nil
}

func PyR2(i, a int) ([]estruct.PyR, error) {
	pyrP := []estruct.PyR{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT pregunta,respuesta,iduser,idservicio,inserta FROM pyr WHERE idservicio = ? AND iduser = ? ORDER BY inserta LIMIT 10", i, a)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.PyR
	for filas.Next() {
		err = filas.Scan(&c.Pregunta, &c.Respuesta, &c.Iduser, &c.Idservicio, &c.Inserta)
		if err != nil {
			return nil, err
		}
		pyrP = append(pyrP, c)
	}
	return pyrP, nil
}

func Rating(i int) ([]estruct.Rating, error) {
	ratingP := []estruct.Rating{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT valorizacion,comentario FROM rating WHERE idservicio = ?", i)
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.Rating
	for filas.Next() {
		err = filas.Scan(&c.Valorizacion, &c.Comentario)
		if err != nil {
			return nil, err
		}
		ratingP = append(ratingP, c)
	}
	return ratingP, nil
}

func CodigosPaises() ([]estruct.CodigosPaises, error) {
	codpasP := []estruct.CodigosPaises{}
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idcodpais,pais,codigo FROM codigosPaises")
	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var c estruct.CodigosPaises
	for filas.Next() {
		err = filas.Scan(&c.Idcodpais, &c.Pais, &c.Codigo)
		if err != nil {
			return nil, err
		}
		codpasP = append(codpasP, c)
	}
	return codpasP, nil
}

func RegisterGet(u estruct.Usuarios) ([]estruct.Usuarios, error) {
	db, err := conexion.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare("INSERT INTO usuarios (nombres, password, correo, tipUser, codeCountry, inserta) VALUES(?,?,?,'01',?,NOW())")
	if err != nil {
		return nil, err
	}
	defer sentenciaPreparada.Close()
	_, err = sentenciaPreparada.Exec(u.Nombres, u.Correo, u.TipUser, u.CodeCountry)
	if err != nil {
		return nil, err
	}
	return UsuarioGet(u)
}
