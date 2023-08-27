# CTF-1

Moj prvi jednostavni slična CTF-u projekt.

Nasumiče ponekad ne završava baš kao treba.

Nema flagu, za razliku od uobičajenog CTF-a, ali ima kritičan bag koji omogućava kreiranje neograničenog broja usera na tuđim emailovima i zaobilazak captche. Srećno pri traženju i iskorišćavanju.


### Code structure
in `cmd` located application entripoints, files where execution is started.

in `internal/setup` located all the app initialization.

in `pkg` located simple tools (that mostly can be useful in latter projects).

`internal/domain` containes application logic, main entities and data validation.

`internal/repository` has all the storage-related code (DB, cache, disk).

in `internal/api` located everything related to an API (Handlers, routes).
