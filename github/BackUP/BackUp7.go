package main

import (
	"fmt"
	"os"
	"time"
)

type character struct {
	name       string
	class      string
	lvl        int
	lvlup      int
	maxlife    int
	life       int
	manamax    int
	mana       int
	initiative int
	money      int
	slotinv    int
	inventory  []string
	skill      []string
	equip      stuff
}

type stuff struct {
	head string
	body string
	foot string
	hand string
}

type Monstre struct {
	name       string
	pvmax      int
	pv         int
	damage     int
	initiative int
}

func (C *character) Init(name string, class string, lvl int, lvlup int, maxlife int, life int, manamax int, mana int, initiative int, money int, slotinv int, inventory []string, skill []string, head string, body string, foot string, hand string) {
	C.name = name
	C.class = class
	C.lvl = lvl
	C.lvlup = lvlup
	C.maxlife = maxlife
	C.life = life
	C.manamax = manamax
	C.mana = mana
	C.initiative = initiative
	C.slotinv = slotinv
	C.inventory = inventory
	C.money = money
	C.skill = skill
	C.equip.head = head
	C.equip.body = body
	C.equip.foot = foot
	C.equip.hand = hand
}

func (M *Monstre) Init(name string, pvmax int, pv int, damage int, initiative int) {
	M.name = name
	M.pvmax = pvmax
	M.pv = pv
	M.damage = damage
	M.initiative = initiative
}

//   Variable globale   //
var T = 1
var O = 0
var popogratuite int
var C *character = &character{}
var C1 character = character{}
var M *Monstre = &Monstre{}
var M1 Monstre = Monstre{}

func main() {
	C1.Init(C.name, C.class, 1, 1000, C.maxlife, 1, C.manamax, C.mana, C.initiative, 100, 10, []string{"potion de santé", "potion de santé", "potion de santé"}, []string{"coup de poing"}, "empty", "empty", "empty", "empty")
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

//   Initialisation dans le jeu   //
func (C *character) charcreation() {
	a := 0
	fmt.Println("importation des données...")
	time.Sleep(1 * time.Second)
	fmt.Println("chargement de la fluctlit : utilisateur #7851")
	time.Sleep(1 * time.Second)
	fmt.Println("/Syncronisation terminé/")
	time.Sleep(1 * time.Second)
	for {
		if a <= 0 {
			fmt.Println("Pour rentrer dans le monde virtuel, écrivez 'link_start'")
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

//   Création du personnage : nom   //
func (C *character) start() {
	g := 0
	a := 0
	for g == 0 {
		var interaction string
		fmt.Println("Comment vous appelez vous ? (seulement les lettres sont acceptées, tout le reste sera supprimé)")
		fmt.Scanf("%s\n", &interaction)
		res := []rune(interaction)
		for i := range res {
			if res[i] >= 65 && res[i] <= 90 || res[i] >= 97 && res[i] <= 122 {
				for a <= 0 {
					if res[i] >= 'a' && res[i] <= 'z' {
						res[i] -= 32
						fmt.Println(res[i])
						a++
					} else if a > 0 && res[i] >= 'A' && res[i] <= 'Z' {
						res[i] += 32
					}
				}
			}
			if res[i] > 90 && res[i] < 97 || res[i] < 65 || res[i] > 122 {
				res[i] = 0
			}
			g++
		}
		C.name = string(res)
	}
	C.start2()
}

//   Création du personnage : classe   //
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
				C.manamax = 50
				C.mana = 40
				C.initiative = 2
				a++
			case "Elfe":
				C.class = "Elfe"
				C.maxlife = 80
				C.manamax = 70
				C.mana = 60
				C.initiative = 1
				a++
			case "Nain":
				C.class = "Nain"
				C.maxlife = 120
				C.manamax = 30
				C.mana = 20
				C.initiative = 3
				a++
			}
		} else {
			fmt.Println("choisissez votre classe : 'Humain', 'Elfe', 'Nain'")
		}
		if C.maxlife >= 80 && a > 0 {
			C.life = C.maxlife / 2
			fmt.Println("Bonjour héros, je me présente, je suis Xana et je serai votre guide durant votre aventure.")
			fmt.Println("Laissez moi vous aider, pour interagir avec moi, il vous suffit d'écrire le mot qui correspond a vos attentes :)")
			break
		}
	}
}

