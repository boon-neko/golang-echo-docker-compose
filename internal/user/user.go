package user

type User struct {
	DisplayFirstName string `json:"display_first_name"`
	DisplayLastName string `json:"display_last_name"`
	ImageUrl string `json:"image_url"`
	//items []profileItem `json:"items"`
}

type profileItem struct {
	itemType string `json"item_type"`
	title string `json:"title"`
	contents interface{} `json:"contents"`
}

func (u *User) GetUser(){
	(*u).DisplayFirstName = "Cats"
	u.DisplayLastName = "Unger"

	u.ImageUrl = "https://upload.wikimedia.org/wikipedia/commons/thumb/3/33/Hannibal_Poenaru_-_Nasty_cat_%21_%28by-sa%29.jpg/540px-Hannibal_Poenaru_-_Nasty_cat_%21_%28by-sa%29.jpg"
	//item := profileItem{}

	//u.items = append(u.items, item)
}