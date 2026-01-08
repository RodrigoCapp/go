package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Meu App Android")

	// 1. FORÇAR MODO CLARO AO INICIAR
	// Assim a aplicação começa sempre "branca", independente do sistema
	myApp.Settings().SetTheme(theme.LightTheme()) // <--- MUDANÇA AQUI

	myWindow.Resize(fyne.NewSize(400, 700))

	// --- Área de Conteúdo ---
	contentContainer := container.NewMax()

	// Tela Inicial
	textoHome := widget.NewLabel("Bem-vindo ao Sistema Mobile!")
	textoHome.Alignment = fyne.TextAlignCenter
	telaHome := container.NewVBox(
		widget.NewIcon(theme.HomeIcon()),
		textoHome,
		widget.NewButton("Ação Principal", func() { println("Clicou no Home") }),
	)

	// --- 2. TELA DE CONFIGURAÇÕES COM LÓGICA DE TEMA ---
	
	// Criamos o Checkbox com uma função de "callback"
	checkTema := widget.NewCheck("Modo Escuro", func(ativado bool) {
		if ativado {
			// Se marcou, muda para Escuro
			myApp.Settings().SetTheme(theme.DarkTheme()) // <--- MUDANÇA AQUI
		} else {
			// Se desmarcou, volta para Claro
			myApp.Settings().SetTheme(theme.LightTheme()) // <--- MUDANÇA AQUI
		}
	})
	// Garante que o checkbox comece desmarcado (pois forçamos o tema claro no início)
	checkTema.SetChecked(false)

	telaConfig := container.NewVBox(
		widget.NewLabelWithStyle("Configurações", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewCheck("Notificações", nil), // Checkbox sem ação
		checkTema,                            // <--- Nosso checkbox com ação
		widget.NewSlider(0, 100),
	)

	// Tela de Perfil
	telaPerfil := container.NewVBox(
		widget.NewIcon(theme.AccountIcon()),
		widget.NewLabel("Usuário: Dev Go"),
		widget.NewEntry(),
		widget.NewButton("Salvar", nil),
	)

	// Inicia na Home
	contentContainer.Objects = []fyne.CanvasObject{telaHome}

	// --- Menu Lateral ---
	mudarTela := func(tela fyne.CanvasObject) {
		contentContainer.Objects = []fyne.CanvasObject{tela}
		contentContainer.Refresh()
	}

	btnHome := widget.NewButtonWithIcon("Início", theme.HomeIcon(), func() { mudarTela(telaHome) })
	btnPerfil := widget.NewButtonWithIcon("Perfil", theme.AccountIcon(), func() { mudarTela(telaPerfil) })
	btnConfig := widget.NewButtonWithIcon("Ajustes", theme.SettingsIcon(), func() { mudarTela(telaConfig) })
	btnSair := widget.NewButtonWithIcon("Sair", theme.LogoutIcon(), func() { myWindow.Close() })

	menuLateral := container.NewVBox(
		widget.NewLabel(" MENU "),
		widget.NewSeparator(),
		btnHome,
		btnPerfil,
		btnConfig,
		widget.NewSeparator(),
		btnSair,
	)

	// Montagem
	split := container.NewHSplit(menuLateral, contentContainer)
	split.SetOffset(0.3)

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}