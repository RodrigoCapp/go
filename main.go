package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Variáveis globais para facilitar o acesso entre telas
var myApp fyne.App
var myWindow fyne.Window
var resourceLogo fyne.Resource // Vamos carregar a imagem na memória aqui

func main() {
	myApp = app.New()
	myApp.Settings().SetTheme(theme.LightTheme())

	myWindow = myApp.NewWindow("Acessa ERP")
	myWindow.Resize(fyne.NewSize(400, 700))

	// 1. CARREGAR A IMAGEM COMO UM "RECURSO"
	var err error
	resourceLogo, err = fyne.LoadResourceFromPath("logo_acessa_erp.png")
	if err != nil {
		resourceLogo = theme.WarningIcon()
		println("Erro: Imagem logo_acessa_erp.png não encontrada")
	}

	// Inicia mostrando a Tela de Splash (Abertura)
	mostrarTelaSplash()

	myWindow.ShowAndRun()
}

// --- TELA 1: SPLASH SCREEN (Abertura) ---
func mostrarTelaSplash() {
	// 1. Criamos a imagem separada do botão para termos controle total do tamanho
	imgLogo := canvas.NewImageFromResource(resourceLogo)
	imgLogo.FillMode = canvas.ImageFillContain
	
	// FORÇAMOS UM TAMANHO MÍNIMO (Aqui você controla o tamanho do logo)
	imgLogo.SetMinSize(fyne.NewSize(250, 250)) 

	// 2. Botão para entrar
	btnEntrar := widget.NewButton("Acessar Sistema", func() {
		mostrarTelaPrincipal()
	})
	// Define o botão como "Primário" (Azul/Destacado)
	btnEntrar.Importance = widget.HighImportance 

	// 3. Montamos o layout: Imagem em cima, botão embaixo
	boxVertical := container.NewVBox(
		imgLogo,
		widget.NewSeparator(), // Um espacinho entre a imagem e o botão
		btnEntrar,
	)

	// Centralizamos tudo na tela
	conteudo := container.NewCenter(boxVertical)

	myWindow.SetContent(conteudo)
}

// --- TELA 2: SISTEMA PRINCIPAL (Menu + Conteúdo) ---
func mostrarTelaPrincipal() {
	// 1. Área de Conteúdo (Onde as telas vão mudar)
	conteudoPrincipal := container.NewMax()

	// Define a tela inicial (Dashboard)
	conteudoPrincipal.Objects = []fyne.CanvasObject{criarDashboard()}

	// Função auxiliar para trocar o conteúdo do centro
	navegarPara := func(tela fyne.CanvasObject) {
		conteudoPrincipal.Objects = []fyne.CanvasObject{tela}
		conteudoPrincipal.Refresh()
	}

	// 2. Barra Superior (Logo pequena + Título)
	// Usamos NewGridWrap para forçar a logo a ficar pequena (ex: 50x50)
	logoPequena := widget.NewIcon(resourceLogo)
	containerLogo := container.NewGridWrap(fyne.NewSize(50, 50), logoPequena)

	tituloTopo := widget.NewLabel("Sistema Integrado")
	tituloTopo.TextStyle = fyne.TextStyle{Bold: true}

	// Botão para voltar à tela de splash (Logout)
	btnSair := widget.NewButtonWithIcon("", theme.LogoutIcon(), func() { mostrarTelaSplash() })

	barraSuperior := container.NewHBox(
		containerLogo,
		tituloTopo,
		layout.NewSpacer(), // Empurra o botão sair para a direita
		btnSair,
	)

	// 3. Menu Lateral
	btnDash := widget.NewButtonWithIcon("Início", theme.HomeIcon(), func() {
		navegarPara(criarDashboard())
	})
	btnClientes := widget.NewButtonWithIcon("Clientes", theme.AccountIcon(), func() {
		navegarPara(criarTelaClientes())
	})
	btnEstoque := widget.NewButtonWithIcon("Estoque", theme.StorageIcon(), func() {
		navegarPara(criarTelaEstoque())
	})

	menuLateral := container.NewVBox(
		widget.NewLabel("MENU"),
		widget.NewSeparator(),
		btnDash,
		btnClientes,
		btnEstoque,
	)

	// 4. Montagem do Layout (BorderLayout)
	layoutFinal := container.NewBorder(barraSuperior, nil, menuLateral, nil, conteudoPrincipal)

	myWindow.SetContent(layoutFinal)
}

// --- FUNÇÕES AUXILIARES QUE CRIAM O CONTEÚDO DAS TELAS ---

func criarDashboard() fyne.CanvasObject {
	return container.NewCenter(
		widget.NewLabel("Bem-vindo ao Dashboard! Selecione uma opção no menu."),
	)
}

func criarTelaClientes() fyne.CanvasObject {
	lista := widget.NewList(
		func() int { return 5 }, // 5 clientes de exemplo
		func() fyne.CanvasObject { return widget.NewLabel("Template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("Cliente Exemplo " + string(rune('A'+i)))
		},
	)
	return container.NewBorder(widget.NewLabelWithStyle("Lista de Clientes", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), nil, nil, nil, lista)
}

func criarTelaEstoque() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle("Controle de Estoque", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewProgressBarInfinite(), // Apenas visual
		widget.NewButton("Adicionar Produto", func() { println("Adicionando...") }),
	)
}
