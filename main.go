package main

import (
	"log"
	"os"
	"time"

	"github.com/jeffemart/GoGirias/game"

	"gopkg.in/telebot.v3"
)

func main() {
	// Ler o token do Telegram a partir da variável de ambiente
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN não está definido.")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	game.InitGame()

	bot.Handle("/jogar", game.StartGame)
	bot.Handle("/pular", game.SkipWord)
	bot.Handle("/parar", game.StopGame)

	// Adicionar log de todos os tipos de chat
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		log.Printf("Recebendo mensagem em chat tipo: %s, de usuário: %s", c.Chat().Type, c.Sender().Username)
		log.Printf("Mensagem recebida: %s", c.Text())

		// Verificar se o chat é de um grupo
		if c.Chat().Type == telebot.ChatGroup || c.Chat().Type == telebot.ChatSuperGroup {
			log.Printf("Mensagem recebida no grupo: %s", c.Text())
		} else {
			log.Printf("Mensagem recebida em chat privado.")
		}

		// Verifica se o jogo está em andamento
		if !game.IsGameRunning() {
			return c.Send("⏳ Nenhuma partida ativa. Use /jogar para começar!")
		}

		// Log de recepção do palpite
		log.Printf("Palpite recebido: %s", c.Text())

		// Verifica a resposta do palpite
		return game.CheckAnswer(c)
	})

	log.Println("🤖 Bot está rodando...")
	bot.Start()
}
