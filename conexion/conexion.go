package conexion

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var usuario = "weik_admin"
var pass = "silcoral9225"
var host = "tcp(systemsweb.net)"
var nombreBaseDeDatos = "weik_admin"

func ObtenerBaseDeDatos() (db *sql.DB, e error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*func ObtenerBaseDeDatos() (db *sql.DB, e error) {
	var dire, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	dire = strings.Replace(dire, `\`, `/`, -1)

	var host = "tcp(" + getCfg("db.server", dire+"/config.ini") + ")"
	var usuario = getCfg("db.user", dire+"/config.ini")
	var pass = getCfg("db.password", dire+"/config.ini")
	var nombreBaseDeDatos = getCfg("db.db", dire+"/config.ini")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getCfg(tag string, filepath string) string {
	dat, err := ioutil.ReadFile(filepath) // Leer archivo
	checkErr(err)                         // Verificar errores
	cfg := string(dat)                    // Convierte el archivo de configuración de llegada de lectura en una cadena
	var str string
	s1 := fmt.Sprintf("[^;]%s *= *.{1,}\\n", tag)
	s2 := fmt.Sprintf("%s *= *", tag)
	reg, err := regexp.Compile(s1)
	if err == nil {
		tag_str := reg.FindString(cfg) // Buscar en la cadena de configuración
		if len(tag_str) > 0 {
			r, _ := regexp.Compile(s2)
			i := r.FindStringIndex(tag_str) // Encuentra la posición inicial exacta de la cadena de configuración
			var h_str = make([]byte, len(tag_str)-i[1])
			copy(h_str, tag_str[i[1]:])
			str1 := fmt.Sprintln(string(h_str))
			str2 := strings.Replace(str1, "\n", "", -1)
			str = strings.Replace(str2, "\r", "", -1)
		}
	}
	return str
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}*/
