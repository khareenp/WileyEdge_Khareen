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
func CreateGoblin(goblin *Goblin) Goblin{
	rand.Seed(time.Now().UnixNano())
	max_h:= 10
	min_h:=5
	max_a:= 3
	min_a:=2
	max_d:= 2
	min_d:=1
	goblin.MaxHealth=  rand.Intn(max_h-min_h) + min_h    // range is min to max
	goblin.Attack=  rand.Intn(max_a-min_a) + min_a    // range is min to max
	goblin.Defense=  rand.Intn(max_d-min_d) + min_d    // range is min to max
	
	
	fmt.Println(*goblin)	
	return *goblin
}


//---------------------------------------------------------------------------------
//Start of game
func GoblinTower(){
	var new_hero Hero
	var choice int

	fmt.Println("Press 1 to start game:")
	fmt.Println("Press 2 to quit game:")
	fmt.Scanf("%d",&choice)

	switch  {
	case choice == 1:
		Game()	
		CreateHero(&new_hero)
	case choice==2:
		End()
	} 
		
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

func End(){
	fmt.Println("Thankyou. Have a nice day!!")
}