/*
https://www.codewars.com/kata/577bd8d4ae2807c64b00045b/train/go

Create a function that returns the name of the winner in a fight between two fighters.

Each fighter takes turns attacking the other and whoever kills the other first is victorious. Death is defined as having health <= 0.

Each fighter will be a Fighter object/instance. See the Fighter class below in your chosen language.

Both health and damagePerAttack (damage_per_attack for python) will be integers larger than 0. You can mutate the Fighter objects.

Your function also receives a third argument, a string, with the name of the fighter that attacks first.

Example:
  declare_winner(Fighter("Lew", 10, 2), Fighter("Harry", 5, 4), "Lew") => "Lew"

  Lew attacks Harry; Harry now has 3 health.
  Harry attacks Lew; Lew now has 6 health.
  Lew attacks Harry; Harry now has 1 health.
  Harry attacks Lew; Lew now has 2 health.
  Lew attacks Harry: Harry now has -1 health and is dead. Lew wins.
type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

*/

package main

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			firstAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			firstAttacker = fighter1.Name
		}

		if fighter1.Health <= 0 {
			return fighter2.Name
		}

		if fighter2.Health <= 0 {
			return fighter1.Name
		}
	}
}

