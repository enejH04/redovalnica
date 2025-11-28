# redovalnica

Package redovalnica provides functionality for managing student grades.

The package allows adding grades to students, printing all grades, and calculating final performance based on average grades.
Note that grades must be within specified minimum and maximum bounds.

Example usage:

```go
studenti := map[string]redovalnica.Student{
    "2021001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{8, 8, 9, 8}},
    "2021002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{10, 10, 9, 10, 10, 8}},
}

// Add grade to Ana Novak
redovalnica.DodajOceno(studenti, "2021001", 9, 5, 10)
redovalnica.IzpisVsehOcen(studenti)
redovalnica.IzpisiKoncniUspeh(studenti, 6, 10)
```

redovalnica would give us the following output:

```
REDOVALNICA:
2021001 - Ana Novak: [8 8 9 8 9]
2021002 - Boris Kralj: [10 10 9 10 10 8]
Ana Novak: povprečna ocena 0.0 -> Neuspešen študent
Boris Kralj: povprečna ocena 9.5 -> Odličen student!
```
