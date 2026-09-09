package http

import (
	"log"
	"net/http"

	"github.com/joaovitormine/dio-financial-planning-system/adapter/http/actuator"
	"github.com/joaovitormine/dio-financial-planning-system/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
