package task

import "fmt"

/*-----------------------------------------------------------------------
- Librerias utilizadas													-
- 														-
-----------------------------------------------------------------------*/

/*-----------------------------------------------------------------------
- Se crea la estructura del archivo json															-
-----------------------------------------------------------------------*/
type Task struct {
	ID         int    `json:"id"`
	From       string `json:"from"`
	Label      string `json:"label"`
	Run        string `json:"run"`
	Copy       string `json:"copy"`
	Workdir    string `json:"workdir"`
	Expose     string `json:"expose"`
	Env        string `json:"env"`
	Cmd        string `json:"cmd"`
	Entrypoint int    `json:"entrypoint"`
	User       string `json:"user"`
	Completado bool   `json:"completado"`
}

// Creando el metodo para listar las imagenes en el archivo JSON pasando los parametros tasks y el arreglo []Task
func ListTask(tasks []Task) {
	// En caso de que no hallan imagenes cargadas se indica. En caso contrario se recorre el archivo y se muestra la información
	if len(tasks) == 0 {
		fmt.Println("\033[32;40m No existen imagenes cargadas en archivo JSON \033[0m")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d %s\n", i, task.From)
	}
}
