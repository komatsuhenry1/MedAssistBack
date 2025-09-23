package dto

type AllNursesListDto struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Specialization  string  `json:"specialization"`
	YearsExperience int     `json:"years_experience"`
	Price           float32 `json:"price"`
	Shift           string  `json:"shift"`
	Department      string  `json:"department"`
	Image           string  `json:"image"`
	Available       bool    `json:"available"`
	Location        string  `json:"location"`
}
