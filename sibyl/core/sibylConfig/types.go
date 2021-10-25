package sibylConfig

type SibylSystemConfig struct {
	TokenSize int64  `json:"toke_size"`
	MasterId  int64  `json:"masterid"`
	MaxPanic  int64  `json:"max_panic"`
	DbUrl     string `json:"db_url"`
	DbName    string `json:"db_name"`
	UseSqlite bool   `json:"use_sqlite"`
	Port      string `json:"port"`
}
