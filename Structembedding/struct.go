package main

type Account struct {
	Name   string
	Active bool
}

type Seller struct {
	Account
	Sales  int
	Rating int
}
