package sibylConfig

type SibylSystemConfig struct {
	TokenSize  int64  `json:"toke_size"`
	MasterId   int64  `json:"masterid"`
	DbUrl      string `json:"db_url"`
	DbName     string `json:"db_name"`
	UseSqllite bool   `json:"use_sqlite"`
	Port       string `json:"port"`
}
