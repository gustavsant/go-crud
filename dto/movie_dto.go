package dto

type CreateMovieDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating" binding:"required,gte=0,lte=10"`
	Cover       string  `json:"cover" binding:"required,url"`
}

type MovieResponseDTO struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Cover       string  `json:"cover"`
}
