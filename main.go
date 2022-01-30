// Test Identity Generator
// Copyright (C) 2019-2022  obsti8383 (https://github.com/obsti8383) &
//					  threepw0od (https://github.com/threepw0od)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const standardEmailDomain = "company.test"

// Quellen für Namen:
// https://offenedaten-koeln.de/dataset/vornamen
// https://github.com/HBehrens/phonet4n/blob/master/src/Tests/data/nachnamen.txt
// Quellen für Berufe: Wikipedia.de
// Quellen für Abteilungen:
// http://abimagazin.de/beruf-karriere/arbeitsmarkt/unternehmensportraits/infotext-typische-abteilungen-04868.htm

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

func main() {
	// parse command line parameters/flags
	anzahlIdentitaeten := flag.Int("anzahl", 100, "Anzahl der zu generierenden Identitäten")
	emailDomain := flag.String("domain", standardEmailDomain, "Valid Email Domain")
	jsonFlag := flag.Bool("json", false, "for output in JSON")
	managerFlag := flag.Int("manager", 0, "Anzahl der zu generierenden Manager unter den Identitäten")
	flag.Parse()
	if !isEmailDomainValid(*emailDomain) {
		fmt.Println("Error: Invalid EMail Domain")
		//flag.PrintDefaults()
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

	for i := 0; i < *anzahlIdentitaeten; i++ {
		vornm := vornamen[rand.Intn(anzahlVornamen)]
		nachnm := nachnamen[rand.Intn(anzahlNachnamen)]
		beruf := berufe[rand.Intn(anzahlBerufe)]
		abteilung := abteilungen[rand.Intn(anzahlAbteilungen)]
		email := Accents(strings.ToLower(vornm.Vorname)) + "." + Accents(strings.ToLower(nachnm.Nachname)) + "@" + *emailDomain
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

		if *managerFlag > 0 && *managerFlag < *anzahlIdentitaeten && i < *managerFlag {
			id.Manager = "isManager"
			managerIdentitaeten = append(managerIdentitaeten, id)
		}

		alleIdentitaeten = append(alleIdentitaeten, id)
	}

	if *managerFlag > 0 && *managerFlag < *anzahlIdentitaeten {
		for i, id := range alleIdentitaeten {
			if id.Manager != "isManager" {
				alleIdentitaeten[i].Manager = selectManager(managerIdentitaeten)
			} else {
				alleIdentitaeten[i].Manager = "" // set manager fields for managers to empty string
			}
		}
	}

	if *jsonFlag {
		encodedJSON, err := json.MarshalIndent(alleIdentitaeten, "", "    ")
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}
		fmt.Printf("%s\n", string(encodedJSON))
		//printIdAsJSON(id)
	} else {
		fmt.Println("id;firstName;firstname;lastname;sex;mail;birthday;job;department;manager")
		for _, id := range alleIdentitaeten {
			printIDAsCSV(id)
		}
	}
}

func selectManager(managers []Identitaet) string {
	managerID := managers[rand.Intn(len(managers))].ID

	return managerID
}

func printIDAsJSON(id Identitaet) {
	jb, _ := json.Marshal(id)
	fmt.Println(string(jb))
}

func printIDAsCSV(id Identitaet) {
	fmt.Println(id.ID + ";" + id.Vorname + ";" + id.Nachname + ";" + id.Geschlecht + ";" + id.Email + ";" + id.Geburtstag + ";" + id.Beruf + ";" + id.Abteilung + ";" + id.Manager)
}

func holeVornamen(filename string) (vornamen []Vorname, err error) {
	vornamen = make([]Vorname, 0, 0)
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return vornamen, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		line := scanner.Text()
		if !utf8.ValidString(line) {
			fmt.Println("Vornamen: no valid UTF8!")
		}
		splitLine := strings.Split(line, ";")
		name := Vorname{
			Vorname:    splitLine[0],
			Geschlecht: splitLine[1],
		}
		vornamen = append(vornamen, name)
	}

	if err := scanner.Err(); err != nil {
		return vornamen, err
	}

	return vornamen, nil
}

func holeNachnamen(filename string) (nachnamen []Nachname, err error) {
	nachnamen = make([]Nachname, 0, 0)
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nachnamen, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		line := scanner.Text()
		if !utf8.ValidString(line) {
			fmt.Println("Nachnamen: no valid UTF8!")
		}
		name := Nachname{
			Nachname: line,
		}
		nachnamen = append(nachnamen, name)
	}

	if err := scanner.Err(); err != nil {
		return nachnamen, err
	}

	return nachnamen, nil
}

func geburtstag(minAge, maxAge int, rand64 rand.Rand) time.Time {
	if minAge > maxAge {
		panic("invalid range")
	}
	now := time.Now()
	from := now.AddDate(-maxAge, 0, 0)
	to := now.AddDate(-minAge, 0, 0)

	unixTime := from.Unix() + rand64.Int63n(to.Unix()-from.Unix()+1)
	birthday := time.Unix(unixTime, 0)
	roundedBirthday := time.Date(birthday.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, birthday.Location())
	return roundedBirthday
}

func isEmailDomainValid(domain string) bool {
	emailDomainRegexp := regexp.MustCompile("^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailDomainRegexp.MatchString(domain) {
		return false
	}
	return true
}

// following code excerpts:
// Copyright (c) 2017 Florian Carrere <florian@carrere.cc>
// (The MIT License)
var (
	ErrBadFormat = errors.New("invalid format")

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func validateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}
