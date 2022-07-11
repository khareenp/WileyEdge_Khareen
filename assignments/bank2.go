//Khareen Proverbs
//Banking System

package assignments

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type account struct{
	Type string
   	Balance float64
   	FirstName string
   	LastName string
   	UserID int
	Address string
	Interest float64
   	CreatedAt time.Time
}

type wallet struct{
	WalletID string
	Accounts account

}




func Bank2(){

}