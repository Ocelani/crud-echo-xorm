package main

import (
	"fmt"
	"time"

	"github.com/Ocelani/mercafacil/pkg"
	"github.com/Ocelani/mercafacil/pkg/simulation"
)

func main() {
	go pkg.Server()

	time.Sleep(time.Second)
	fmt.Println()

	for i := 5; i >= 0; i-- {
		time.Sleep(time.Second)
		fmt.Printf("\r ••• Simulation starting in: %v", i)
	}

	simulation.Run()

	fmt.Printf(`
			# Amigos Gophers de Curitiba,

			Confesso que não completei todos os requisitos propostos.
			Entretanto, inseri algumas firulas, como essa simulação concorrente vista logo acima.

			No README.md consta um checklist das tarefas que cumpri.

			Destaco que não finalizei a parte de persistência através do ORM entre a API e o Banco de Dados.
			Apesar disso, a configuração e conexão foi realizada com sucesso (dir "./internal").
			Até então, me parece que as rotas e funcionalidades implementadas tem executado adequadamente.

			Foi uma boa oportunidade para testar a lib "xORM", que já estava na minha lista.

			Espero que tenham gostado.
			Obrigado!

			Otávio Celani
	`)
}
