// Package redovalnica provides functionality for managing student grades.
//
// The package allows adding grades to students, printing all grades, and calculating final performance based on average grades.
// Note that grades must be within specified minimum and maximum bounds.
//
// Example usage:
//
//	studenti := map[string]redovalnica.Student{
//		"2021001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{8, 8, 9, 8}},
//		"2021002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{10, 10, 9, 10, 10, 8}},
//	}
//	redovalnica.DodajOceno(studenti, "2021001", 9, 5, 10)
//	redovalnica.IzpisVsehOcen(studenti)
//	redovalnica.IzpisiKoncniUspeh(studenti, 6, 10)
//
// redovalnica would give us the following output:
//
//	REDOVALNICA:
//	2021001 - Ana Novak: [8 8 9 8 9]
//	2021002 - Boris Kralj: [10 10 9 10 10 8]
//	Ana Novak: povprečna ocena 0.0 -> Neuspešen študent
//	Boris Kralj: povprečna ocena 9.5 -> Odličen student!
package redovalnica

import "fmt"

// Student represents a student with a name, surname, and a slice of grades.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno adds a grade to a student's record if the student exists and the grade is in the valid range.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	// Mapi se posiljajo po referenci
	// ocen pod 0 ali nad 10 ne sprejemamo
	if ocena < minOcena || ocena > maxOcena {
		fmt.Printf("Neveljavna ocena (%d). Ocene morajo biti na intervalu od %d do %d!\n", ocena, minOcena, maxOcena)
		return
	}

	student, ok := studenti[vpisnaStevilka]

	// Ce studenta ni v slovarju, mu ne moremo vpisati ocene
	if !ok {
		fmt.Printf("Študenta z vpisno številko %s ni na seznamu!\n", vpisnaStevilka)
		return
	}

	// Slovar pa je reference
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
}

// IzpisVsehOcen prints all students and their grades.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	// Zgoscena tabela kljucev nam ne zagotavlja vrstnega reda
	for vpisna, student := range studenti {
		ime := student.Ime
		priimek := student.Priimek
		ocene := student.Ocene
		fmt.Printf("%s - %s %s: %v\n", vpisna, ime, priimek, ocene)
	}
}

// IzpisiKoncniUspeh calculates and prints the final performance of each student based on their average grade relative to the max possible Grade.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int, maxOcena int) {
	for vpisna, student := range studenti {
		ime := student.Ime
		priimek := student.Priimek
		povprecnaOcena := povprecje(studenti, vpisna, stOcen)

		var status string

		switch {
		case povprecnaOcena >= float64(maxOcena)*0.9:
			status = "Odličen student!"
		case povprecnaOcena >= float64(maxOcena)*0.6:
			status = "Povprečen študent"
		default:
			status = "Neuspešen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", ime, priimek, povprecnaOcena, status)
	}
}

// povprecje calculates the average grade of a student if they have enough grades; otherwise, it returns 0.0 for insufficient grades or -1.0 if the student does not exist.
func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return -1.0
	}

	if len(student.Ocene) < stOcen {
		return 0.0
	}

	sum := 0
	for _, ocena := range student.Ocene {
		sum += ocena
	}
	avg := float64(sum) / float64(len(student.Ocene))

	return avg
}
