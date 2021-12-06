package courses

type Course struct {
	Name     string   `json:"name"`
	Material []string `json:"material"`
	Teacher  string   `json:"teacher"`
	Files    []string `json:"file"`
}
