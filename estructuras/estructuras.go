package estruct

type Usuarios struct {
	Iduser      int    `json:"iduser"`
	Celular     string `json:"celular"`
	Password    string `json:"password"`
	Nombres     string `json:"nombres"`
	Cedula      string `json:"cedula"`
	Fdn         string `json:"fdn"`
	Correo      string `json:"correo"`
	TipUser     string `json:"tipuser"`
	Inserta     string `json:"inserta"`
	Token       string `json:"token"`
	CodeCountry string `json:"codecountry"`
}

type Negocios struct {
	Idnegocio      int    `json:"idnegocio"`
	Nombre         string `json:"nombre"`
	Nit            string `json:"nit"`
	Direccion      string `json:"direccion"`
	ImagePrincipal string `json:"imgPrinc"`
	Image          string `json:"image"`
	Iduser         string `json:"iduser"`
	Prioridad      string `json:"prioridad"`
	Inserta        string `json:"inserta"`
	Telefonos      string `json:"telefonos"`
	Celulares      string `json:"celulares"`
	Valoriz        string `json:"valoriz"`
	Domicilio      int    `json:"domicilio"`
	Ubicacion      string `json:"ubicacion"`
}

type Tipservicio struct {
	Idtipserv int    `json:"idtipserv"`
	Img       string `json:"img"`
	Tipserv   string `json:"tipserv"`
	Estado    string `json:"estado"`
}

type Servicio struct {
	Idservicio   int    `json:"idservicio"`
	Servicio     string `json:"servicio"`
	Descripcion  string `json:"descripcion"`
	Valor        int    `json:"valor"`
	Idnegocio    string `json:"idnegocio"`
	Inserta      string `json:"inserta"`
	ImgPrincipal string `json:"imgprinc"`
	Imagenes     string `json:"imagenes"`
	Tipserv      string `json:"idtipserv"`
	Prioridad    string `json:"prioridad"`
	NegNombre    string `json:"negnombre"`
	ImgNegPrinc  string `json:"imgnegprinc"`
	Domicilio    int    `json:"domicilio"`
	Direccion    string `json:"direccion"`
	TipCobro     string `json:"tipcobro"`
	Valoriz      string `json:"valoriz"`
}

type Rating struct {
	Valorizacion string `json:"valorizacion"`
	Comentario   string `json:"comentario"`
}

type Menu_options struct {
	Option string `json:"option"`
	Direct string `json:"direct"`
	Icon   string `json:"icon"`
	Sesion int    `json:"sesion"`
}

type PyR struct {
	Pregunta   string `json:"pregunta"`
	Respuesta  string `json:"respuesta"`
	Iduser     string `json:"iduser"`
	Idservicio string `json:"idnegocio"`
	Inserta    string `json:"inserta"`
}

type Chat struct {
	IduserC int    `json:"iduserc"`
	IduserV int    `json:"iduserv"`
	Mensaje string `json:"mensaje"`
	Inserta string `json:"inserta"`
}

type CodigosPaises struct {
	Idcodpais int    `json:"idcodpais"`
	Pais      string `json:"pais"`
	Codigo    string `json:"codigo"`
}
