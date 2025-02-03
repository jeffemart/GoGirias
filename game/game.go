package game

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"gopkg.in/telebot.v3"
)

type Gira struct {
	Palavra string
	Dica    string
}

var (
	mu      sync.Mutex
	jogando bool
	girias  []Gira
	atual   *Gira
)

func InitGame() {
	carregarGirias()
}

func carregarGirias() {
	file, err := os.Open("girias.csv")
	if err != nil {
		log.Fatal("Erro ao abrir o CSV:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Erro ao ler o CSV:", err)
	}

	for _, line := range records {
		if len(line) < 2 {
			continue
		}
		girias = append(girias, Gira{Dica: line[1], Palavra: strings.ToLower(strings.TrimSpace(line[0]))})
	}
}

// Função para verificar se o jogo está ativo
func IsGameRunning() bool {
	mu.Lock()
	defer mu.Unlock()
	return jogando
}

// Inicia o jogo
func StartGame(c telebot.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if jogando {
		return c.Send("🎮 O jogo já está em andamento!")
	}

	jogando = true
	sortearPalavra()
	return c.Send("💡 Dica: " + atual.Dica + "\n🌩️ Qual é a gíria?")
}

// Verifica se a resposta está correta
func CheckAnswer(c telebot.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if !jogando {
		log.Println("Jogo não iniciado. Ignorando resposta.")
		return nil
	}

	resposta := strings.ToLower(strings.TrimSpace(c.Text()))
	log.Printf("Recebendo resposta: %s | Palavra correta: %s", resposta, atual.Palavra)

	if resposta == atual.Palavra {
		log.Println("Resposta correta!")
		jogando = false
		return c.Send("✅ Correto! A gíria era: *" + atual.Palavra + "* 🎉", telebot.ModeMarkdown)
	}

	log.Println("Resposta errada!")
	return c.Send("❌ Errado! Tente novamente.")
}

// Pula a palavra atual e escolhe outra
func SkipWord(c telebot.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if !jogando {
		return c.Send("⏳ Nenhuma partida ativa. Use /jogar para iniciar.")
	}

	sortearPalavra()
	return c.Send("⏭️ Palavra pulada! Nova dica: " + atual.Dica)
}

// Para o jogo
func StopGame(c telebot.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if !jogando {
		return c.Send("⚠️ Nenhuma partida ativa para parar.")
	}

	jogando = false
	return c.Send("🛑 O jogo foi encerrado! A gíria era: *" + atual.Palavra + "*", telebot.ModeMarkdown)
}

// Sorteia uma nova palavra
func sortearPalavra() {
	if len(girias) == 0 {
		log.Println("Nenhuma gíria carregada!")
		return
	}
	rand.Seed(time.Now().UnixNano())
	atual = &girias[rand.Intn(len(girias))]
}
