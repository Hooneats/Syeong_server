package server

type Server struct {
	Mode        string   `json:"mode"`
	Port        string   `json:"port"`
	Version     string   `json:"version"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      []string `json:"author"`
	Spec        []string `json:"spec"`
	Blog        string   `json:"blog"`
	Github      string   `json:"github"`
}
