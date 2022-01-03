package models

type NotificationsPage struct {
	Notifications []*Notification `json:"notifications"`
}

type Notification struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	ServiceName     string `json:"service_name"`
	ServiceImageURL string `json:"service_image_url"`
}
