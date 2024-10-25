package main

type fentry struct {
	name     string
	sidecard string
	hash     string
}

type recursive struct {
	name      string
	completed bool
}

var fentries []fentry

// var fedelete []fentry
var rentries []recursive
