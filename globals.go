package main

const VERSION = "0.0.1-alpha"

type fileentry struct {
	name     string
	sidecard string
	hash     string
}

type recursive struct {
	name      string
	completed bool
}

var FileEntry []fileentry

// var fedelete []fentry
var rdirectories []recursive

type duplicate struct {
	