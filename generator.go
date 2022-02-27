package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func generate(anzahlIdentitaeten int, emailDomain string, managerFlag int) []Identitaet {
	if !isEmailDomainValid(emailDomain) {
		fmt.Println("Error: Invalid EMail Domain")
		//flag.PrintDefaults()
		// TODO: remove from func, return error instead
		os.Exit(1)
	}

	vornamen, _ := holeVornamen("vornamen.txt")
	nachnamen, _ := holeNachnamen("nachnamen.txt")
	berufe, _ := holeBerufe("berufe.txt")
	abteilungen, _ := holeAbteilungen("abteilungen.txt")

	anzahlVornamen := len(vornamen)
	anzahlNachnamen := len(nachnamen)
	anzahlBerufe := len(berufe)
	anzahlAbteilungen := len(abteilungen)
	//fmt.Println("#Vornamen: " + strconv.Itoa(anzahlVornamen))
	//fmt.Println("#Nachnamen: " + strconv.Itoa(anzahlNachnamen))
	//fmt.Println("#Berufe: " + strconv.Itoa(anzahlBerufe))
	//fmt.Println("#Abteilungen: " + strconv.Itoa(anzahlAbteilungen))

	alleIdentitaeten := make([]Identitaet, 0)
	managerIdentitaeten := make([]Identitaet, 0)

	rand.Seed(time.Now().UnixNano())                                                // rand seed
	rand64 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()).(rand.Source64)) // rand fuer Geburtstage

	for i := 0; i < anzahlIdentitaeten; i++ {
		vornm := vornamen[rand.Intn(anzahlVornamen)]
		nachnm := nachnamen[rand.Intn(anzahlNachnamen)]
		beruf := berufe[rand.Intn(anzahlBerufe)]
		abteilung := abteilungen[rand.Intn(anzahlAbteilungen)]
		email := Accents(strings.ToLower(vornm.Vorname)) + "." + Accents(strings.ToLower(nachnm.Nachname)) + "@" + emailDomain
		validateErr := validateFormat(email)
		if validateErr != nil {
			panic("Validation error: " + email)
		}
		id := Identitaet{
			ID:         strings.ToUpper(Accents(vornm.Vorname)[0:1]+Accents(nachnm.Nachname)[0:2]) + strconv.Itoa(i),
			Vorname:    vornm.Vorname,
			Nachname:   nachnm.Nachname,
			Geschlecht: vornm.Geschlecht,
			Email:      email,
			Geburtstag: geburtstag(16, 105, *rand64).Format("2006-01-02"),
			Beruf:      beruf,
			Abteilung:  abteilung,
		}

		if managerFlag > 0 && managerFlag < anzahlIdentitaeten && i < managerFlag {
			id.Manager = "isManager"
			managerIdentitaeten = append(managerIdentitaeten, id)
		}

		alleIdentitaeten = append(alleIdentitaeten, id)
	}

	if managerFlag > 0 && managerFlag < anzahlIdentitaeten {
		for i, id := range alleIdentitaeten {
			if id.Manager != "isManager" {
				alleIdentitaeten[i].Manager = selectManager(managerIdentitaeten)
			} else {
				alleIdentitaeten[i].Manager = "" // set manager fields for managers to empty string
			}
		}
	}

	return alleIdentitaeten
}

func isEmailDomainValid(domain string) bool {
	emailDomainRegexp := regexp.MustCompile("^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailDomainRegexp.MatchString(domain) {
		return false
	}
	return true
}
