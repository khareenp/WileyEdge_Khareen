//Khareen Proverbs
//Goblin Tower Game

package assignments

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct{
	MaxHealth int
	Attack int
	Defense int
	Potions [5]int
	Gold int

}
type Goblin struct{
	MaxHealth int
	Attack int
	Defense int

}


//---------------------------------------------------------------------------------
//Function to create hero
func CreateHero(hero *Hero) Hero{
	rand.Seed(time.Now().UnixNano())
	max_h:= 30
	min_h:=20
	max_a:= 30
	min_a:=20
	max_d:= 30
	min_d:=20
	hero.MaxHealth=  rand.Intn(max_h-min_h) + min_h    // range is min to max
	hero.Attack=  rand.Intn(max_a-min_a) + min_a    // range is min to max
	hero.Defense=  rand.Intn(max_d-min_d) + min_d    // range is min to max
	
	for i := 0; i < len(hero.Potions); i++{
		hero.Potions[i] = rand.Intn(5-1) + 1
	}
	
	fmt.Println(*hero)	
	return *hero
}

//---------------------------------------------------------------------------------
//Function to create goblin
func CreateGoblin(hero *Hero) Hero{
	rand.Seed(time.Now().UnixNano())
	max_h:= 30
	min_h:=20
	max_a:= 30
	min_a:=20
	max_d:= 30
	min_d:=20
	hero.MaxHealth=  rand.Intn(max_h-min_h) + min_h    // range is min to max
	hero.Attack=  rand.Intn(max_a-min_a) + min_a    // range is min to max
	hero.Defense=  rand.Intn(max_d-min_d) + min_d    // range is min to max
	
	for i := 0; i < len(hero.Potions); i++{
		hero.Potions[i] = rand.Intn(5-1) + 1
	}
	
	fmt.Println(*hero)	
	return *hero
}


//---------------------------------------------------------------------------------
//Start of game
func GoblinTower(){
	var new_hero Hero
	Game()
	CreateHero(&new_hero)
   // fmt.Println(new_hero)

}


//---------------------------------------------------------------------------------
//Welcome message
func Game(){
	fmt.Println("**************************************")
	fmt.Println("***             WELCOME            ***")
	fmt.Println("************ GOBLIN TOWER!!! *********")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
}