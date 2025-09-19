package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmailNurseRegister(email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)

	m.SetHeader("Subject", "🔑 Análise de cadastro - Bem-vindo à Plataforma")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<title>Senha de Acesso</title>
		<style>
			body {
				background-color: #f9f9f9;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				color: #333333;
				padding: 0;
				margin: 0;
			}
			.container {
				max-width: 600px;
				margin: 40px auto;
				background-color: #ffffff;
				border-radius: 10px;
				box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
				padding: 30px 40px;
			}
			h2 {
				color: #1E88E5;
				text-align: center;
			}
			p {
				line-height: 1.6;
				font-size: 15px;
			}
			.code-box {
				background-color: #f1f1f1;
				border-radius: 6px;
				padding: 10px;
				font-family: monospace;
				font-size: 16px;
				color: #333333;
				margin: 15px 0;
				text-align: center;
				font-weight: bold;
			}
			.footer {
				margin-top: 30px;
				font-size: 12px;
				color: #999999;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>🔑 Sua conta está em analise para ser cadastrada no sistema como enfermeiro(a).</h2>
			<p>Olá,</p>
			<p><strong>E-mail cadastrado:</strong></p>
			<div class="code-box">%s</div>

			<p><strong>Sua conta está em analise para ser cadastrada no sistema como enfermeiro(a).</strong></p>

			<p>⚠️ Caso necessário, você pode alterar sua senha assim que fizer o primeiro login.</p>

			<div class="footer">
				<p>Se você não solicitou esta conta, apenas ignore este e-mail.</p>
				<p>Este é um e-mail automático. Por favor, não responda.</p>
			</div>
		</div>
	</body>
	</html>
	`, email)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailUserRegister(email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)

	m.SetHeader("Subject", "🔑 Cadastro de conta - Bem-vindo à Plataforma")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<title>Senha de Acesso</title>
		<style>
			body {
				background-color: #f9f9f9;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				color: #333333;
				padding: 0;
				margin: 0;
			}
			.container {
				max-width: 600px;
				margin: 40px auto;
				background-color: #ffffff;
				border-radius: 10px;
				box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
				padding: 30px 40px;
			}
			h2 {
				color: #1E88E5;
				text-align: center;
			}
			p {
				line-height: 1.6;
				font-size: 15px;
			}
			.code-box {
				background-color: #f1f1f1;
				border-radius: 6px;
				padding: 10px;
				font-family: monospace;
				font-size: 16px;
				color: #333333;
				margin: 15px 0;
				text-align: center;
				font-weight: bold;
			}
			.footer {
				margin-top: 30px;
				font-size: 12px;
				color: #999999;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>🔑 Cadastro de conta</h2>
			<p>Olá,</p>
			<p>Seja bem-vindo! Sua conta foi criada com sucesso.</p>
			<p><strong>E-mail cadastrado:</strong></p>
			<div class="code-box">%s</div>

			<p>⚠️ Caso necessário, você pode alterar sua senha assim que fizer o primeiro login.</p>

			<div class="footer">
				<p>Se você não solicitou esta conta, apenas ignore este e-mail.</p>
				<p>Este é um e-mail automático. Por favor, não responda.</p>
			</div>
		</div>
	</body>
	</html>
	`, email)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendAuthCode(email string, code int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)

	m.SetHeader("Subject", "🔑 Código de Acesso - Bem-vindo à Plataforma")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<title>Senha de Acesso</title>
		<style>
			body {
				background-color: #f9f9f9;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				color: #333333;
				padding: 0;
				margin: 0;
			}
			.container {
				max-width: 600px;
				margin: 40px auto;
				background-color: #ffffff;
				border-radius: 10px;
				box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
				padding: 30px 40px;
			}
			h2 {
				color: #1E88E5;
				text-align: center;
			}
			p {
				line-height: 1.6;
				font-size: 15px;
			}
			.code-box {
				background-color: #f1f1f1;
				border-radius: 6px;
				padding: 10px;
				font-family: monospace;
				font-size: 16px;
				color: #333333;
				margin: 15px 0;
				text-align: center;
				font-weight: bold;
			}
			.footer {
				margin-top: 30px;
				font-size: 12px;
				color: #999999;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>🔑 Seu código de acesso</h2>

			<p><strong>Code:</strong></p>
			<div class="code-box">%s</div>

			<p>⚠️ Por motivos de segurança, recomendamos que você altere sua senha no menu de segurança.</p>

			<div class="footer">
				<p>Se você não solicitou esta conta, apenas ignore este e-mail.</p>
				<p>Este é um e-mail automático. Por favor, não responda.</p>
			</div>
		</div>
	</body>
	</html>
	`, code)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailForAdmin(email string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)

	m.SetHeader("Subject", "🔑 Sua senha de acesso - Bem-vindo à Plataforma")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<title>Senha de Acesso</title>
		<style>
			body {
				background-color: #f9f9f9;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				color: #333333;
				padding: 0;
				margin: 0;
			}
			.container {
				max-width: 600px;
				margin: 40px auto;
				background-color: #ffffff;
				border-radius: 10px;
				box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
				padding: 30px 40px;
			}
			h2 {
				color: #1E88E5;
				text-align: center;
			}
			p {
				line-height: 1.6;
				font-size: 15px;
			}
			.code-box {
				background-color: #f1f1f1;
				border-radius: 6px;
				padding: 10px;
				font-family: monospace;
				font-size: 16px;
				color: #333333;
				margin: 15px 0;
				text-align: center;
				font-weight: bold;
			}
			.footer {
				margin-top: 30px;
				font-size: 12px;
				color: #999999;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>🔑 Sua Senha de Acesso (ADMINISTRADOR)</h2>
			<p>Olá,</p>
			<p>Seja bem-vindo! Sua conta de administrador foi criada com sucesso.</p>
			<p><strong>E-mail cadastrado:</strong></p>
			<div class="code-box">%s</div><br />


			<p><strong>Sua senha de acesso é a mesma que solicitou a nossa equipe na criação da conta.</strong></p>

			<div class="footer">
				<p>Se você não solicitou esta conta, apenas ignore este e-mail.</p>
				<p>Este é um e-mail automático. Por favor, não responda.</p>
			</div>
		</div>
	</body>
	</html>
	`, email)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailForgotPassword(email, id, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)

	// Link agora inclui o token no botão
	link := os.Getenv("LOCAL_FRONTEND_URL") + "?token=" + token

	m.SetHeader("Subject", "🔐 Recuperação de senha - MEDASSIST")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
	<meta charset="UTF-8">
	<title>Recuperação de Senha - CTF ARENA</title>
	<style>
	body {
		background-color: #f9f9f9;
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		color: #333333;
		padding: 0;
		margin: 0;
	}
	.container {
		max-width: 600px;
		margin: 40px auto;
		background-color: #ffffff;
		border-radius: 10px;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
		padding: 30px 40px;
	}
	h2 {
		color: #1E88E5;
		text-align: center;
	}
	p {
		line-height: 1.6;
		font-size: 15px;
	}
	.button {
		display: inline-block;
		padding: 12px 20px;
		margin: 20px 0;
		background-color: #1E88E5;
		color: #ffffff !important;
		text-decoration: none;
		border-radius: 6px;
		font-weight: 600;
		text-align: center;
	}
	.code-box {
		background-color: #f1f1f1;
		border-radius: 6px;
		padding: 10px;
		font-family: monospace;
		font-size: 14px;
		color: #333333;
		margin: 10px 0;
	}
	.footer {
		margin-top: 30px;
		font-size: 12px;
		color: #999999;
		text-align: center;
	}
	</style>
	</head>
	<body>
	<div class="container">
		<h2>🔐 Recuperação de Senha</h2>
		<p>Olá,</p>
		<p>Recebemos uma solicitação para redefinir a senha da sua conta associada ao e-mail:</p>
		<div class="code-box">%s</div>

		<p>Para criar uma nova senha, clique no botão abaixo:</p>
		<a href="%s" class="button">Redefinir Senha</a>

		<p>Se você não solicitou essa alteração, apenas ignore este e-mail. Nenhuma ação será realizada.</p>

		<div class="footer">
			<p>CTF ARENA - Este é um e-mail automático, por favor não responda.</p>
		</div>
	</div>
	</body>
	</html>
	`, email, link)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailRejectedNurse(email, reason string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "❌ Cadastro Rejeitado - MEDASSIST")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
	<meta charset="UTF-8">
	<title>Cadastro Rejeitado - MEDASSIST</title>
	<style>
	body {
		background-color: #f9f9f9;
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		color: #333333;
		padding: 0;
		margin: 0;
	}
	.container {
		max-width: 600px;
		margin: 40px auto;
		background-color: #ffffff;
		border-radius: 10px;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
		padding: 30px 40px;
	}
	h2 {
		color: #E53935;
		text-align: center;
	}
	p {
		line-height: 1.6;
		font-size: 15px;
	}
	.code-box {
		background-color: #f1f1f1;
		border-radius: 6px;
		padding: 10px;
		font-family: monospace;
		font-size: 14px;
		color: #333333;
		margin: 10px 0;
	}
	.footer {
		margin-top: 30px;
		font-size: 12px;
		color: #999999;
		text-align: center;
	}
	</style>
	</head>
	<body>
	<div class="container">
		<h2>❌ Cadastro Rejeitado</h2>
		<p>Olá,</p>
		<p>Infelizmente, sua solicitação de cadastro no sistema foi rejeitada.</p>

		<p>Motivo:</p>
		<div class="code-box">%s</div>

		<p>Se você acredita que isso foi um engano, entre em contato com o suporte para mais informações.</p>

		<div class="footer">
			<p>MEDASSIST - Este é um e-mail automático, por favor não responda.</p>
		</div>
	</div>
	</body>
	</html>
	`, reason)

	m.SetBody("text/html", html)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL_SENDER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

