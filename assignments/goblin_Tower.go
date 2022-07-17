//Khareen Proverbs
//Goblin Tower Game

package assignments

import (
	"fmt"
	"math"
	"math/rand"
	"rogue/characters"
	"strings"
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
	
	rand.Seed(time.Now().Unix())
	hero := characters.NewHero()
	fmt.Println(hero)

	var response string
	for {
		for i := 1; ; i++ {
			goblin := hero.Step()
			if goblin != nil {
				fmt.Printf("step %v has goblin\n", i)
				fmt.Println(goblin)
				heroWins := hero.Fight(goblin)
				if heroWins {
					hero.RewardFight()
					fmt.Printf("Hero win\n%v\n", hero)
				} else {
					fmt.Print("Goblin win!\n")
					fmt.Printf("Hero dies!\n%v\n", hero)
					break
				}
			} else {
				fmt.Println("nothing on step", i)
			}
			if i%10 == 0 {
				hero.LevelUp()
				fmt.Println(hero)
				fmt.Print("visit potion shop? (y/n) ")
				fmt.Scan(&response)
				response = strings.ToLower(response)
				if response == "y" {
					var numPotions = 0
					fmt.Println("How many potions to buy?")
					fmt.Printf("Bag contains %v potions\n", hero.PotionCount())
					fmt.Printf("You currently have %v gold; 1 potion = 4 gold\n", hero.Gold)
					bagCapacity := len(hero.Potions) - hero.PotionCount()
					maxBuyable := int(math.Min(float64(hero.Gold/4), float64(bagCapacity)))
					fmt.Printf("(Enter number 0-%v): ", maxBuyable)
					fmt.Scan(&numPotions)
					for numPotions > maxBuyable {
						fmt.Print("Can't hold that many. Enter a lower amount: ")
						fmt.Scan(&numPotions)
					}
					fmt.Printf("Bought %v potions\n", numPotions)
					hero.FillPotions(numPotions)
				}
			}
		}
		fmt.Print("Would you like to play again? (y/n)")
		fmt.Scan(&response)
		response = strings.ToLower(response)
		if response != "y" {
			break
		}
		fmt.Println("Making a new character...")
		lastHerosGold := hero.Gold
		hero := characters.NewHero()
		hero.Gold = lastHerosGold
	}

	fmt.Printf("Hero had %v levels, %v gold, and killed %v goblins\n", hero.Level, hero.Gold, hero.GoblinsKilled)
}