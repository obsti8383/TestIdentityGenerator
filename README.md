![Go Vet and Lint Status](https://github.com/obsti8383/TestIdentityGenerator/.github/workflows/golang_test.yml/badge.svg)

# TestIdentityGenerator
Tool zum Generieren von Test-Identitäten (z.B. für den Bereich Identity und Access Management)

Die Identitäten werden per Zufallsgenerator mit Hilfe von Daten aus Text-Dateien (vornamen.txt, nachnamen.txt, berufe.txt, abteilungen.txt) gebildet.

Die Ausgabe kann als CSV oder JSON erfolgen.

Beispiel für die entstehenden Test-Identitäten:

```
ERA95;Eva;Raue;w;eva.raue@company.test;1951-02-09;Weinküfer;Vertrieb
ALA96;Ayt;Lamers;m;ayt.lamers@company.test;1938-10-21;Fahrradmonteur;Finanzen
ABR97;Albane;Brunken;w;albane.brunken@company.test;1930-01-17;Landwirt;Öffentlichkeitsarbeit
THI98;Tracy;Hillers;w;tracy.hillers@company.test;1955-01-01;Informatikkaufmann;Controlling
AWI99;Apostolova;Wickel;w;apostolova.wickel@company.test;1998-08-26;Fachkraft für Fruchtsafttechnik;Rechtsabteilung
```


Der Aufruf ist mit oder ohne Parameter möglich. Folgende Parameter sind vorhanden:

```
  -anzahl int
        Anzahl der zu generierenden Identitäten (default 100)
  -domain string
        Valid Email Domain (default "company.test")
  -json
        for output in JSON
  -manager
  		set number of managers      
```

Beispielaufrufe:

```
> ./TestIdentityGenerator --anzahl 4 --domain emailtestdomain.test --json
{"id":"MNE0","firstName":"Milenov","surname":"Nees","sex":"m","mail":"milenov.nees@emailtestdomain.test","birthday":"1939-07-07","job":"Schiffsmechaniker","department":"Vertrieb"}
{"id":"SHE1","firstName":"Shamanta","surname":"Helling","sex":"w","mail":"shamanta.helling@emailtestdomain.test","birthday":"1962-07-27","job":"Fachangestellter für Arbeitsmarktdienstleistungen","department":"Poststelle"}
{"id":"MBR2","firstName":"Milanea","surname":"Brutscher","sex":"w","mail":"milanea.brutscher@emailtestdomain.test","birthday":"1918-08-28","job":"Elektroniker","department":"Haustechnik"}
{"id":"KNE3","firstName":"Keslin","surname":"Nellen","sex":"w","mail":"keslin.nellen@emailtestdomain.test","birthday":"1932-11-20","job":"Weber","department":"Marketing"}
```

```
> ./TestIdentityGenerator --anzahl 3 
PKE0;Putu;Kessel;m;putu.kessel@company.test;1946-12-24;Biologisch-Technischer Assistent;Poststelle
RRA1;Redzhebov;Rauch;m;redzhebov.rauch@company.test;1929-01-11;Modeschneider;Geschäftsführung
FNA2;Francesca;Naujoks;w;francesca.naujoks@company.test;1927-12-30;Galvaniseur;Rechnungswesen
```

```
> ./TestIdentityGenerator
ADO0;Alexs;Dosch;m;alexs.dosch@company.test;1929-05-23;Elektroniker für Betriebstechnik;Controlling
YKA1;Yousef;Kapfer;m;yousef.kapfer@company.test;1924-07-24;Weinküfer;Öffentlichkeitsarbeit
AHA2;Ahmed;Haake;m;ahmed.haake@company.test;1978-04-07;Fachmann für Systemgastronomie;IT
[...]
```

```
> ./TestIdentityGenerator --anzahl 4  --manager 2
MPI0;Marina;Pille;w;marina.pille@company.test;1951-12-16;Textilreiniger;Logistik;
CMO1;Christiana;Möhring;w;christiana.moehring@company.test;1969-03-21;Maschinenzusammensetzer;Rechnungswesen;
MBO2;Mohmed;Böhner;m;mohmed.boehner@company.test;1974-04-22;Industriekeramiker Verfahrenstechnik;Poststelle;CMO1
LOE3;Liya;Oellers;w;liya.oellers@company.test;1989-07-13;Landwirtschaftlicher Laborant;Qualitätssicherung;MPI0
[...]
```