//   Menu   //
func (C1 *character) menu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("info. accéder à vos informations")
	fmt.Println("inv. accéder à votre inventaire")
	fmt.Println("stuff. gérer vos équipements")
	fmt.Println("trade. parler au marchand")
	fmt.Println("forge. permet d'accéder à la forge")
	fmt.Println("fight. permet de s'excercer au combat")
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
	case "fight":
		C1.TrainingFight(M)
	case "left":
		os.Exit(0)
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Info personnage   //
func (C *character) displayinfo() {
	tab := ""
	fmt.Println("------------------------------------------------")
	fmt.Println("name :", C.name)
	fmt.Println("class :", C.class)
	fmt.Println("lvl :", C.lvl)
	fmt.Println("lvl up :", C.lvlup, "point d'xp restant")
	fmt.Println("maxlife :", C.maxlife)
	fmt.Println("life :", C.life)
	fmt.Println("maxmana :", C.manamax)
	fmt.Println("mana :", C.mana)
	for i := 0; i < len(C.skill); i++ {
		if i > 0 {
			tab += ", "
		}
		tab += "["
		tab += C.skill[i]
		tab += "]"
	}
	fmt.Println("skill :", tab)
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   mort du perso   //
func (C *character) dead() {
	if C.life <= 0 {
		println("VOUS ETES MORT")
		time.Sleep(2 * time.Second)
		println("Vous revenez à la vie avec 50% de vos PV")
		C.life = C.maxlife / 2
	}
}

//   commandes liées à l'inventaire   //
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
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Permet de voir l'inventaire   //
func (C *character) see() {
	nb := 1
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous possedez", C.money, "pièces d'or")
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
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Utiliser un objet dans l'inventaire   //
func (C *character) use() {
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("écrivez le numérot qui correspond à l'objet que vous voulez utiliser")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case 0:
		return
	case k:
		if k < 0 || k > len(C.inventory) {
			fmt.Println("Je n'ai pas compris votre requête")
		} else {
			C.verf(k - 1)
		}
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Utiliser potion santé  //
func (C *character) takeapot() {
	if C.life == C.maxlife {
		fmt.Println("vous etes déjà full pv")
		return
	}
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "potion de santé" {
			break
		}
	}
	if C.life >= C.maxlife-50 {
		fmt.Println("Vous utilisez une potion de santé, vous regagnez 50 pv")
		C.life = C.maxlife
		C.removestuff(i)
	} else {
		fmt.Println("Vous utilisez une potion de santé, vous regagnez 50 pv")
		C.life += 50
		C.removestuff(i)
	}
}

//   Utiliser potion mana   //
func (C *character) PotiondeMana() {
	if C.mana == C.manamax {
		fmt.Println("vous etes déjà full mana")
		return
	}
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "potion de mana" {
			break
		}
	}
	if C.mana >= C.manamax-25 {
		fmt.Println("Vous utilisez une potion de mana, vous regagnez 25 pts de mana")
		C.mana = C.manamax
		C.removestuff(i)
	} else {
		fmt.Println("Vous utilisez une potion de mana, vous regagnez 25 pts de mana")
		C.mana += 25
		C.removestuff(i)
	}
}

//   poison sur perso  //
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

//   Utiliser un objet en combat   //
func (C *character) use2() {
	nb := 1
	var k int
	a := 0
	for a == 0 {
		nb = 1
		fmt.Println("------------------------------------------------")
		for i := 0; i < len(C.inventory); i++ {
			if C.inventory[i] == "potion de santé" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "potion de mana" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "fiole de poison" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "livre de sort (boule de feu)" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "livre de sort (pichenette)" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			}
		}
		fmt.Println("------------------------------------------------")
		fmt.Println("écrivez le numéro qui correspond à l'objet que vous voulez utiliser")
		fmt.Println("utilisez '0' pour revenir en arrière")
		fmt.Println("------------------------------------------------")
		fmt.Scanf("%d\n", &k)
		switch k {
		case 0:
			return
		case k:
			if k > len(C.inventory) || k < 1 {
				a = 0
			} else if C.verf2(k - 1) {
				fmt.Println("------------------------------------------------")
				a++
				M.gobelinPattern(C)
			}
		default:
			fmt.Println("Je n'ai pas compris votre requête")
		}
	}
}

