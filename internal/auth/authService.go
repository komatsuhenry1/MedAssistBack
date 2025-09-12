package auth

import (
	"fmt"
	"medassist/internal/auth/dto"
	"medassist/internal/model"
	"medassist/internal/repository"
	"medassist/utils"
	"strings"
	"time"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	UserRegister(registerRequestDTO dto.UserRegisterRequestDTO) (model.User, error)
	NurseRegister(nurseRequestDTO dto.NurseRegisterRequestDTO, files map[string][]*multipart.FileHeader) (model.Nurse, error)
	LoginUser(loginRequestDTO dto.LoginRequestDTO) (string, dto.AuthUser, error)
	SendCodeToEmail(emailAuthRequestDTO dto.EmailAuthRequestDTO) (dto.CodeResponseDTO, error)
	ValidateUserCode(inputCodeDto dto.InputCodeDto) (string, error)
}

type authService struct {
	userRepository  repository.UserRepository
	nurseRepository repository.NurseRepository
}

func NewAuthService(userRepository repository.UserRepository, nurseRepository repository.NurseRepository) AuthService {
	return &authService{userRepository: userRepository, nurseRepository: nurseRepository}
}

func (s *authService) UserRegister(registerRequestDTO dto.UserRegisterRequestDTO) (model.User, error) {
	if err := registerRequestDTO.Validate(); err != nil {
		return model.User{}, err
	}

	normalizedEmail, err := utils.EmailRegex(registerRequestDTO.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("email invalido")
	}

	// Verifica se usuário existe (sem erro se não achar)
	_, err = s.userRepository.FindUserByEmail(normalizedEmail)
	if err == nil {
		return model.User{}, fmt.Errorf("o usuario com o email '%s' ja existe", normalizedEmail)
	}

	_, err = s.userRepository.FindUserByCpf(registerRequestDTO.Cpf)
	if err == nil {
		return model.User{}, fmt.Errorf("o usuario com o CPF '%s' ja existe", registerRequestDTO.Cpf)
	}

	hashedPassword, err := utils.HashPassword(registerRequestDTO.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	// var role string
	// if registerRequestDTO.Nurse {
	// 	role = "NURSE"
	// } else {
	// 	role = "USER"
	// }

	user := model.User{
		ID:          primitive.NewObjectID(),
		Name:        registerRequestDTO.Name,
		Cpf:         registerRequestDTO.Cpf,
		Phone:       registerRequestDTO.Phone,
		Address:     registerRequestDTO.Address,
		Email:       normalizedEmail,
		Password:    hashedPassword,
		Role:        "USER",
		Hidden:      false,
		FirstAccess: true,
		TempCode:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.userRepository.CreateUser(&user); err != nil {
		return model.User{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	password := "password string test"

	if err := utils.SendEmail(registerRequestDTO.Email, password); err != nil {
		return model.User{}, fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return user, nil
}

func (s *authService) NurseRegister(nurseRequestDTO dto.NurseRegisterRequestDTO, files map[string][]*multipart.FileHeader) (model.Nurse, error) {
	if err := nurseRequestDTO.Validate(); err != nil { // valida se nao falta nenhum campo
		return model.Nurse{}, err
	}

	normalizedEmail, err := utils.EmailRegex(nurseRequestDTO.Email)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("email invalido")
	}

	// Verifica se usuário existe (sem erro se não achar)
	_, err = s.nurseRepository.FindNurseByEmail(normalizedEmail)
	if err == nil {
		return model.Nurse{}, fmt.Errorf("O(A) enfermeiro(a) com o email '%s' ja existe", normalizedEmail)
	}

	_, err = s.nurseRepository.FindNurseByCpf(nurseRequestDTO.Cpf)
	if err == nil {
		return model.Nurse{}, fmt.Errorf("O(A) enfermeiro(a) com o CPF '%s' ja existe", nurseRequestDTO.Cpf)
	}

	hashedPassword, err := utils.HashPassword(nurseRequestDTO.Password)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	// FUNCAO QUE VALIDA O RG / LICENSE_ID / ANTECEDENTES

	nurse := model.Nurse{
		ID:       primitive.NewObjectID(),
		Name:     nurseRequestDTO.Name,
		Cpf:      nurseRequestDTO.Cpf,
		Phone:    nurseRequestDTO.Phone,
		Address:  nurseRequestDTO.Address,
		Email:    normalizedEmail,
		Password: hashedPassword,
		PixKey:   nurseRequestDTO.PixKey,
		VerificationSeal: false,

		LicenseNumber:   nurseRequestDTO.LicenseNumber,
		Specialization:  nurseRequestDTO.Specialization,
		Shift:           nurseRequestDTO.Shift,
		Department:      nurseRequestDTO.Department,
		YearsExperience: nurseRequestDTO.YearsExperience,

		Role:        "NURSE",
		Hidden:      false,
		Online:      false,
		FirstAccess: true,
		TempCode:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	    // faz o upload de todos os arquivos e preenche os IDs no objeto nurse
		for fieldName, fileHeaders := range files {
			if len(fileHeaders) == 0 {
				continue // pula se não houver arquivo para este campo
			}
			fileHeader := fileHeaders[0] // pegamos apenas o primeiro arquivo por campo
	
			file, err := fileHeader.Open()
			if err != nil {
				return model.Nurse{}, fmt.Errorf("erro ao abrir o arquivo %s: %w", fileHeader.Filename, err)
			}
			defer file.Close()
			
			// cria um nome de arquivo único e descritivo
			uniqueFileName := fmt.Sprintf("%s_%s_%s", nurse.ID.Hex(), fieldName, fileHeader.Filename) // <nurse_id><license_number><image_name>
	
			// usa o método genérico do repositório
			fileID, err := s.nurseRepository.UploadFile(file, uniqueFileName) // sobe pro mongodb esse arquivo e gera o registroem fs.files
			if err != nil {
				// se um upload falhar, a operação inteira é cancelada
				return model.Nurse{}, fmt.Errorf("erro no upload do arquivo %s: %w", fileHeader.Filename, err)
			}
			
			// atribui o id ao campo correto no nosso objeto `nurse`
			switch fieldName {
			case "license_document":
				nurse.LicenseDocumentID = fileID
			case "qualifications":
				nurse.QualificationsID = fileID
			case "general_register":
				nurse.GeneralRegisterID = fileID
			case "residence_comprovant":
				nurse.ResidenceComprovantid = fileID
			}
		}
	
    if err := s.nurseRepository.CreateNurse(&nurse); err != nil {
        return model.Nurse{}, fmt.Errorf("erro ao criar o registro final do enfermeiro(a): %w", err)
    }

	password := "password string test"

	if err := utils.SendEmail(nurseRequestDTO.Email, password); err != nil {
		return model.Nurse{}, fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nurse, nil
}

func (s *authService) LoginUser(loginRequestDTO dto.LoginRequestDTO) (string, dto.AuthUser, error) {
	if err := loginRequestDTO.Validate(); err != nil {
		return "", dto.AuthUser{}, err
	}

	loginRequestDTO.Email = strings.ToLower(loginRequestDTO.Email)

	authUser, err := s.userRepository.FindUserByEmail(loginRequestDTO.Email)
	if err != nil && err.Error() == "usuário não encontrado" {
		authUser, err = s.nurseRepository.FindNurseByEmail(loginRequestDTO.Email)

		if err != nil {
			return "", dto.AuthUser{}, fmt.Errorf("Credenciais incorretas.")
		}
	} else if err != nil {
		return "", dto.AuthUser{}, err
	}

	if authUser.Hidden {
		return "", dto.AuthUser{}, fmt.Errorf("Usuário não permitido para login.")
	}
	if !utils.ComparePassword(authUser.Password, loginRequestDTO.Password) {
		return "", dto.AuthUser{}, fmt.Errorf("Credenciais incorretas.")
	}

	fmt.Println("authUser.Role: ", authUser.Role)
	fmt.Println("authUser.ID: ", authUser.ID)

	token, err := utils.GenerateToken(authUser.ID.Hex(), authUser.Role, authUser.Hidden)
	if err != nil {
		return "", dto.AuthUser{}, fmt.Errorf("erro ao gerar token: %w", err)
	}

	// Retorna o token e o usuário autenticado genérico
	return token, authUser, nil
}

func (s *authService) SendCodeToEmail(emailAuthRequestDTO dto.EmailAuthRequestDTO) (dto.CodeResponseDTO, error) {

	//busca o usuario pelo email
	user, err := s.userRepository.FindUserByEmail(emailAuthRequestDTO.Email)
	if err != nil {
		return dto.CodeResponseDTO{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	//gera o codigo
	code, err := utils.GenerateAuthCode()
	if err != nil {
		return dto.CodeResponseDTO{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	// atualiza o campo temp_code no db
	err = s.userRepository.UpdateTempCode(user.ID.Hex(), code)
	if err != nil {
		return dto.CodeResponseDTO{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	//manda para o email
	err = utils.SendAuthCode(emailAuthRequestDTO.Email, code)
	if err != nil {
		return dto.CodeResponseDTO{}, fmt.Errorf("erro ao enviar codigo de verificacao")
	}

	//retorna o code para o dto
	codeResponseDTO := dto.CodeResponseDTO{
		Code: code,
	}

	return codeResponseDTO, nil
}

func (s *authService) ValidateUserCode(inputCodeDto dto.InputCodeDto) (string, error) {

	//busca o usuario pelo email
	user, err := s.userRepository.FindUserByEmail(inputCodeDto.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao buscar user by email")
	}

	//valida o codigo inputado com o do banco
	userCode := user.TempCode

	if inputCodeDto.Code == userCode {
		token, err := utils.GenerateToken(user.ID.Hex(), user.Role, user.Hidden)
		if err != nil {
			return "", fmt.Errorf("Erro ao gerar token.")
		}
		return token, nil
	}

	return "", fmt.Errorf("Código inválido.")
}
