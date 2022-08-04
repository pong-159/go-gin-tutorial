package entity

type Person struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Age uint8 `json:"age" binding:"gte=1,lte=150"`
	Email string `json:"email" validate:"required,email"`
}

type Video struct{
	Title string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	Url string `json:"url" binding:"required,url"`
	Author Person `json:"author" binding:"required"`

}
