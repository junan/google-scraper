package apiv1

type User struct {
	Id   int    `jsonapi:"primary,user"`
	Name string `jsonapi:"attr,name"`
}

type Me struct {
	baseAPIController
}

func (c *Me) Get() {
	data := User{
		Id:   1,
		Name: "John Smith",
	}

	c.serveJSON(&data)
}
