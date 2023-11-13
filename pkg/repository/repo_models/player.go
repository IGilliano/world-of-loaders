package repo_models

type Player struct {
	ID       int
	Login    string `json:"login"`
	Password string `json:"password"`
	Class    string `json:"class"`
}
