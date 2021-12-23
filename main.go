package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/urfave/cli/v2"
)

func getChars() []string {
	return []string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		"0123456789",
		"!\"#$%&'()*+,-.:;<=>?@[]^_"}
}

type Key struct {
	length   int
	isUpper  bool
	isLower  bool
	isNumber bool
	isSymbol bool
	value    string
	strength float64
}

func isUpper(noUpper bool) bool {
	return !noUpper
}

func isLower(noLower bool) bool {
	return !noLower
}

func isNumber(noNumber bool) bool {
	return !noNumber
}

func isSymbol(noSymbol bool) bool {
	return !noSymbol
}

func NewKey(length int, noUpper bool, noLower bool, noNumber bool, noSymbol bool) Key {
	k := Key{}
	k.length = length
	k.isUpper = isUpper(noUpper)
	k.isLower = isLower(noLower)
	k.isNumber = isNumber(noNumber)
	k.isSymbol = isSymbol(noSymbol)
	k.value = generate(k)
	k.strength = calculateEntropy(k.value, k.length)
	return k
}

func getRandomChars(chars string) string {
	i := rand.Intn(len(chars))
	return string(chars[i])
}

func generate(k Key) string {
	if !k.isUpper && !k.isLower && !k.isNumber && !k.isSymbol {
		fmt.Println("[\033[5;38;5;160mERROR\033[0m] Cannot create key with no chars!")
		os.Exit(1)
	}

	for len(k.value) != k.length {
		gindex := rand.Intn(4)
		if gindex == 0 && k.isUpper {
			k.value += getRandomChars(getChars()[0])
		} else if gindex == 1 && k.isLower {
			k.value += getRandomChars(getChars()[1])
		} else if gindex == 2 && k.isNumber {
			k.value += getRandomChars(getChars()[2])
		} else if gindex == 3 && k.isSymbol {
			k.value += getRandomChars(getChars()[3])
		}
	}
	return k.value
}

func calculateStrength(key string) int {
	strength := 0
	patterns := []*regexp.Regexp{regexp.MustCompile(`^.*[a-z]`), regexp.MustCompile(`^.*[A-Z]`), regexp.MustCompile(`^.*[0-9]`), regexp.MustCompile(`/^.*[!\"#$%&'()*+,-.:;<=>?@[]^_"]/`)}
	for i := 0; i < len(patterns); i++ {
		if patterns[i].MatchString(key) {
			switch i {
			case 0:
				strength += 26
			case 1:
				strength += 26
			case 2:
				strength += 10
			case 3:
				strength += 32
			default:
				strength += 0
			}
		}
	}
	return strength
}

func calculateEntropy(key string, l int) float64 {
	if l == 0 {
		return 0
	}
	strength := calculateStrength(key)
	return float64(l) * math.Log2(float64(strength))
}

type StrengthColor string

const (
	LOW    StrengthColor = "\033[38;5;124m"
	MEDIUM StrengthColor = "\033[38;5;202m"
	STRONG StrengthColor = "\033[38;5;28m"
)

func exportKeyInFile(k []Key) {
	f, err := os.OpenFile("password.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < len(k); i++ {
		if _, err := f.WriteString(k[i].value + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// Flags
	var length int
	var quantity int
	var noUpper bool
	var noLower bool
	var noNumber bool
	var noSymbol bool
	var export bool

	var color StrengthColor

	app := &cli.App{
		Name:    "gktool",
		Usage:   "generate key",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "length",
				Aliases:     []string{"l"},
				Usage:       "define length of key",
				Value:       16,
				Destination: &length,
			},
			&cli.IntFlag{
				Name:        "quantity",
				Aliases:     []string{"q"},
				Usage:       "define quantity of key",
				Value:       1,
				Destination: &quantity,
			},
			&cli.BoolFlag{
				Name:        "no-upper",
				Usage:       "define if you don't want upper chars",
				Value:       false,
				Destination: &noUpper,
			},
			&cli.BoolFlag{
				Name:        "no-lower",
				Usage:       "define if you don't want lower chars",
				Value:       false,
				Destination: &noLower,
			},
			&cli.BoolFlag{
				Name:        "no-number",
				Usage:       "define if you don't want number chars",
				Value:       false,
				Destination: &noNumber,
			},
			&cli.BoolFlag{
				Name:        "no-symbol",
				Usage:       "define if you don't want symbol chars",
				Value:       false,
				Destination: &noSymbol,
			},
			&cli.BoolFlag{
				Name:        "export",
				Aliases:     []string{"e"},
				Usage:       "export generate key(s) in file",
				Value:       false,
				Destination: &export,
			},
		},
		Action: func(c *cli.Context) error {
			// ASCII art
			fmt.Printf("\n  \033[48;0;36m██████\033[37m╗ \033[48;0;36m██\033[37m╗  \033[48;0;36m██\033[37m╗\033[38;5;125m████████\033[37m╗ \033[38;5;125m██████\033[37m╗  \033[38;5;125m██████\033[37m╗ \033[38;5;125m██\033[37m╗     \n")
			fmt.Printf(" \033[48;0;36m██\033[37m╔════╝ \033[48;0;36m██\033[37m║ \033[48;0;36m██\033[37m╔╝╚══\033[38;5;125m██\033[37m╔══╝\033[38;5;125m██\033[37m╔═══\033[38;5;125m██\033[37m╗\033[38;5;125m██\033[37m╔═══\033[38;5;125m██\033[37m╗\033[38;5;125m██\033[37m║     \n")
			fmt.Printf(" \033[48;0;36m██\033[37m║  \033[48;0;36m███\033[37m╗\033[48;0;36m█████\033[37m╔╝    \033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║\033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║\033[38;5;125m██\033[37m║     \n")
			fmt.Printf(" \033[48;0;36m██\033[37m║   \033[48;0;36m██\033[37m║\033[48;0;36m██\033[37m╔═\033[48;0;36m██\033[37m╗    \033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║\033[38;5;125m██\033[37m║   \033[38;5;125m██\033[37m║\033[38;5;125m██\033[37m║     \n")
			fmt.Printf(" ╚\033[48;0;36m██████\033[37m╔╝\033[48;0;36m██\033[37m║  \033[48;0;36m██\033[37m╗   \033[38;5;125m██\033[37m║   ╚\033[38;5;125m██████\033[37m╔╝╚\033[38;5;125m██████\033[37m╔╝\033[38;5;125m███████\033[37m╗\n")
			fmt.Printf("  ╚═════╝ ╚═╝  ╚═╝   ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝\033[0m\n\n")

			rand.Seed(time.Now().UTC().UnixNano())

			var ks []Key

			for i := 0; i < quantity; i++ {
				k := NewKey(length, noUpper, noLower, noNumber, noSymbol)
				ks = append(ks, k)
				switch {
				case int(k.strength) >= 0 && int(k.strength) < 60:
					color = LOW
				case int(k.strength) >= 60 && int(k.strength) < 90:
					color = MEDIUM
				case int(k.strength) >= 90:
					color = STRONG
				}
				fmt.Printf("\033[1mLENGTH\033[0m:%s%d\033[0m\n", color, k.length)
				fmt.Printf("\033[1mKEY%d\033[0m:%s%s\033[0m\n", i+1, color, k.value)
				fmt.Printf("\033[1mSTRENGTH\033[0m:%s%f\033[0m\n\n", color, k.strength)
			}

			if export {
				exportKeyInFile(ks)
				fmt.Printf("[\033[5;38;5;75mINFO\033[0m] Key(s) export in 'password.txt'\n")
			}

			return nil
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
