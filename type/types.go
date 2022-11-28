package _type

type User struct {
	ID        string
	Username  string `json:"username" form:"username" xml:"username" binding:"required"`
	Password  string `json:"password" form:"password" xml:"password" binding:"required"`
	Usertype  int    `json:"usertype" form:"usertype" xml:"usertype" binding:"required"`
	Apartname string `json:"apartname" form:"apartname" xml:"apartname" binding:"required"'`
	Business  string `json:"business" form:"business" xml:"business" binding:"required"`
	Address   string `json:"address" form:"address" xml:"address" binding:"required"`
	Postcode  string `json:"postcode" form:"postcode" xml:"postcode" binding:"required"`
	Tel       string `json:"tel" form:"tel" xml:"tel" binding:"required"`
	Linkman   string `json:"linkman" form:"linkman" xml:"linkman" binding:"required"`
	Creid     string `json:"creid" form:"creid" xml:"creid" binding:"required"`
	Remark    string `json:"remark" form:"remark" xml:"remark" binding:"required"`
}
