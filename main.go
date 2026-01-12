package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Variáveis globais
var myApp fyne.App
var myWindow fyne.Window

// Variável que virá do arquivo bundled.go
var resourceLogo fyne.Resource

func main() {
	myApp = app.New()
	
	// Tema escuro para visual moderno
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow = myApp.NewWindow("Acessa ERP")
	myWindow.Resize(fyne.NewSize(900, 600))
	myWindow.CenterOnScreen()

	// -------------------------------------------------------------------------
	// IMPORTANTE: Certifique-se de que a variável abaixo tem o MESMO nome
	// da variável gerada dentro do seu arquivo 'bundled.go'.
	// O padrão do fyne bundle é 'resourceNomeDoArquivoPng'.
	// -------------------------------------------------------------------------
	resourceLogo = resourceLogoacessaerpPng

	// Inicia na Tela de Login
	mostrarTelaLogin()

	myWindow.ShowAndRun()
}

// --- TELA 1: LOGIN (Entrada do Sistema) ---
func mostrarTelaLogin() {
	// 1. Logo
	imgLogo := canvas.NewImageFromResource(resourceLogo)
	imgLogo.FillMode = canvas.ImageFillContain
	imgLogo.SetMinSize(fyne.NewSize(180, 180))

	// 2. Campos de Entrada
	inputUsuario := widget.NewEntry()
	inputUsuario.PlaceHolder = "Usuário"
	inputUsuario.ActionItem = widget.NewIcon(theme.AccountIcon()) 

	inputSenha := widget.NewPasswordEntry()
	inputSenha.PlaceHolder = "Senha"
	// CORREÇÃO: Substituímos theme.KeyIcon (que não existe) por theme.LoginIcon
	inputSenha.ActionItem = widget.NewIcon(theme.LoginIcon()) 

	// 3. Botão de Entrar
	btnEntrar := widget.NewButton("ACESSAR SISTEMA", func() {
		// --- LÓGICA DE VALIDAÇÃO ---
		if inputUsuario.Text == "admin" && inputSenha.Text == "123" {
			mostrarTelaPrincipal()
		} else {
			dialog.ShowInformation("Erro de Acesso", "Usuário ou senha incorretos.\nTente: admin / 123", myWindow)
		}
	})
	btnEntrar.Importance = widget.HighImportance // Botão destacado

	// 4. Montagem do Layout
	formContainer := container.NewVBox(
		widget.NewLabelWithStyle("Bem-vindo de volta!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		inputUsuario,
		inputSenha,
		layout.NewSpacer(),
		btnEntrar,
	)

	cardLogin := widget.NewCard("", "", container.NewPadded(formContainer))
	
	conteudo := container.NewCenter(
		container.NewVBox(
			imgLogo,
			container.NewGridWrap(fyne.NewSize(300, 260), cardLogin),
			widget.NewLabelWithStyle("Versão 1.0.0", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
		),
	)

	myWindow.SetContent(conteudo)
}

// --- TELA 2: SISTEMA PRINCIPAL ---
func mostrarTelaPrincipal() {
	conteudoPrincipal := container.NewMax()
	conteudoPrincipal.Objects = []fyne.CanvasObject{criarDashboard()}
	conteudoPrincipal.Refresh()

	navegarPara := func(tela fyne.CanvasObject) {
		conteudoPrincipal.Objects = []fyne.CanvasObject{tela}
		conteudoPrincipal.Refresh()
	}

	// Header
	logoPequena := widget.NewIcon(resourceLogo)
	tituloTopo := widget.NewLabel("SISTEMA INTEGRADO")
	tituloTopo.TextStyle = fyne.TextStyle{Bold: true}

	btnSair := widget.NewButtonWithIcon("Sair", theme.LogoutIcon(), func() { 
		mostrarTelaLogin() 
	})
	btnSair.Importance = widget.LowImportance

	header := container.NewPadded(container.NewHBox(
		container.NewGridWrap(fyne.NewSize(40, 40), logoPequena),
		tituloTopo,
		layout.NewSpacer(),
		widget.NewLabel("Usuário: Admin"),
		btnSair,
	))

	// Menu
	criarBotao := func(t string, i fyne.Resource, f func()) *widget.Button {
		b := widget.NewButtonWithIcon(t, i, f)
		b.Alignment = widget.ButtonAlignLeading
		b.Importance = widget.MediumImportance
		return b
	}

	menu := container.NewHBox(
		container.NewPadded(container.NewVBox(
			criarBotao("Início", theme.HomeIcon(), func() { navegarPara(criarDashboard()) }),
			criarBotao("Clientes", theme.AccountIcon(), func() { navegarPara(criarTelaClientes()) }),
			criarBotao("Estoque", theme.StorageIcon(), func() { navegarPara(criarTelaEstoque()) }),
		)),
		widget.NewSeparator(),
	)

	layoutFinal := container.NewBorder(header, nil, menu, nil, container.NewPadded(conteudoPrincipal))
	myWindow.SetContent(layoutFinal)
}

// --- FUNÇÕES AUXILIARES ---

func criarDashboard() fyne.CanvasObject {
	lbl := canvas.NewText("Visão Geral", theme.ForegroundColor())
	lbl.TextSize = 24
	lbl.TextStyle = fyne.TextStyle{Bold: true}

	grid := container.NewGridWithColumns(3,
		cardMetrica("Faturamento", "R$ 4.350", theme.ConfirmIcon(), color.RGBA{0, 150, 0, 255}),
		cardMetrica("Pedidos", "12", theme.InfoIcon(), color.RGBA{0, 0, 150, 255}),
		cardMetrica("Alertas", "3", theme.WarningIcon(), color.RGBA{200, 100, 0, 255}),
	)
	return container.NewVBox(lbl, widget.NewSeparator(), grid)
}

func cardMetrica(t, v string, i fyne.Resource, c color.Color) fyne.CanvasObject {
	return widget.NewCard("", "", container.NewPadded(
		container.NewBorder(nil, nil, widget.NewIcon(i), nil,
			container.NewVBox(widget.NewLabel(t), 
			canvas.NewText(v, theme.ForegroundColor()))),
	))
}

func criarTelaClientes() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Tela de Clientes"))
}

func criarTelaEstoque() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Tela de Estoque"))
}
