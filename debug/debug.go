package debug

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Menserr(mensaje error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	path = strings.Replace(path, `\`, `/`, -1)
	if err != nil {
		fmt.Print(err)
	}
	path = path + "/debug.txt"
	crearArchivo(path)
	escribeArchivo(path, mensaje)
}

func crearArchivo(path string) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Created Successfully", path)
}

func escribeArchivo(path string, mensaje error) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(mensaje.Error())
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
