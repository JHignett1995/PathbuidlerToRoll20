package Model

type Config struct {
	TestMode bool   `json:"testMode"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Game     string `json:"game"`
	File     string `json:"file"`
}
