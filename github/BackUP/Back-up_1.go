package main

import (
	"fmt"
	"os"

	// "github.com/fatih/color"
	"time"
)

type character struct {
	name      string
	class     string
	lvl       int
	maxlife   int
	life      int
	money     int
	inventory []string
}

func (C *character) Init(name string, class string, lvl int, maxlife int, life int, money int, inventory []string) {
	C.name = name
	C.class = class
	C.lvl = lvl
	C.maxlife = maxlife
	C.life = life
	C.inventory = inventory
	C.money = money
}

func main() {
	var C1 character
	fmt.Println("Bonjour héro, je me présente, je suis Xana et je serai votre guide durant votre aventure")
	C1.Init("Granolax", "Armagicien", 1, 200, 40, 100, []string{"potion de santé", "potion de santé", "potion de santé"})
	fmt.Println("Laissez moi vous aider, pour interagir avec moi, il vous suffit d'écrire le mot qui correspond a vos attentes :)")
	for {
		if C1.life <= 0 {
			C1.dead()
		}
		C1.menu()
	}
}

func (C1 *character) menu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("info. accéder à vos informations")
	fmt.Println("inv. accéder à votre inventaire")
	fmt.Println("stuff. gérer vos équipements")
	fmt.Println("trade. parler au marchand")
	fmt.Println("forge. permet d'accerder à la forge")
	fmt.Println("left. quitter le jeu")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "info":
		C1.displayinfo()
	case "inv":
		C1.accesinventory()
	case "trade":
		C1.trader()
	case "left":
		os.Exit(0)
	default:
		fmt.Println("Je n'ai pas compris votre requete")
	}
}

func (C *character) displayinfo() {
	fmt.Println("------------------------------------------------")
	fmt.Println("name :", C.name)
	fmt.Println("class :", C.class)
	fmt.Println("lvl :", C.lvl)
	fmt.Println("maxlife :", C.maxlife)
	fmt.Println("life :", C.life)
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
}

func (C *character) accesinventory() {
	fmt.Println("------------------------------------------------")
	fmt.Println("inventory :")
	fmt.Println("[Money :", C.money, "pièces d'or]")
	if len(C.inventory) == 0 {
		fmt.Println("votre inventaire est vide")
	}
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println("-", C.inventory[i])
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("1) utiliser une potion de santé")
	fmt.Println("utilisez 'return' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "1":
		C.takeapot()
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
}

func (C *character) takeapot() {
	i := 0
	g := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "potion de santé" {
			g++
		}
	}
	if g == 0 {
		println("vous n'avez plus de potion de santé")
	}
	if C.life == C.maxlife {
		fmt.Println("vous etes déjà full pv")
		return
	}
	if g > 0 && C.life >= C.maxlife-50 {
		fmt.Println("Vous utilisez une potion de soin, vous regagnez 50 pv")
		C.life = C.maxlife
		C.remove("potion de santé")
	} else {
		fmt.Println("Vous utilisez une potion de soin, vous regagnez 50 pv")
		C.life += 50
		C.remove("potion de santé")
	}
}

func (C *character) remove(s string) {
	var i int
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == s {
			break
		}
	}
	if i < len(C.inventory) {
		C.inventory = append(C.inventory[:i], C.inventory[i+1:]...)
		return
	}
}

func (C *character) dead() {
	if C.life <= 0 {
		println("VOUS ETES MORT")
		time.Sleep(2 * time.Second)
		println("Vous revenez à la vie avec 50% de vos PV")
		C.life = C.maxlife / 2
		C.displayinfo()
	}
}

func (C *character) poison() {
	seconds := 3
	fmt.Println("Vous buvez une fiole de poison, vous obtenez des effets indésirables pendant 3s")
	for s := seconds; s >= 1; s-- {
		if C.life-10 < 0 {
			C.life = 0
			time.Sleep(2 * time.Second)
			fmt.Println("-10 pv")
			fmt.Println("life :", C.life)
			print("\n")
			time.Sleep(1 * time.Second)
			C.dead()
			time.Sleep(2 * time.Second)
			print("\n")
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("-10 pv")
		C.life -= 10
		fmt.Println("life :", C.life)
	}
}

var popogratuite int

func (C *character) trader() {
	fmt.Println("Marchand itinérant : jette un coup d'oeil à ma boutique, j'ai plein de choses à vendre.")
	fmt.Println("------------------------------------------------")
	fmt.Println("1) potion de santé: Une potion qui rend 50 pv [3 po]")
	fmt.Println("2) fiole de poison: Une fiole de poison qui inflige 10 pv pendant 3s [6 po]")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	case "1":
		C.add("potion de santé", 3)
		popogratuite++
	case "2":
		C.add("fiole de poison", 6)
	}
}

func (C *character) add(object string, price int) {
	if popogratuite == 0 && object == "potion de santé" {
		price = 0
		println("En tant que bon vendeur, je vous offre votre premiere potion de santé")
		popogratuite++
	}
	if C.money > price {
		C.money -= price
		C.inventory = append(C.inventory, object)
		fmt.Println(object, "à été ajouté à votre inventaire")
	} else {
		fmt.Println("Vous n'avez pas assez d'argent")
	}
}