//   Pour supprimer un objet   //
func (C *character) delete() {
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("écrivez le numérot qui correspond à l'objet que vous voulez supprimer")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case 0:
		return
	case k:
		C.remove(C.inventory[k-1])
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Ajouter un objet   //
func (C *character) add(object string, price int) {
	if popogratuite == 0 && object == "potion de santé" {
		price = 0
		println("En tant que bon vendeur, je vous offre votre premiere potion de santé")
		popogratuite++
	}
	if object == "livre de sort (pichenette)" {
		println("Le", object, "a été ajouté à votre inventaire.")
		C.inventory = append(C.inventory, object)
		return
	}
	if object == "livre de sort (tatane celeste)" {
		println("Le", object, "a été ajouté à votre inventaire.")
		C.inventory = append(C.inventory, object)
		return
	}
	if object == "épée rouillée" {
		println("L'", object, "a été ajoutée à votre inventaire.")
		C.inventory = append(C.inventory, object)
		return
	}
	if len(C.inventory) >= C.slotinv {
		fmt.Println("votre inventaire est plein")
		return
	}
	if C.money > price {
		C.money -= price
		C.inventory = append(C.inventory, object)
		// sortie = append(sortie, object)
		fmt.Println(object, "a été ajouté à votre inventaire")
	} else {
		fmt.Println("Vous n'avez pas assez d'argent")
	}
}

//   Ajouter un équipement   //
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
	if equip == "épée empoisonnée" && C.equip.hand != "épée empoisonnée" {
		C.equip.hand = "épée empoisonnée"
		fmt.Println("Vous vous équipez de l'", equip)
		C.remove("épée empoisonnée")
		C.skill = append(C.skill, "épée empoisonnée")
		return
	}
	println("Vous etes déja équipé de cet objet")
}

//   Supprimer un objet   //
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

//   Supprimer un équipement   //
func (C *character) removestuff(number int) {
	if number < len(C.inventory)-1 {
		C.inventory = append(C.inventory[:number], C.inventory[number+1:]...)
	} else if number == len(C.inventory)-1 {
		C.inventory = C.inventory[:number]
	}
}

