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
	slotinv   int
	inventory []string
	skill     []string
	equip     stuff
}

type stuff struct {
	head string
	body string
	foot string
}

type Monstre struct {
	name  string
	pvmax int
	pv    int
	dps   int
}

func (C *character) Init(name string, class string, lvl int, maxlife int, life int, money int, slotinv int, inventory []string, skill []string, head string, body string, foot string) {
	C.name = name
	C.class = class
	C.lvl = lvl
	C.maxlife = maxlife
	C.life = life
	C.slotinv = slotinv
	C.inventory = inventory
	C.money = money
	C.skill = skill
	C.equip.head = head
	C.equip.body = body
	C.equip.foot = foot
}

func (M *Monstre) Init(name string, pvmax int, pv int, dps int) {
	M.name = name
	M.pvmax = pvmax
	M.pv = pv
	M.dps = dps
}

func main() {
	var C1 character
	var M1 Monstre
	fmt.Println("Bonjour héro, je me présente, je suis Xana et je serai votre guide durant votre aventure")
	C1.Init("Granolax", "Armagicien", 1, 100, 40, 200, 10, []string{"potion de santé", "potion de santé", "potion de santé"}, []string{"coup de poing"}, "empty", "empty", "empty")
	M1.Init("Gobelin d'entrainement", 40, 40, 5)
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
	case "stuff":
		C1.stuff()
	case "trade":
		C1.trader()
	case "forge":
		C1.forge()
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
	fmt.Println("skill :", C.skill)
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
	fmt.Println("2) apprendre le livre de sort boule de feu")
	fmt.Println("utilisez 'return' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "1":
		C.takeapot()
	case "2":
		C.spellBook("boule de feu")
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
	fmt.Println("3) livre de sort (boule de feu): un sort qui permet de projeter une boule de feu [25 po]")
	fmt.Println("4) Fourrure de Loup: une fourrure banale de gentil chien-chien[4 po]")
	fmt.Println("5) Peau de Troll: une belle peau fraichement dépecée d'un troll [7 po]")
	fmt.Println("6) Cuir de Sanglier: un cuir plutot résistant [3 po]")
	fmt.Println("7) Plume de Corbeau: une couleur magnifique pour une simple plume noire, non ? [1 po]")
	fmt.Println("8) augmentation d'inventaire: ajoute 10 places à votre inventaire [30 po]")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	case "1":
		if !C.fullinv() {
			return
		}
		C.add("potion de santé", 3)
		popogratuite++
	case "2":
		if !C.fullinv() {
			return
		}
		C.add("fiole de poison", 6)
	case "3":
		if !C.fullinv() {
			return
		}
		for i := 0; i < len(C.inventory); i++ {
			if C.inventory[i] == "livre de sort (boule de feu)" {
				println("Vous avez déjà ce sort en votre possession")
				return
			}
		}
		for i := 0; i < len(C.skill); i++ {
			if C.skill[i] == "boule de feu" {
				println("Vous avez déjà ce sort en votre possession")
				return
			}
		}
		C.add("livre de sort (boule de feu)", 25)
	case "4":
		if !C.fullinv() {
			return
		}
		C.add("fourrure de loup", 10)
	case "5":
		if !C.fullinv() {
			return
		}
		C.add("peau de troll", 7)
	case "6":
		if !C.fullinv() {
			return
		}
		C.add("cuir de sanglier", 3)
	case "7":
		if !C.fullinv() {
			return
		}
		C.add("plume de corbeau", 1)
	case "8":
		C.upgradeinventoryslot(30)
	default:
		fmt.Println("Je n'ai pas compris votre requette")
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

func (C *character) spellBook(sort string) {
	g := 0
	for i := 0; i < len(C.skill); i++ {
		if C.skill[i] == sort {
			g = 1
		}
	}
	for i := 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "livre de sort (boule de feu)" {
			g = 2
		}
	}
	if g == 2 {
		C.skill = append(C.skill, sort)
		println("Vous avez appris un nouveau sort")
		C.remove("livre de sort (boule de feu)")
	}
	if g == 0 {
		println("Vous n'avez pas ce libre dans votre inventaire")
	}
	if g == 1 {
		println("Vous avez déjà appris ce sort")
	}
}

func (C *character) forge() {
	fmt.Println("Forgeron : bienvenu dans ma forge, je peux vous craft ce que vous voulez contre 5 pièces d'or, enfin que si vous m'apportez les matériaux.")
	fmt.Println("------------------------------------------------")
	fmt.Println("1) chapeau de l'aventurier (1 plume de corbeau, 1 cuir de sanglier)")
	fmt.Println("2) tunique de l'aventurier (2 fourrure de loup, 1 peau de troll)")
	fmt.Println("3) bottes de l'aventurier (1 fourrure de loup, 1 cuir de sanglier)")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "1":
		C.craft("plume de corbeau", "cuir de sanglier", "chapeau de l'aventurier")
	case "2":
		C.craft("fourrure de loup", "peau de troll", "tunique de l'aventurier")
	case "3":
		C.craft("fourrure de loup", "cuir de sanglier", "bottes de l'aventurier")
	case "return":
		return
	default:
		fmt.Println("Xana left the game")
	}
}

