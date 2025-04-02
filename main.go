/*
-----------------------------------------------------------------------
- Programa: CLI_docker_images											-
- Autor: Walter Avendaño												-
- Fecha: 01/04/2025														-
-----------------------------------------------------------------------
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	task "github.com/WalterAvendano/CLI_docker_images/tasks"
)

/*-----------------------------------------------------------------------
- Librerias utilizadas													-
- "os" permite interactuar con el sistema operativo (abrir, leer, etc)  -
- y archivos.															-
- "io" proporciona interfaces básicas para primitivas de entrada/salida -
- (I/O). Estas interfaces son fundamentales para trabajar con flujos de -
- datos													                -
- "encoding/json" fundamental para trabajar con datos JSON			    -
-----------------------------------------------------------------------*/

func main() {
	file, err := os.OpenFile("estructura_images.json", os.O_RDWR|os.O_CREATE, 0666) // Crea si no existe el archivo con sufiientes
	//  privilegios
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// Si la variable info tiene un tamaño lee la información del archivo, la guarda en la variable bytes y de tener información lo pasa
	// al archivb JSON{}
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
		// Si la información esta vacia la asigna a una variable vacia de tareas
	} else {
		tasks = []task.Task{}
	}

	// Recibiendo los argumentos del usuario, los cuales deben ser superiores a dos, pues el primero [0] es la información del programa
	if len(os.Args) < 2 {
		printUsage()
	}

	// Creando las opciones del menu
	switch os.Args[1] {
	case "Listar":
		task.ListTask(tasks)
	case "Agregar":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\033[32;40m FROM:\033[0m")
		//fmt.Println("FROM:")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSpace(from)
		// Se actualiza "tasks" por ser un arreglo que se esta pasando
		tasks = task.AddTasks(tasks, from)
		// Invocando al metodo para guardar pasandole la variable del archivo y el arreglo
		task.SaveTasks(file, tasks)

	}
}

// Funcion creada para dar un mensaje al usuario en caso de que no pase como minimo dos argumentos
func printUsage() {
	fmt.Println("\033[32;40m El programa CLI_docker_images se usa para administrar imagenes Docker Opciones: | Listar | Agregar | Completar | Borrar |\033[0m")
}
