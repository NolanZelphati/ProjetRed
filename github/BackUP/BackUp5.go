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

//   Variable globale   //
var T int = 0
var C *character = &character{}
var C1 character = character{}
var M *Monstre = &Monstre{}
var M1 Monstre = Monstre{}

func main() {
	C1.Init(C.name, C.class, 1, C.maxlife, 1, 100, 10, []string{"potion de santé", "potion de santé", "potion de santé"}, []string{"coup de poing"}, "empty", "empty", "empty")
	M1.Init("Gobelin d'entrainement", 40, 40, 5)
	C1.charcreation()
	for {
		if C1.life <= 0 {
			C1.dead()
			return
		} else {
			C1.menu()
		}
	}
}

func (C1 *character) menu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("info. accéder à vos informations")
	fmt.Println("inv. accéder à votre inventaire")
	fmt.Println("stuff. gérer vos équipements")
	fmt.Println("trade. parler au marchand")
	fmt.Println("forge. permet d'accerder à la forge")
	fmt.Println("combat. permet de s'excercer au combat")
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
	case "combat":
		C1.TrainingFight(M)
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

var popogratuite int

func (C *character) takeapot() {
	if C.life == C.maxlife {
		fmt.Println("vous etes déjà full pv")
		return
	}
	if C.life >= C.maxlife-50 {
		fmt.Println("Vous utilisez une potion de soin, vous regagnez 50 pv")
		C.life = C.maxlife
	} else {
		fmt.Println("Vous utilisez une potion de soin, vous regagnez 50 pv")
		C.life += 50
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

func (C *character) dead() {
	if C.life <= 0 {
		println("VOUS ETES MORT")
		time.Sleep(2 * time.Second)
		println("Vous revenez à la vie avec 50% de vos PV")
		C.life = C.maxlife / 2
	}
}

func (C *character) accesinventory() {
	if len(C.inventory) == 0 {
		fmt.Println("votre inventaire est vide")
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("see permet de voir les objets dans votre inventaire")
	fmt.Println("use permet d'utiliser un objet")
	fmt.Println("delete permet de supprimer un objet")
	fmt.Println("utilisez 'return' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)
	switch interaction {
	case "see":
		C.see()
	case "use":
		C.use()
	case "delete":
		C.delete()
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
}

func (C *character) see() {
	nb := 1
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez 'return' pour revenir en arrière")
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

func (C *character) delete() {
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case k:
		C.supp(k - 1)
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
}

func (C *character) supp(rang int) {
	var g string
	if rang >= len(C.inventory) {
		return
	}
	if C.inventory[rang] == "potion de santé" {
		g = "potion de santé"
	} else if C.inventory[rang] == "fiole de poison" {
		g = "fiole de poison"
	} else if C.inventory[rang] == "livre de sort (boule de feu)" {
		g = "livre de sort (boule de feu)"
	} else if C.inventory[rang] == "chapeau de l'aventurier" {
		g = "chapeau de l'aventurier"
	} else if C.inventory[rang] == "tunique de l'aventurier" {
		g = "tunique de l'aventurier"
	} else if C.inventory[rang] == "bottes de l'aventurier" {
		g = "bottes de l'aventurier"
	}
	C.remove(g)
	fmt.Println(g, "a été supprimé de l'inventaire")
}

func (C *character) use() {
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case k:
		C.verf(k - 1)
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
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
	case "1":
		C.add("potion de santé", 3)
		popogratuite++
	case "2":
		C.add("fiole de poison", 6)
	case "3":
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
		C.add("fourrure de loup", 10)
	case "5":
		C.add("peau de troll", 7)
	case "6":
		C.add("cuir de sanglier", 3)
	case "7":
		C.add("plume de corbeau", 1)
	case "8":
		C.upgradeinventoryslot(30)
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requette")
	}
}

func (C *character) spellBook(sort string) {
	g := 0
	rang := 0
	for i := 0; i < len(C.skill); i++ {
	}
	for i := 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "livre de sort (boule de feu)" {
			rang = i
			g = 1
		}
	}
	if g == 1 {
		C.skill = append(C.skill, sort)
		println("Vous avez appris un nouveau sort")
		C.removestuff(rang)
	}
	if g == 0 {
		println("Vous n'avez pas ce libre dans votre inventaire")
	}
}

func (C *character) forge() {
	fmt.Println("Forgeron : bienvenu dans ma forge, je peux vous craft ce que vous voulez contre 5 pièces d'or, enfin que si vous m'apportez les mattériaux.")
	fmt.Println("------------------------------------------------")
	fmt.Println("1) chapeau de l'aventurier (1 plume de corbeau, 1 cuir de sanglier)")
	fmt.Println("2) tunique de l’aventurier (2 fourrure de loup, 1 peau de troll)")
	fmt.Println("3) bottes de l’aventurier (1 fourrure de loup, 1 cuir de sanglier)")
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
		fmt.Println("Je n'ai pas compris votre requette")
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
	fmt.Println("vous n'avez pas les matériaux")
}

func (C *character) add(object string, price int) {
	if popogratuite == 0 && object == "potion de santé" {
		price = 0
		println("En tant que bon vendeur, je vous offre votre premiere potion de santé")
		popogratuite++
	}
	if len(C.inventory) == C.slotinv {
		fmt.Println("votre inventaire est plein")
		return
	}
	if C.money > price {
		C.money -= price
		C.inventory = append(C.inventory, object)
		// sortie = append(sortie, object)
		fmt.Println(object, "à été ajouté à votre inventaire")
	} else {
		fmt.Println("Vous n'avez pas assez d'argent")
	}
}

func (C *character) addstuff(equip string) {
	if equip == "chapeau de l'aventurier" && C.equip.head != "chapeau de l'aventurier" {
		C.equip.head = "chapeau de l'aventurier"
		fmt.Println("Vous vous équipez du", equip)
		C.remove("chapeau de l'aventurier")
		C.maxlife += 10
		return
	}
	if equip == "tunique de l'aventurier" && C.equip.body != "tunique de l'aventurier" {
		C.equip.body = "tunique de l'aventurier"
		fmt.Println("Vous vous équipez de la", equip)
		C.remove("tunique de l'aventurier")
		C.maxlife += 25
		return
	}
	if equip == "bottes de l'aventurier" && C.equip.foot != "bottes de l'aventurier" {
		C.equip.foot = "bottes de l'aventurier"
		fmt.Println("Vous vous équipez des", equip)
		C.remove("bottes de l'aventurier")
		C.maxlife += 15
		return
	}
	println("Vous etes déja équipé de cet objet")
}

func (C *character) remove(s string) {
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == s {
			break
		}
	}
	if i < len(C.inventory)-1 {
		C.inventory = append(C.inventory[:i], C.inventory[i+1:]...)
	} else if i == len(C.inventory)-1 {
		C.inventory = C.inventory[:i]
	}
}

func (C *character) removestuff(number int) {
	if number < len(C.inventory)-1 {
		C.inventory = append(C.inventory[:number], C.inventory[number+1:]...)
	} else if number == len(C.inventory)-1 {
		C.inventory = C.inventory[:number]
	}
}

func (C *character) verf(rang int) {
	if rang >= len(C.inventory) {
		return
	}
	if C.inventory[rang] == "potion de santé" {
		C.takeapot()
		C.removestuff(rang)
	} else if C.inventory[rang] == "fiole de poison" {
		C.poison()
		C.removestuff(rang)
	} else if C.inventory[rang] == "livre de sort (boule de feu)" {
		C.spellBook("boule de feu")
	} else if C.inventory[rang] == "chapeau de l'aventurier" {
		C.addstuff("chapeau de l'aventurier")
	} else if C.inventory[rang] == "tunique de l'aventurier" {
		C.addstuff("tunique de l'aventurier")
	} else if C.inventory[rang] == "bottes de l'aventurier" {
		C.addstuff("bottes de l'aventurier")
	}
}

func (C *character) TrainingFight(M *Monstre) {
	T := 0
	for {
		if T == 0 {
			M1.pv = M1.pvmax
		}
		if C1.life > 0 && M1.pv > 0 {
			T += 1
			fmt.Println("------------------------------------------------")
			fmt.Println("Vous etes au tour n°", T)
			C.charTurn()
			time.Sleep(2 * time.Second)
			M.gobelinPattern(C, T)
		} else {
			if M1.pv == 0 {
				time.Sleep(2 * time.Second)
				fmt.Println("Vous avez vaincu le Gobelin d'entrainement.")
				time.Sleep(2 * time.Second)
				fmt.Println("Vous vous dirigez victorieusement vers la sortie mais votre soif de combat n'a pas été rassasiée.")
				time.Sleep(2 * time.Second)
				C.menu()
				break
			}
			if C1.life == 0 {
				fmt.Println("------------------------------------------------")
				time.Sleep(2 * time.Second)
				fmt.Println("Le Gobelin d'entrainement vous a mis une raclée.")
				time.Sleep(2 * time.Second)
				C.dead()
				time.Sleep(2 * time.Second)
				fmt.Println("Vous reprenez connaisance en dehors du donjon")
				time.Sleep(2 * time.Second)
				C.menu()
				break
			}
		}
	}
}

func (C *character) charTurn() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Bienvenue dans ce combat qu'elle sera votre action ?")
	fmt.Println("1) Menu Attack")
	fmt.Println("2) acceder à l'inventaire")
	fmt.Println("------------------------------------------------")
	var interaction int
	fmt.Scanf("%d\n", &interaction)

	switch interaction {
	case 1:
		C.MenuAttack()
	case 2:
		C.accesinventory()
	}
}

func (C *character) MenuAttack() {
	nb := 1
	fmt.Println("------------------------------------------------")
	fmt.Println("Quelle attaque allez vous lancer ?")
	for i := 0; i < len(C.skill); i++ {
		fmt.Println(nb, ")", C.skill[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	var interaction int
	fmt.Scanf("%d\n", &interaction)

	switch interaction {
	case 1:
		C.verfsort(0)
	case 2:
		C.verfsort(1)
	}
}

func (C *character) verfsort(rang int) {
	if rang >= len(C.skill) {
		return
	}
	if C.skill[rang] == "coup de poing" {
		if M1.pv-8 < 0 {
			M1.pv = 0
			fmt.Println("------------------------------------------------")
			fmt.Println(C1.name, "inflige à", M1.name, 8, "pts de dégat")
			fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
			fmt.Println("------------------------------------------------")
			return
		} else {
			M1.pv -= 8
			fmt.Println("------------------------------------------------")
			fmt.Println(C1.name, "inflige à", M1.name, 8, "pts de dégat")
			fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
			fmt.Println("------------------------------------------------")
		}
	}
	if C.skill[rang] == "boule de feu" {
		if M1.pv-18 < 0 {
			M1.pv = 0
			fmt.Println("------------------------------------------------")
			fmt.Println(C1.name, "inflige à", M1.name, 18, "pts de dégat")
			fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
			fmt.Println("------------------------------------------------")
			return
		} else {
			M1.pv -= 18
			fmt.Println("------------------------------------------------")
			fmt.Println(C1.name, "inflige à", M1.name, 18, "pts de dégat")
			fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
			fmt.Println("------------------------------------------------")
		}
	}
}

func (M *Monstre) gobelinPattern(C *character, T int) {
	Crit := 10
	if M1.pv <= 0 {
		return
	}
	if T%3 == 0 {
		if C.life-Crit <= 0 {
			C.life = 0
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, Crit, "pts de dégat")
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
			return
		} else {
			C.life -= Crit
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, Crit, "pts de dégat")
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		}
		return
	}
	if C.life-M1.dps <= 0 {
		C.life = 0
		fmt.Println(M1.name, "inflige à", C1.name, M1.dps, "pts de dégat")
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		return
	} else {
		C.life -= M1.dps
		fmt.Println(M1.name, "inflige à", C1.name, M1.dps, "pts de dégat")
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
	}
}

func (C *character) charcreation() {
	a := 0
	fmt.Println("importation des données...")
	time.Sleep(2 * time.Second)
	fmt.Println("chargement de la fluctlit : utilisateur #7851")
	time.Sleep(2 * time.Second)
	fmt.Println("/Syncronisation terminé/")
	time.Sleep(2 * time.Second)
	for {
		if a <= 0 {
			fmt.Println("Pour démarrer le jeu, écrivez 'link_start'")
			var interaction string
			fmt.Scanf("%s\n", &interaction)

			switch interaction {
			case "link_start":
				C.start()
				a++
				return
			}
		} else {
			fmt.Println("Pour démarrer le jeu, écrivez 'link_start'")
		}
	}
}

func (C *character) start() {
	var interaction string
	fmt.Println("Comment vous appelez vous ?")
	fmt.Scanf("%s\n", &interaction)
	res := []rune(interaction)
	for i := range res {
		if res[i] >= 65 && res[i] <= 90 || res[i] >= 97 && res[i] <= 122 {
			if res[0] >= 'a' && res[0] <= 'z' {
				res[0] = res[0] - 32
			}
			if res[i] >= 'A' && res[i] <= 'Z' {
				res[i] = res[i] + 32
			}
		}
		if res[i] > 90 && res[i] < 97 || res[i] < 65 || res[i] > 122 {
			res[i] = 0
		}
		C.name = string(res)
	}
	C.start2()
}

func (C *character) start2() {
	var interaction string
	a := 0
	for {
		if a <= 0 {
			fmt.Println("choisissez votre classe : 'Humain', 'Elfe', 'Nain'")
			fmt.Scanf("%s\n", &interaction)

			switch interaction {
			case "Humain":
				C.class = "Humain"
				C.maxlife = 100
				a++
			case "Elfe":
				C.class = "Elfe"
				C.maxlife = 80
				a++
			case "Nain":
				C.class = "Nain"
				C.maxlife = 120
				a++
			}
		} else {
			fmt.Println("choisissez votre classe : 'Humain', 'Elfe', 'Nain'")
		}
		if C.maxlife >= 80 && a > 0 {
			C.life = C.maxlife / 2
			fmt.Println("Bonjour héro, je me présente, je suis Xana et je serai votre guide durant votre aventure")
			fmt.Println("Laissez moi vous aider, pour interagir avec moi, il vous suffit d'écrire le mot qui correspond a vos attentes :)")
			C1.menu()
		}
	}
}

// func ShowMainMenu() {
// 	bold(format: 'Select something :')
// 	PrintLnNtimes( 2)
// }
// color.H1Cyant (format "1: Show character information.")
// fmt.Printf(format)
