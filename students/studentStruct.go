package students

type Student struct {
	Name     string `json:"name"`
	Password []byte `json:"password"`
}
