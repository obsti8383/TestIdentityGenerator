// Test Identity Generator
// Copyright (C) 2019  Florian Probst
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
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const EMAIL_DOMAIN = "company.test"

// Quellen für Namen:
// https://offenedaten-koeln.de/dataset/vornamen
// https://github.com/HBehrens/phonet4n/blob/master/src/Tests/data/nachnamen.txt

type Vorname struct {
	Vorname    string
	Geschlecht string
}

type Nachname struct {
	Nachname string
}

type Identitaet struct {
	Vorname, Nachname, Geschlecht, Email string
}

func main() {
	vornamen, _ := holeVornamen("vornamen.txt")
	nachnamen, _ := holeNachnamen("nachnamen.txt")

	anzahlVornamen := len(vornamen)
	anzahlNachnamen := len(nachnamen)
	fmt.Println("#Vornamen: " + strconv.Itoa(anzahlVornamen))
	fmt.Println("#Nachnamen: " + strconv.Itoa(anzahlNachnamen))
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		vornm := vornamen[rand.Intn(anzahlVornamen)]
		nachnm := nachnamen[rand.Intn(anzahlNachnamen)]
		id := Identitaet{
			Vorname:    vornm.Vorname,
			Nachname:   nachnm.Nachname,
			Geschlecht: vornm.Geschlecht,
			Email:      vornm.Vorname + "." + nachnm.Nachname,
		}
		fmt.Println(id.Vorname + ";" + id.Nachname + ";" + id.Geschlecht + ";" + id.Email + "@" + EMAIL_DOMAIN)
	}
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
