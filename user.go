package API

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Secondname string `json:"secondname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Group      int    `json:"group"`
	Role       string `json:"role"`
}
