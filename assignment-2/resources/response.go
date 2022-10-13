package resources

type Response struct {
	Message string `json:"message"`
	Data    Order  `json:"data"`
}
