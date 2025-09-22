package dto

// id: 1,
// name: "Ana Silva",
// specialization: "Pediatria",
// experience: 5,
// rating: 4.8,
// price: 80,
// shift: "Manhã",
// department: "Departamento de Pediatria",
// image: "/nurse-woman-professional.jpg",
// available: true,
// location: "São Paulo - SP",
// },

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
