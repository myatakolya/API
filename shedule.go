package API

type Teacher struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Secondname string `json:"secondname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

type Student struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Secondname string `json:"secondname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Group      int    `json:"groupe"`
	Role       string `json:"role"`
}

type Roles struct {
	Id            int    `json:"id"`
	Auth          bool   `json:"auth"`
	Naming        string `json:"naming"`
	Sheduleaccess bool   `json:"shedule"`
	Fullaccess    bool   `json:"fullaccess"`
	Comments      bool   `json:"comments"`
}

type Lesson struct {
	Subject string `json:"subject"`
	Teacher string `json:"teacher"`
	Type    string `json:"type"`
	Time    string `json:"time"`
}

type Lessons struct {
	First  Lesson `json:"first"`
	Second Lesson `json:"second"`
	Third  Lesson `json:"third"`
	Fourth Lesson `json:"fourth"`
}

type Shedule struct {
	Day     string  `json:"day"`
	Group   int     `json:"group"`
	Lessons Lessons `json:"lessons"`
}
