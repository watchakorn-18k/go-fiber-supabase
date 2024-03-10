package entities

type CourseModel struct {
	Id       string `json:"id"`
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	CreateAt string `json:"create_at"`
	Price    int    `json:"price"`
	Details  string `json:"details"`
}