func (C *character) craft(q string, r string, p string) {
	verf1 := 0
	verf2 := 0
	for i := 0; i < len(C.inventory); i++ {
		if C.inventory[i] == q {
			verf1++
		}
		if C.inventory[i] == r {
			verf2++
		}
		if p == "tunique de l'aventurier" && verf1 > 1 && verf2 > 0 {
			C.remove(q)
			C.remove(q)
			C.remove(r)
			C.add(p, 5)
			return
		} else if p != "tunique de l'aventurier" && verf1 > 0 && verf2 > 0 {
			C.remove(q)
			C.remove(r)
			C.add(p, 5)
			return
		}
	}
	fmt.Println("vous n'avez pas les matériaux")
}

func (C *character) stuff() {
	fmt.Println("------------------------------------------------")
	fmt.Println("head stuff :", C.equip.head)
	fmt.Println("body stuff :", C.equip.body)
	fmt.Println("foot stuff :", C.equip.foot)
	fmt.Println("------------------------------------------------")
	fmt.Println("1) chapeau de l'aventurier")
	fmt.Println("2) tunique de l'aventurier")
	fmt.Println("3) bottes de l'aventurier")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "1":
		C.addstuff("chapeau de l'aventurier")
	case "2":
		C.addstuff("tunique de l'aventurier")
	case "3":
		C.addstuff("bottes de l'aventurier")
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requete")
	}
}

func (C *character) addstuff(equip string) {
	i := 0
	g := 0
	p := 0
	s := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "chapeau de l'aventurier" {
			g++
		}
		if C.inventory[i] == "tunique de l'aventurier" {
			p++
		}
		if C.inventory[i] == "bottes de l'aventurier" {
			s++
		}
	}
	if equip == "chapeau de l'aventurier" && g > 0 {
		C.equip.head = "chapeau de l'aventurier"
		fmt.Println("Vous vous équipez du", equip)
		C.remove("chapeau de l'aventurier")
		C.maxlife += 10
		return
	}
	if equip == "tunique de l'aventurier" && p > 0 {
		C.equip.body = "tunique de l'aventurier"
		fmt.Println("Vous vous équipez de la", equip)
		C.remove("tunique de l'aventurier")
		C.maxlife += 25
		return
	}
	if equip == "bottes de l'aventurier" && s > 0 {
		C.equip.foot = "bottes de l'aventurier"
		fmt.Println("Vous vous équipez des", equip)
		C.remove("bottes de l'aventurier")
		C.maxlife += 15
		return
	}
	println("Vous ne posséder pas cet objet dans votre inventaire")
}

func (C *character) fullinv() bool {
	if len(C.inventory) == C.slotinv {
		fmt.Println("votre inventaire est plein")
		return false
	}
	return true
}

func (C *character) upgradeinventoryslot(price int) {
	if C.money > price && C.slotinv < 40 {
		C.money -= price
		C.slotinv += 10
		fmt.Println("votre inventaire s'agrandit de 10 places")
	} else if C.slotinv == 40 {
		fmt.Println("vous ne pouvez pas agrandir plus votre inventaire")
	} else {
		fmt.Println("Vous n'avez pas assez d'argent")
	}
}

// func ShowMainMenu() {
// 	bold(format: 'Select something :')
// 	PrintLnNtimes( 2)
// }
// color.H1Cyant (format "1: Show character information.")
// fmt.Printf(format)
