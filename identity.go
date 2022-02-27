package main

// Vorname mit Geschlecht
type Vorname struct {
	Vorname    string
	Geschlecht string
}

// Nachname ohne weitere Attribute
type Nachname struct {
	Nachname string
}

// Identitaet mit allen Attributen
type Identitaet struct {
	ID         string `json:"id"`
	Vorname    string `json:"firstName"`
	Nachname   string `json:"surname"`
	Geschlecht string `json:"sex"`
	Email      string `json:"mail"`
	Geburtstag string `json:"birthday"`
	Beruf      string `json:"job"`
	Abteilung  string `json:"department"`
	Manager    string `json:"manager"`
}
