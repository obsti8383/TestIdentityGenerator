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

// Quellen für Namen:
// https://offenedaten-koeln.de/dataset/vornamen
// https://github.com/HBehrens/phonet4n/blob/master/src/Tests/data/nachnamen.txt
// Quellen für Berufe: Wikipedia.de
// Quellen für Abteilungen:
// http://abimagazin.de/beruf-karriere/arbeitsmarkt/unternehmensportraits/infotext-typische-abteilungen-04868.htm

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
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const standardEmailDomain = "company.test"

func main() {
	// parse command line parameters/flags
	anzahlIdentitaeten := flag.Int("anzahl", 100, "Anzahl der zu generierenden Identitäten")
	emailDomain := flag.String("domain", standardEmailDomain, "Valid Email Domain")
	jsonFlag := flag.Bool("json", false, "for output in JSON")
	managerFlag := flag.Int("manager", 0, "Anzahl der zu generierenden Manager unter den Identitäten")
	flag.Parse()

	alleIdentitaeten := generate(*anzahlIdentitaeten, *emailDomain, *managerFlag)

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
