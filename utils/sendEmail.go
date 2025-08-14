package utils

import (
	"fmt"
	"os"
	"gopkg.in/gomail.v2"
) 

func SendEmail(email string, password string) error {
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
			<h2>🔑 Sua Senha de Acesso</h2>
			<p>Olá,</p>
			<p>Seja bem-vindo! Sua conta foi criada com sucesso.</p>
			<p><strong>E-mail cadastrado:</strong></p>
			<div class="code-box">%s</div>

			<p><strong>Sua senha de acesso inicial:</strong></p>
			<div class="code-box">%s</div>

			<p>⚠️ Por motivos de segurança, recomendamos que você altere sua senha assim que fizer o primeiro login.</p>

			<div class="footer">
				<p>Se você não solicitou esta conta, apenas ignore este e-mail.</p>
				<p>Este é um e-mail automático. Por favor, não responda.</p>
			</div>
		</div>
	</body>
	</html>
	`, email, password)

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
