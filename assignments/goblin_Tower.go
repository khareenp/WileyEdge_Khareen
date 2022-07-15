//Khareen Proverbs
//Goblin Tower Game

package assignments

import (
	"fmt"
	"math/rand"
	"time"
)

// NOTE: We will need to keep track of how many Goblins were slain.
// NOTE: We have a MaxHealth, but what about the current health?
// NOTE: Do we need a Level field in here?
type Hero struct {
	MaxHealth int
	Attack    int
	Defense   int
	Potions   [5]int
	Gold      int
	CurrentHealth int
	Level []int
}

// NOTE: We have a MaxHealth here, but can the Goblin take potions? Or does their health monotonically decrease?
// In that case, maybe we just need current health.
// Or you can add a way for the Goblin to take a potion, suck life from Hero, etc.
type Goblin struct {
	CurentHealth int
	Attack    int
	Defense   int
}

//---------------------------------------------------------------------------------
//Function to create hero
// NOTE: This function takes a pointer, modifies the value at it, and returns that value. Thoughts?
func CreateHero() Hero {
	rand.Seed(time.Now().UnixNano())
	// NOTE: I've updated some of these to the assigned amounts.
	max_h := 30
	min_h := 20
	max_a := 3
	min_a := 1
	max_d := 1
	min_d := 5
	hero.MaxHealth = rand.Intn(max_h-min_h) + min_h // range is min to max
	hero.Attack = rand.Intn(max_a-min_a) + min_a    // range is min to max
	hero.Defense = rand.Intn(max_d-min_d) + min_d   // range is min to max

	for i := 0; i < len(hero.Potions); i++ {
		hero.Potions[i] = 2 // NOTE: The hero gets 5 2-health potions to start, not random-health potions.
	}

	fmt.Println(hero)
	return hero
}

//---------------------------------------------------------------------------------
//Function to create goblin
// NOTE: This function takes a pointer, modifies the value at it, and returns that value. Thoughts?
func CreateGoblin(goblin Goblin) Goblin {
	rand.Seed(time.Now().UnixNano())
	max_h := 10
	min_h := 5
	max_a := 3
	min_a := 2
	max_d := 2
	min_d := 1
	goblin.MaxHealth = rand.Intn(max_h-min_h) + min_h // range is min to max
	goblin.Attack = rand.Intn(max_a-min_a) + min_a    // range is min to max
	goblin.Defense = rand.Intn(max_d-min_d) + min_d   // range is min to max

	fmt.Println(goblin)
	return goblin
}

// NOTE: Probably we need some methods or functions to make things happen to the Hero and Goblins.
// Whether we choose to make these methods or functions is a matter of style.
func (hero *Hero) Fight(goblin *Goblin) {
	goblinturn:=rand.Intn(10-1) + 1    // range is min to max
	heroturn:=rand.Intn(10-1) + 1    // range is min to max
	if heroturn > goblinturn{
		//hero goes first
		for _,i:= range hero.MaxHealth{
			hero.CurrentHealth = hero.MaxHealth
			if hero.Defense < goblin.Attack
			hero.CurrentHealth = hero.CurrentHealth - goblin.Attack
		}
	}else{
		//goblin goes first
		goblin.CurrentHealth
		if hero
	}
	// NOTE: Who goes first?
	// Up to you. You can use a random number, compare health/level, Hero always goes first, etc.
	// While both are alive, keep fighting.
	// Does the Hero have the option to drink potions in here?
	// If the Goblin dies, the Hero gets some Gold and the kill is counted.
	// If the Hero dies, game over.
	// Does this need to return anything?
}

func (hero *Hero) TakePotion() {
	// NOTE: Check if we have any potions left.
	// If we do, take one, and restore some health.
	// Decrement the number of potions we have left.
	// Maybe say something to the user about what happened?
}

func (hero *Hero) VisitPotionShop() {
	// NOTE: Do we need this method?
	// Could we just let the player buy potions inline in the loop of some other function?
	// What if we want to let them buy multiple potions in one visit?
	// The design is up to you.
}

