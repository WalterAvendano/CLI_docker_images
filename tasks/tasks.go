package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*-----------------------------------------------------------------------
- Librerias utilizadas													-
- 														-
-----------------------------------------------------------------------*/

/*
-----------------------------------------------------------------------
- Se crea la estructura del archivo json															-
-----------------------------------------------------------------------
*/
type Task struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	/*	Label      string `json:"label"`
		Run        string `json:"run"`
		Copy       string `json:"copy"`
		Workdir    string `json:"workdir"`
		Expose     string `json:"expose"`
		Env        string `json:"env"`
		Cmd        string `json:"cmd"`
		Entrypoint int    `json:"entrypoint"`
		User       string `json:"user"` */
	Completado bool `json:"completado"`
}

// Creando el metodo para listar las imagenes en el archivo JSON pasando los parametros tasks y el arreglo []Task
func ListTask(tasks []Task) {
	// En caso de que no hallan imagenes cargadas se indica. En caso contrario se recorre el archivo y se muestra la información
	if len(tasks) == 0 {
		fmt.Println("\033[32;40m No existen imagenes cargadas en archivo JSON \033[0m")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Completado {
			status = "✓"
		}
		fmt.Printf("[%s] %d %s\n", status, task.ID, task.From)
	}
}

// Creando el metodo para agregar parametros para crear la imagen y guardar los mismos en archivo JSON
func AddTasks(tasks []Task, from string) []Task {
	newTask := Task{
		ID:   len(tasks) + 1,
		From: from,
		/*		Label: 	   label,
				Run:	  run,
				Copy:	   name,
				Workdir:   workdir,
				Expose:		expose,
				Env:		env,
				Cmd:		cmd,
				Entrypoint:	1,
				User:		user,*/
		Completado: false,
	}
	return append(tasks, newTask)
}

// Creando metodo para guardar la información en archivo JSON
func SaveTasks(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// Antes de grabar el archivo se debe limpiar
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}
