package main

const VERSION = "0.0.1-alpha"

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