func (hero *Hero) BuyPotion() {
	// NOTE: Check if we have enough gold.
	// Check if we have room in our inventory for a new potion.
	// If we do, buy a potion, and add it to our inventory.
	// Decrement the number of Gold we have left.
	// Maybe say something to the user about what happened?
}


//---------------------------------------------------------------------------------
//Start of game
func GoblinTower() {
	// NOTE: The player gets to play as long as they want, so we will need a main game loop in here somewhere.
	// The loop should create a new Hero with random stats except for the Gold, which should be set to the previous run's Gold when
	// the Hero died.
	var new_hero Hero
	var choice int

	Game()
	fmt.Println("Press 1 to start game:")
	fmt.Println("Press 2 to quit game:")
	fmt.Scanf("%d", &choice)

	switch {
	case choice == 1:
		CreateHero(new_hero)
		// NOTE: Here we could have some kind of `EnterTower()` function (or just inline it) to loop over the Hero walking and 
		// encountering Goblins, etc.
	case choice == 2:
		End()
	}
}

// NOTE: We are not given an encounter rate, but you can play with it and find a rate that allows the player to win sometimes.
var encounterRate float64 = 0.99

func EnterTower(hero *Hero) {
	// NOTE: The Hero needs to take steps and randomly encounter Goblins, fight them, and be able to drink potions.
	// NOTE: We need to keep track of the number of steps the Hero has taken to decide if he can visit the potion shop.
	currentSteps := 0
	var choice int
	for {
		// Let's find out what the player wants to do.
		fmt.Println("Possible actions: ")
		fmt.Println("1. Take a step.")
		if len(hero.Potions) > 0 {
			fmt.Println("2. Drink a potion.")
		}
		// NOTE: You could implement this a number of ways.
		// Do you want the hero to have to wait at least another 10 steps after every visit the potion shop?
		// Or do you want them to be able to visit the potion shop 3 times in a row if they have reached level 3 for example?
		// Or do you want to make vising the potion shop an option only on the step that they attain a new level?
		// How many potions can the Hero buy at the potion shop in one visit?
		// Up to you.
		if {// SOME_CONDITION 
			fmt.Println("3. Visit the potion shop.")
		}
	
		// NOTE: Maybe another option would be to just quit the game.
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			NOTE: // Take a step and count it.
			if rand.Float64() < encounterRate {
				// NOTE: We need to create a Goblin here.
				var new_goblin Goblin
				CreateGoblin(&new_goblin)
				// NOTE: The Hero needs to fight the Goblin.
				// If he wins: 
				// 		record the slain Goblin and give the Hero some gold.
				// If he loses, 
				// 		end this run, show the user the game over screen, and ask if they want to play again.
				// 		don't forget to record the amount of Gold the Hero had for the next run.
				// So do we want to reward the Hero with gold right here?
				// Or should we do it in a method we create to have the Hero and Goblin?
				// Where should we record the number of Goblins the Hero has killed?
				// Should that be a global variable or live on the Hero struct?
				// Many ways are acceptable. How would you like to think about it?
			}
		// NOTE: Probably want to check that the player has met the same conditions we had to show the options for these in the first place.
		case 2:
			// NOTE: The Hero wants to drink a potion.
		case 3:
			// NOTE: The Hero wants to visit the potion shop.
		default:
			fmt.Println("Please select one of the actions.")
		}
	}
}

//---------------------------------------------------------------------------------
//Welcome message
func Game() {
	fmt.Println("**************************************")
	fmt.Println("***             WELCOME            ***")
	fmt.Println("************ GOBLIN TOWER!!! *********")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
}

func End() {
	// NOTE: Print the number of Goblins slain and the Level attained by the Hero.
	fmt.Println("Thank you. Have a nice day!!")
}

// NOTE: We always need a main function.
func main() {
	// We can just call the GoblinTower function in here though.
	GoblinTower()
}