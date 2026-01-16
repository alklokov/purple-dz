package product

type ProductToCreate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductToUpdate struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price"`
}
