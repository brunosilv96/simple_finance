package main

import (
	"fmt"
	"log"
	usecases "simple_finance/internal/finance/use_cases"
	"time"
)

func main() {
	// Create new user
	user, userError := usecases.CreateUserExecute(
		"Bruno Silva",
	)
	if userError != nil {
		log.Fatal("Error to create new user. Error: %w",userError)
	}
	
	// Create new category
	category, categoryError := usecases.CreateCategoryExecute(
		"Alimentos",
		"Gastos com alimentação",
	)
	if categoryError != nil {
		log.Fatal("Error to create new category. Error: %w",categoryError)
	}
	
	// Create new debit
	dateDebit, dateError := time.Parse("2006-01-02 15:04:05", "2026-02-01 17:53:00")
	if dateError != nil {
		log.Fatal("Error to date convert. Error: %w", dateError)
	}
	
	_, debitError := usecases.CreateNewDebit(
		*category,
		*user,
		"Compras do mês 03/26",
		"Compra unica do mês com todos os itens necessários até o mês de abril",
		dateDebit,
		756.50,
	)
	if debitError != nil {
		log.Fatal("Error to register new debit. Error: %w",debitError)
	}

	fmt.Println("Débido registrado com sucesso!")
}