package domain

type Category string

const (
	CategoryClothing	Category = "CLOTHING"
	CategoryAcessories	Category = "ACCESSORIES"
	CategoryPersonality	Category = "PERSONALITY"
)

var validCategories = []Category{
	CategoryClothing,
	CategoryAcessories,
	CategoryPersonality,
}

func IsValidCategory(category Category) bool {
	for _, v := range validCategories {
		if v == category {
			return true
		}
	}
	return false
}