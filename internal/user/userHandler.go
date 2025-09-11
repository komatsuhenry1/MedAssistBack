package user

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}



// func (h *UserHandler) GetAllVisits(c *gin.Context) {
// 	//VALIDACAO DE ROLE
// 	claims, exists := c.Get("claims")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
// 		return
// 	}

// 	// valido a role do user 
// 	role, ok := claims.(jwt.MapClaims)["role"].(string)
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "role inválido no token"})
// 		return
// 	}

// 	fmt.Println("role: ", role)

// 	if role != "NURSE" {
// 		utils.SendErrorResponse(c, "Rota exclusiva para enfermeiros.", http.StatusBadRequest)
// 	}
	
// 	consults, err := h.service.GetAllVisits()
// 	if err != nil{
// 		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	utils.SendSuccessResponse(c, "Consultas encontradas com sucesso.", consults)
// }