//   Pour voir ses équipements   //
func (C *character) stuff() {
	fmt.Println("------------------------------------------------")
	fmt.Println("head stuff :", C.equip.head)
	fmt.Println("body stuff :", C.equip.body)
	fmt.Println("foot stuff :", C.equip.foot)
	fmt.Println("hand stuff :", C.equip.hand)
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Pour marchander   //
func (C *character) trader() {
	fmt.Println("Marchand itinérant : jette un coup d'oeil à ma boutique, j'ai plein de choses à vendre.")
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous possedez", C.money, "pièces d'or")
	fmt.Println("------------------------------------------------")
	fmt.Println("1) potion de santé: Une potion qui rend 50 pv [3 po]")
	fmt.Println("2) potion de mana: Une potion qui rend 25 points de mana [3 po]")
	fmt.Println("3) fiole de poison: Une fiole de poison qui inflige 10 pv pendant 3s [6 po]")
	fmt.Println("4) livre de sort (boule de feu): un sort qui permet de projeter une boule de feu [25 po]")
	fmt.Println("5) Fourrure de Loup: une fourrure banale de gentil chien-chien[4 po]")
	fmt.Println("6) Peau de Troll: une belle peau fraichement dépecée d'un troll [7 po]")
	fmt.Println("7) Cuir de Sanglier: un cuir plutot résistant [3 po]")
	fmt.Println("8) Plume de Corbeau: une couleur magnifique pour une simple plume noire, non ? [1 po]")
	fmt.Println("9) augmentation d'inventaire: ajoute 10 places à votre inventaire [30 po]")
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "0":
		return
	case "1":
		C.add("potion de santé", 3)
		popogratuite++
	case "2":
		C.add("potion de mana", 3)
	case "3":
		C.add("fiole de poison", 6)
	case "4":
		for i := 0; i < len(C.inventory); i++ {
			if C.inventory[i] == "livre de sort (boule de feu)" {
				println("Vous avez déjà ce sort en votre possession.")
				return
			}
		}
		for i := 0; i < len(C.skill); i++ {
			if C.skill[i] == "boule de feu" {
				println("Vous avez déjà appris ce sort.")
				return
			}
		}
		C.add("livre de sort (boule de feu)", 25)
	case "5":
		C.add("fourrure de loup", 10)
	case "6":
		C.add("peau de troll", 7)
	case "7":
		C.add("cuir de sanglier", 3)
	case "8":
		C.add("plume de corbeau", 1)
	case "9":
		C.upgradeinventoryslot(30)
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Augmentation d'inventaire   //
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

//   Apprendre un nouveau sort  //
func (C *character) spellBook(sort string, rang int) {
	for i := 0; i < len(C.skill); i++ {
	}
	C.skill = append(C.skill, sort)
	println("Vous avez appris un nouveau sort")
	C.removestuff(rang)
}

//   Forger un équipement   //
func (C *character) forge() {
	if O == 2 {
		fmt.Println("forgeron : comme tu est un bon client, je te confie cette épée que j'ai trouvé ce matin.")
		time.Sleep(2 * time.Second)
		C.add("épée rouillée", 0)
		time.Sleep(2 * time.Second)
		fmt.Println("N'hésite pas à revenir me voir")
		time.Sleep(2 * time.Second)
		O++
		return
	} else {
		fmt.Println("Forgeron : bienvenu dans ma forge, je peux vous craft ce que vous voulez contre 5 pièces d'or, enfin que si vous m'apportez les mattériaux.")
		fmt.Println("------------------------------------------------")
		fmt.Println("Vous possedez", C.money, "pièces d'or")
		fmt.Println("------------------------------------------------")
		fmt.Println("1) chapeau de l'aventurier (1 plume de corbeau, 1 cuir de sanglier)")
		fmt.Println("2) tunique de l’aventurier (2 fourrure de loup, 1 peau de troll)")
		fmt.Println("3) bottes de l’aventurier (1 fourrure de loup, 1 cuir de sanglier)")
		fmt.Println("4) épée empoisonnée (1 épee rouillée, 1 fiole de poison)")
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
		case "4":
			C.craft("épée rouillée", "fiole de poison", "épée empoisonnée")
		case "return":
			return
		default:
			fmt.Println("Je n'ai pas compris votre requête")
		}
	}
}

//   Vérifier les objets avant de forger l'équipement   //
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

//   Trouver un objet en fonction de son rang dans l'inventaire puis l'utiliser   //
func (C *character) verf(rang int) {
	if rang >= len(C.inventory) {
		return
	}
	if C.inventory[rang] == "potion de santé" {
		C.takeapot()
	} else if C.inventory[rang] == "potion de mana" {
		C.PotiondeMana()
	} else if C.inventory[rang] == "fiole de poison" {
		C.poison()
		C.removestuff(rang)
	} else if C.inventory[rang] == "livre de sort (boule de feu)" {
		C.spellBook("boule de feu", rang)
	} else if C.inventory[rang] == "livre de sort (pichenette)" {
		C.spellBook("pichenette", rang)
	} else if C.inventory[rang] == "livre de sort (tatane celeste)" {
		C.spellBook("tatane celeste", rang)
	} else if C.inventory[rang] == "chapeau de l'aventurier" {
		C.addstuff("chapeau de l'aventurier")
	} else if C.inventory[rang] == "tunique de l'aventurier" {
		C.addstuff("tunique de l'aventurier")
	} else if C.inventory[rang] == "bottes de l'aventurier" {
		C.addstuff("bottes de l'aventurier")
	} else if C.inventory[rang] == "épée empoisonnée" {
		C.addstuff("épée empoisonnée")
	}
}

//   Trouver un objet en fonction de son rang dans l'inventaire puis l'utiliser dans un combat (éviter d'utiliser un objet inutile en combat)   //
func (C *character) verf2(rang int) bool {
	if C.inventory[rang] == "potion de santé" && C.life != C.maxlife {
		C.takeapot()
		return true
	} else if C.inventory[rang] == "potion de mana" && C.mana != C.manamax {
		C.PotiondeMana()
		return true
	} else if C.inventory[rang] == "fiole de poison" {
		C.poison()
		C.removestuff(rang)
		return true
	} else if C.inventory[rang] == "livre de sort (boule de feu)" {
		C.spellBook("boule de feu", rang)
		return true
	} else if C.inventory[rang] == "livre de sort (pichenette)" {
		C.spellBook("pichenette", rang)
		return true
	} else {
		fmt.Println("vous etes déjà full")
		return false
	}
}

//   Gère les monstres, le nombre de tours, la fin des combats et la fin du jeu   //
func (C *character) TrainingFight(M *Monstre) {
	if O == 0 {
		M1.name = "Gobelin d'entrainement"
		M1.pvmax = 40
		M1.pv = 40
		M1.damage = 5
		M1.initiative = 1
	} else if O == 1 {
		M1.name = "Gobelin supérieur"
		M1.pvmax = 80
		M1.pv = 80
		M1.damage = 10
		M1.initiative = 3
	} else if O >= 2 {
		M1.name = "Roi des Gobelins"
		M1.pvmax = 100
		M1.pv = 100
		M1.damage = 15
		M1.initiative = 5
	}
	if O < 2 {
		fmt.Println("Vous entrez dans un donjon, un", M1.name, "vous attaque.")
	} else {
		fmt.Println("Vous entrez dans un chateau, le", M1.name, "vous fait face.")
		time.Sleep(1 * time.Second)
		fmt.Println("Vous vous sentez intimidé par sa présence")
		time.Sleep(1 * time.Second)
		fmt.Println("Vos stats baissent légèrement")
		C.maxlife -= 5
		C.manamax -= 5
	}
	for {
		if T == 1 {
			M1.pv = M1.pvmax
		}
		if C1.life > 0 && M1.pv > 0 {
			fmt.Println("------------------------------------------------")
			fmt.Println("Vous êtes au tour n°", T)
			if !C.charTurn() {
			} else {
				break
			}
		} else {
			if M1.pv == 0 {
				fmt.Println("Vous avez vaincu le", M1.name)
				time.Sleep(2 * time.Second)
				if O == 0 {
					fmt.Println("Vous récuperez 5 pièces d'or.")
					C.money += 5
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 1000 points d'expérience")
					C.xp(1000)
				} else if O == 1 {
					fmt.Println("Vous récuperez 10 pièces d'or.")
					C.money += 10
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 1250 points d'expérience")
					C.xp(1250)
				} else if O == 2 {
					fmt.Println("Vous récuperez 20 pièces d'or.")
					C.money += 20
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 2000 points d'expérience")
					C.xp(2000)
				}
				time.Sleep(2 * time.Second)
				fmt.Println("xp nécessaire pour lvl up :", C.lvlup)
				fmt.Println("------------------------------------------------")
				if O == 0 {
					time.Sleep(2 * time.Second)
					fmt.Println("Le Gobelin d'entrainement laisse tomber derrière lui le livre de sort : pichenette")
					time.Sleep(2 * time.Second)
					C.add("livre de sort (pichenette)", 0)
				} else if O == 1 {
					time.Sleep(2 * time.Second)
					fmt.Println("Le Gobelin supérieur laisse tomber derrière lui le livre de sort : tatane celeste")
					time.Sleep(2 * time.Second)
					C.add("livre de sort (tatane celeste)", 0)
				}
				time.Sleep(2 * time.Second)
				if O < 2 {
					fmt.Println("Vous vous dirigez victorieusement vers la sortie mais votre soif de combat n'a pas été rassasiée.")
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Vous avez vaincu le Roi Gobelin, mais l'intensité de votre combat vous à fait vous évanouir sur place.")
					time.Sleep(3 * time.Second)
					fmt.Println("En fermant les yeux, vous appercevez les larbins du roi qui foncent vers vous.")
					time.Sleep(3 * time.Second)
					fmt.Println("Vous sentez votre fin approcher mais une mystérieuse personne encapuchonnée viens à votre secours...")
					time.Sleep(3 * time.Second)
					fmt.Println("Vous vous reveillez dans une foret au coté de cette personne.")
					time.Sleep(3 * time.Second)
					fmt.Println("Vous vous demandez pourquoi vous à t-elle sauvé.")
					time.Sleep(3 * time.Second)
					fmt.Println("Soudain, la personne retire sa capuche affirme que vous devez vous reveiller pour continuer de bosser sur Ytrack.")
					time.Sleep(3 * time.Second)
					fmt.Println("La suite dans la partie 2...")
					fmt.Println("------------------------------------------------")
					time.Sleep(2 * time.Second)
					fmt.Println("Merci d'avoir joué a notre jeu, nous espérons que ce petit RPG vous aura bien diverti.")
					fmt.Println("------------------------------------------------")
					time.Sleep(2 * time.Second)
					fmt.Println("Crédit")
					time.Sleep(2 * time.Second)
					// blague à Nolan
					time.Sleep(4 * time.Second)
				}
				O++
				T = 1
				break
			}
			if C1.life == 0 {
				fmt.Println("------------------------------------------------")
				time.Sleep(2 * time.Second)
				fmt.Println("Le ", M1.name, "vous a mis une raclée.")
				time.Sleep(2 * time.Second)
				C.dead()
				time.Sleep(2 * time.Second)
				fmt.Println("Vous reprenez connaisance en dehors du donjon.")
				time.Sleep(2 * time.Second)
				T = 1
				break
			}
		}
	}
}

//   Choisir ce que l'on veut faire dans un combat   //
func (C *character) charTurn() bool {
	fmt.Println("------------------------------------------------")
	time.Sleep(1 * time.Second)
	fmt.Println("1) Menu Attack")
	fmt.Println("2) acceder à l'inventaire")
	fmt.Println("3) prendre la fuite ")
	fmt.Println("------------------------------------------------")
	var interaction int
	fmt.Scanf("%d\n", &interaction)

	switch interaction {
	case 1:
		C.MenuAttack()
		return false
	case 2:
		C.use2()
		return false
	case 3:
		T = 1
		fmt.Println("Vous prenez la fuite tel un misérable gueux apeuré.")
		time.Sleep(2 * time.Second)
		return true
	}
	fmt.Println("Je n'ai pas compris votre requête")
	return false
}

//   Choisir une attaque   //
func (C *character) MenuAttack() {
	a := 0
	nb := 1
	for a == 0 {
		nb = 1
		fmt.Println("------------------------------------------------")
		fmt.Println("Quelle attaque allez vous lancer ?")
		fmt.Println("Votre mana actuel est de :", "[", C.mana, "]")
		fmt.Println("------------------------------------------------")
		for i := 0; i < len(C.skill); i++ {
			fmt.Print(nb, ")", C.skill[i])
			if C.skill[i] == "coup de poing" {
				fmt.Println(" 8 damage [5 mana]")
			} else if C.skill[i] == "épée empoisonnée" {
				fmt.Println(" 15 damage [10 mana]")
			} else if C.skill[i] == "boule de feu" {
				fmt.Println(" 18 damage [15 mana]")
			} else if C.skill[i] == "pichenette" {
				fmt.Println(" 5 damage [0 mana]")
			} else if C.skill[i] == "tatane celeste" {
				fmt.Println(" 15 damage [12 mana]")
			}
			nb++
		}
		fmt.Println("------------------------------------------------")
		fmt.Println("écrivez le numéro qui correspond à l'attaque que vous voulez utiliser.")
		fmt.Println("utilisez '0' pour revenir en arrière")
		fmt.Println("------------------------------------------------")
		var k int
		fmt.Scanf("%d\n", &k)

		switch k {
		case 0:
			return
		case k:
			if k > len(C.skill) || k < 1 {
				a = 0
			} else {
				C.verifsort(C.skill[k-1])
				a++
			}
		}
	}
}

//   Utiliser une attaque et vérifier le mana nécéssaire pour la lancer   //
func (C *character) verifsort(rang string) {
	mana := 0
	deg := 0
	if rang == "coup de poing" {
		mana = 5
		deg = 8
	} else if rang == "épée empoisonnée" {
		mana = 5
		deg = 5
	} else if rang == "boule de feu" {
		mana = 15
		deg = 18
	} else if rang == "pichenette" {
		mana = 0
		deg = 5
	} else if rang == "tatane celeste" {
		mana = 10
		deg = 15
	}
	if C.mana < mana {
		fmt.Println("Vous n'avez pas asser de mana.")
		time.Sleep(2 * time.Second)
		C.MenuAttack()
		return
	} else {
		if !C.ynitiative() {
			fmt.Println("Vous avez moins d'initiative que votre adversaire, vous attaquez en dernier.")
			fmt.Println("------------------------------------------------")
			M.gobelinPattern(C)
			fmt.Println("------------------------------------------------")
		} else {
			fmt.Println("Vous avez plus d'initiative que votre adversaire, vous attaquez en premier.")
			fmt.Println("------------------------------------------------")
		}
		C.mana -= mana
		time.Sleep(1 * time.Second)
		fmt.Println(C1.name, "inflige à", M1.name, deg, "pts de dégat", "votre Mana restant est de", C.mana, "/", C.manamax)
		time.Sleep(1 * time.Second)
		if M1.pv-deg < 0 {
			M1.pv = 0
		} else {
			M1.pv -= deg
			if rang == "épée empoisonnée" {
				seconds := 2
				fmt.Println("Vous avez empoisonnée le", M1.name)
				for s := seconds; s >= 1; s-- {
					if M1.pv-5 < 0 {
						M1.pv = 0
						time.Sleep(2 * time.Second)
						fmt.Println("-5 pv")
						fmt.Println("life :", M1.pv)
						break
					}
					time.Sleep(2 * time.Second)
					fmt.Println("-5 pv")
					M1.pv -= 5
					fmt.Println("life :", M1.pv)
				}
			}
		}
		fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
		time.Sleep(1 * time.Second)
		if C.ynitiative() {
			fmt.Println("------------------------------------------------")
			M.gobelinPattern(C)
		}
		return
	}
}

//   Permet d'attaquer le joueur   //
func (M *Monstre) gobelinPattern(C *character) {
	if M1.pv <= 0 {
		return
	}
	if T%3 == 0 {
		if C.life-(M1.damage*2) <= 0 {
			C.life = 0
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, (M1.damage * 2), "pts de dégat")
			time.Sleep(1 * time.Second)
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
			time.Sleep(1 * time.Second)
			return
		} else {
			C.life -= (M1.damage * 2)
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, (M1.damage * 2), "pts de dégat")
			time.Sleep(1 * time.Second)
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
			time.Sleep(1 * time.Second)
			T++
			return
		}
	}
	if C.life-M1.damage <= 0 {
		C.life = 0
		fmt.Println(M1.name, "inflige à", C1.name, M1.damage, "pts de dégat")
		time.Sleep(1 * time.Second)
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		time.Sleep(1 * time.Second)
		return
	} else {
		C.life -= M1.damage
		fmt.Println(M1.name, "inflige à", C1.name, M1.damage, "pts de dégat")
		time.Sleep(1 * time.Second)
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		time.Sleep(1 * time.Second)
		T++
	}
}

//   Lvl up le personnage   //
func (C *character) xp(nb int) {
	for nb > 0 {
		if nb >= C.lvlup {
			nb = nb - C.lvlup
			C.lvl += 1
			C.lvlup += 250
			C.maxlife += 5
			C.manamax += 5
			C.initiative += 1
			fmt.Println("Vous venez de gagner un niveau")
			fmt.Println("Vous vous sentez plus fort")
		} else {
			C.lvlup -= nb
			nb = 0
		}
	}
}

//   Pour savoir qui attaque en premier   //
func (C *character) ynitiative() bool {
	if C.initiative >= M1.initiative {
		return true
	} else {
		return false
	}
}
