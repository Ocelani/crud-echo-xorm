package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

// Simulator type provides an execution
// of an use case scenario simulation.
type Simulator struct {
	Username    string
	Password    string
	BearerToken string
	Contacts    []Contact
}

// Contact definition for json operations.
type Contact struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}

type data struct {
	Contacts []Contact `json:"contacts"`
}

type token struct {
	JWT string `json:"token"`
}

//*----------*//
//* Handlers *//
//*----------*//

// ExecLogin executes app login.
func (sim *Simulator) ExecLogin() {
	username := fmt.Sprintf(`'username=%s'`, sim.Username)
	password := fmt.Sprintf(`'password=%s'`, sim.Password)

	cmd := exec.Command(
		"curl", "-X", "POST", "-d", username, "-d", password, "localhost:1323/login")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	var bearer token
	json.Unmarshal(cmdOutput.Bytes(), &bearer)
	sim.BearerToken = string(bearer.JWT)
	// fmt.Println("BEARER_TOKEN:", sim.BearerToken)

	sim.RegisterContacts()
}

// RegisterContacts registers the contacts with provided admin access.
func (sim *Simulator) RegisterContacts() {
	bearer := fmt.Sprintf("Authorization: Bearer %s", sim.BearerToken)

	for _, c := range sim.Contacts {
		b, _ := json.Marshal(c)

		cmd := exec.Command(
			"curl", "-X", "POST", "localhost:1323/users",
			"-H", "Content-Type: application/json",
			"-H", bearer,
			"-d", string(b))
		/*
			curl -X POST localhost:1323/users -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjA2OTY4MzE1LCJ1c2VybmFtZSI6Im1hY2FwYSJ9.sVet4OcyvPqU2zv3NfZKnh9cjIYT3buKEcnLow-I7fQ' -d '{"name": "Davi Lucca Rocha","cellphone": "5541979210400"}'
		*/
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput

		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		// fmt.Println(sim.Username, cmdOutput.Bytes())
	}
}

// Run the algorithm.
func Run() {
	macapa()
	varejao()
}

//*---------*//
//* Private *//
//*---------*//

func macapa() {
	m := &Simulator{
		Username: "macapa",
		Password: "PASSmacapa",
		Contacts: readJSON("./json/contacts-macapa.json"),
	}
	m.ExecLogin()

}

func varejao() {
	v := &Simulator{
		Username: "varejao",
		Password: "PASSvarejao",
		Contacts: readJSON("./json/contacts-varejao.json"),
	}
	v.ExecLogin()
}

func readJSON(file string) []Contact {
	b, _ := ioutil.ReadFile(file)

	var d data
	json.Unmarshal(b, &d)

	return d.Contacts
}
