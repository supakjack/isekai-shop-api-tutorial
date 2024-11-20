package model

type (
	Item struct {
		ID          uint64
		Name        string
		Description string
		Picture     string
		Price       uint
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
	}
)
