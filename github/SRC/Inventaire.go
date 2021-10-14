package main

import "fmt"

//   Commandes liées à l'inventaire   //
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
		fmt.Println("Vous revenez au menu.")
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
		fmt.Println("Vous revenez au menu.")
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
