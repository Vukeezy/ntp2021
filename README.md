# Predlog projekta NTP - Sistem za pretragu i preporuku treninga

## Opis projekta:

U aplikaciji postoji samo neulogovani korisnik koji ima opcije da pretrazuje sve treninge po odredjenim kriterijumima kao što su: 
 - Stepen neophodne stručne spreme za izvođenje treninga
 - Da li je potrebna oprema
 - Koje grupe mišića trening aktivira, itd. 

Pored pretrage treninga korisnik takođe ima opciju da pregleda sve treninge i da nakon unosa parametara kao što su trenutna težina, godine i stepen stručne spreme dobija listu treninga koje zadovoljavaju njegove fizičke predispozicije.

Sistem takođe podržava ocenjivanje koliko je vežba korisna bila za njega kao vežbača, te kasnije može takođe vežbe pretraživati po najbolje rangiranim.

## Arhitektura sistema:

Kao backend aplikacija bi se Golang aplikacija u kojoj bi bili implementirani svi REST servisi i sama logika sistema. Za frontend, odnosno GUI bi se koristila Pharo aplikacija koja bi omogucila graficki prikaz svega navedenog. Kao baza podataka koristila bi se Postgres baza.



