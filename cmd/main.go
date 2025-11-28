package main

import (
	"context"
	"fmt"
	"os"

	"github.com/enejH04/redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	// Define the variables corresponding to flags

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "redovalnica omogoča upravljanje z redovalnicenico študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 6,
				Usage: "minimum number of grades needed for a positive end grade",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 5,
				Usage: "lowest possible grade",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "highest possible grade",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Printf("Najmanjše število ocen, potrebnih za pozitivno oceno: %d\nNajnižja možna ocena: %d\nNajvišja možna ocena: %d\n", stOcen, minOcena, maxOcena)

			fmt.Println("Primeri: ")

			// Studenti
			ana := redovalnica.Student{
				Ime:     "Ana",
				Priimek: "Novak",
				Ocene:   []int{8, 8, 9, 8},
			}
			boris := redovalnica.Student{
				Ime:     "Boris",
				Priimek: "Kralj",
				Ocene:   []int{10, 10, 9, 10, 10, 8},
			}
			janez := redovalnica.Student{
				Ime:     "Janez",
				Priimek: "Novak",
				Ocene:   []int{6, 7, 8, 5, 5, 6, 8, 6},
			}

			// Slovar studentov
			studenti := map[string]redovalnica.Student{
				"1000": ana,
				"1001": boris,
				"1002": janez,
			}

			fmt.Println("Dodajanje ocen:")
			// Ta student ne obstaja
			redovalnica.DodajOceno(studenti, "9999", 5, minOcena, maxOcena)
			// Te ocene ne moremo dodeliti
			redovalnica.DodajOceno(studenti, "1000", 40, minOcena, maxOcena)
			// Ana je pri OUI dobila 10
			redovalnica.DodajOceno(studenti, "1000", 10, minOcena, maxOcena)
			fmt.Println(studenti)

			fmt.Println()

			fmt.Println("Izpis redovalnice:")
			redovalnica.IzpisVsehOcen(studenti)

			fmt.Println()

			fmt.Println("Izpis končnih uspehov:")
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen, maxOcena)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println("Error:", err)
	}

}
