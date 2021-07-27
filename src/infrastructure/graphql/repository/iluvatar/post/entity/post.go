package entity

type PostsResponse struct {
	Posts []Post `json:"posts"`
}

type PostsByServiceIDResponse struct {
	Service ServiceWithPosts `json:"servicio"`
}

type ServiceWithPosts struct {
	ID string `json:"id"`
	Posts []Post `json:"posts"`
}

type Post struct {
	ID      string  `json:"id"`
	Title   string  `json:"Titulo"`
	Summary string  `json:"Resumen"`
	Body    string  `json:"Cuerpo"`
	Service Service `json:"servicio"`
}

type Service struct {
	ID string `json:"id"`
}
