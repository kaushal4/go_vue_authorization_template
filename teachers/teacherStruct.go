package teachers

type Teacher struct {
	Name     string `json:"name"`
	Password []byte `json:"password"`
}
