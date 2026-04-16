package ui

import (
	"lazyarduino/pkg/commands"
	"lazyarduino/pkg/utils"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// --- DEFINIÇÃO DO ITEM DA LISTA ---
// Como o utils não tem isso, definimos aqui para a List funcionar.
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// --- MENSAGENS ---
type compileFinishedMsg struct {
	err error
	out string
}

type uploadFinishedMsg struct {
	err error
	out string
}

type boardsMsg []list.Item

type model struct {
	List        list.Model
	Spinner     spinner.Model
	Serial      viewport.Model
	ProjectName string
	StatusMsg   string
	Width       int
	Height      int
	Focused     int
	IsWorking   bool
}

// --- COMANDOS (GOROUTINES) ---

func (m model) runCompile() tea.Cmd {
	return func() tea.Msg {
		fqbn, _ := board()
		out, err := commands.Compile(fqbn, m.ProjectName)
		return compileFinishedMsg{err: err, out: out}
	}
}

func (m model) runUpload() tea.Cmd {
	return func() tea.Msg {
		fqbn, port := board()
		if port == "" {
			return uploadFinishedMsg{err: nil, out: "Erro: Nenhuma placa detectada."}
		}
		out, err := commands.Upload(port, fqbn, m.ProjectName)
		return uploadFinishedMsg{err: err, out: out}
	}
}

func (m model) updateBoardsCmd() tea.Cmd {
	return func() tea.Msg {
		boards, _ := commands.ListBoards()
		var items []list.Item
		for _, b := range boards {
			nome := "Desconhecido"
			if len(b.MatchingBoards) > 0 {
				nome = b.MatchingBoards[0].Name
			}
			// Usando a struct local 'item'
			items = append(items, item{
				title: nome,
				desc:  b.Port.Address,
			})
		}
		return boardsMsg(items)
	}
}

func board() (string, string) {
	boards, err := commands.ListBoards()
	if err != nil || len(boards) == 0 {
		return "", ""
	}
	b := boards[0]
	fqbn := ""
	if len(b.MatchingBoards) > 0 {
		fqbn = b.MatchingBoards[0].FQBN
	}
	return fqbn, b.Port.Address
}

// --- CORE ---

func NewModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ab1616"))

	lista := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	lista.SetShowTitle(false)
	lista.SetShowStatusBar(false)
	lista.SetFilteringEnabled(false)
	lista.SetShowHelp(false)

	return model{
		Spinner:     s,
		List:        lista,
		Focused:     2,
		ProjectName: utils.GetProjectName(),
		StatusMsg:   "Pronto",
		IsWorking:   false, // Controle de animação
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, m.updateBoardsCmd())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd // A nossa lista de comandos para despachar

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Se estiver a trabalhar, bloqueia inputs para não bugar
		if m.IsWorking && msg.String() != "ctrl+c" && msg.String() != "q" {
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.Focused = (m.Focused + 1) % 4
		case "1", "2", "3", "0":
			m.Focused = int(msg.String()[0] - '0')
		case "r":
			m.IsWorking = true
			m.StatusMsg = "Buscando placas..."
			// Adicionamos dois comandos ao lote (batch)
			cmds = append(cmds, m.Spinner.Tick, m.updateBoardsCmd())
		case "c":
			m.IsWorking = true
			m.StatusMsg = "Compilando..."
			cmds = append(cmds, m.Spinner.Tick, m.runCompile())
		case "u":
			m.IsWorking = true
			m.StatusMsg = "Fazendo Upload..."
			cmds = append(cmds, m.Spinner.Tick, m.runUpload())
		case "esc":
			m.IsWorking = false
			m.StatusMsg = "Pronto"
		}

	case boardsMsg:
		m.IsWorking = false
		m.StatusMsg = "Placas Atualizadas"
		// Capturamos o comando de atualizar os itens da lista
		cmd = m.List.SetItems(msg)
		cmds = append(cmds, cmd)

	case compileFinishedMsg:
		m.IsWorking = false
		if msg.err != nil {
			m.StatusMsg = "Erro na Compilação"
			m.Serial.SetContent(msg.out)
		} else {
			m.StatusMsg = "Compilado com Sucesso"
			m.Serial.SetContent(msg.out)
		}

	case uploadFinishedMsg:
		m.IsWorking = false
		m.StatusMsg = "Processo Terminado"
		m.Serial.SetContent(msg.out)

	case spinner.TickMsg:
		// O Spinner só anima se IsWorking for true
		if m.IsWorking {
			m.Spinner, cmd = m.Spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	// Só atualiza a lógica da lista se o foco estiver nela
	if m.Focused == 2 && !m.IsWorking {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	// Retorna o modelo atualizado e o lote de comandos acumulados
	return m, tea.Batch(cmds...)
}
