package task

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
