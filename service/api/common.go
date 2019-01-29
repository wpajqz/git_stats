package api

// CommonParams page and lang info
type CommonParams struct {
	Lang  string `form:"lang"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

// Check param check
func (cp *CommonParams) Check() error {
	if cp.Limit <= 0 {
		cp.Limit = 10
	}

	if cp.Page <= 0 {
		cp.Page = 1
	}

	return nil
}
