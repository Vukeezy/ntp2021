# Predlog projekta NTP - Sistem za pretragu i preporuku treninga

<b>Opis projekta:</b>

U aplikaciji postoji samo neulogovani korisnik koji ima opcije da pretrazuje sve treninge po odredjenim kriterijumima kao što su: stepen neophodne stručne spreme za izvođenje treninga, da li je potrebna oprema, koje grupe mišića trening aktivira, itd. Pored pretrage treninga korisnik takođe ima opciju da pregleda sve treninge i da nakon unosa parametara kao što su trenutna težina, godine i stepen stručne spreme dobija listu treninga koje zadovoljavaju njegove fizičke predispozicije.

Sistem takođe podržava ocenjivanje koliko je vežba dobro opisana i smišljena, te kasnije može takođe vežbe pretraživati po najbolje rangiranim.

<b>Arhitektura sistema:</b>

Kao backend aplikacija bi se Golang aplikacija u kojoj bi bili implementirani svi REST servisi i sama logika sistema. Za frontend, odnosno GUI bi se koristila Pharo aplikacija koja bi omogucila graficki prikaz svega navedenog. Kao baza podataka koristila bi se Postgres baza.



