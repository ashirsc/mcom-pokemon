package main

import "log"

type Monster struct {
	health  int
	attack  int
	defense int
}

func main() {

	mon1 := Monster{10, 2, 2}
	mon2 := Monster{25, 1, 1}

	log.Println("before")
	log.Println(mon2.health)
	mon2 = attack(mon1, mon2)
	log.Println("after")
	log.Println(mon2.health)

}

func attack(m1 Monster, m2 Monster) Monster {
	m2.health -= m1.attack
	return m2
}
