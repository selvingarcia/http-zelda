package main

import "github.com/selvingarcia/http-zelda/internal/models"

type Status struct{
	Code int
	Phrase string
}

type StatusEndpoint struct{
	Code int
	Phrase string
	Title string
	Content string
	Meaning string
}

type StatusData struct{
	Status models.Status
}