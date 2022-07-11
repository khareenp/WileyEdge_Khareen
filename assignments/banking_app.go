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

// create the struct( onstructor in java)
type Account struct {
   Type string
   Balance float32
   FirstName string
   LastName string
   UserID int
   CreatedAt time.Time
}

//---------------------------------------------------------------------------------
func (x Account) DisplayBalance()float32{
	fmt.Printf("Your balance is: %.2f", x.Balance)
	//fmt.Println("Transaction completed at")
    return x.Balance
}

//---------------------------------------------------------------------------------
func Deposit(x *Account){
	var amount float64
	deposit:
	fmt.Println("Please enter amount to deposit: ")
	fmt.Scanf("%f",&amount)
	balance := float64(x.Balance)+ amount

	if amount <= 0  {
		fmt.Println("You have entered insuffiction balance, please try again")
		goto deposit
	}else{
		fmt.Printf("You have made a total deposit of: %.2f\n",amount)
		fmt.Printf("Your new balance is: %.2f",balance)
	}
}

//---------------------------------------------------------------------------------
func(x *Account) Withdraw(value float64) (bool,error) {

	var amount float64
	var balance float64
	
	//withdraw:
	fmt.Println("Please enter amount to Withdraw, divisible by 10: ")
	fmt.Scanf("%f",&amount)
	
	balance = float64(x.Balance)-amount

	if amount >= 0 && math.Mod(amount,2)==0{
		balance = float64(x.Balance)-amount
		fmt.Printf("You have made a withdrawal of %.2f\n",amount)
		fmt.Printf("Your balance is: %.2f",balance)
		return true, nil
	} else{
		fmt.Print("You have entered an incorrect value, Please try again.")
		return false, errors.New("You cannot withdraw from this account.")
		//goto withdraw
	// use the errors package to display a new, custom error message
 }

//supposed to have no parameters
// func Withdraw(x *Account) float32{
// 	var amount float64
// 	var balance float64
	
// 	withdraw:
// 	fmt.Println("Please enter amount to Withdraw in divisible by 10: ")
// 	fmt.Scanf("%f",&amount)
	
// 	balance = float64(x.Balance)-amount

// 	if amount >= 0 && math.Mod(amount,2)==0{
// 		balance = float64(x.Balance)-amount
// 		fmt.Printf("You have made a withdrawal of %.2f\n",amount)
// 		fmt.Printf("Your balance is: %.2f",balance)
// 	} else{
// 		fmt.Print("You have entered an incorrect value, Please try again.")
// 		goto withdraw
// 	}

// 	return x.Balance
 }

//---------------------------------------------------------------------------------
//Start application
func Bank(){
	min := 10
    max := 99999
	
	var time_ *time.Time
	var account_num, choice int
	currentTime := time.Now()
	time_= &currentTime
	
	//a:= new (Account
	//a := account{}
	//a := new(account)
	// a:=Account{Name: "Mark" , Balance:140,UserID:(rand.Intn(max - min) + min)}
	a:=Account{FirstName: "Mark" ,LastName: "Vine" , Balance:140,UserID:(rand.Intn(max - min) + min)}
	b:=Account{FirstName: "John" ,LastName: "Doe" , Balance:500,UserID:(rand.Intn(max - min) + min)}
	c:=Account{FirstName: "Jane" ,LastName: "Doe" , Balance:1000,UserID:(rand.Intn(max - min) + min)}
	d:=Account{FirstName: "Sarah" ,LastName: "Gill" , Balance:8500,UserID:(rand.Intn(max - min) + min)}
	
	fmt.Println(d.UserID)
	fmt.Println(time_)




	fmt.Println("Please enter your account number:")
	fmt.Scanf("%d",&account_num)

	if account_num == a.UserID{
		fmt.Println("Please select 1. Withdraw")
		fmt.Println("Please select 2. Deposit")
		fmt.Println("Please select 3. Balance")
		fmt.Scanf("%d",&choice)
		if choice == 1 { a.Withdraw(200)//Withdraw(&a)
			
		}else if choice ==2{
			Deposit(&a)
		}else if choice ==3{
			a.DisplayBalance()
		}
	}
	if account_num == b.UserID{
		fmt.Println("Please select 1. Withdraw")
		fmt.Println("Please select 2. Deposit")
		fmt.Println("Please select 3. Balance")
		fmt.Scanf("%d",&choice)
		if choice == 1 { b.Withdraw(20)//Withdraw(&b)
			
		}else if choice ==2{
			Deposit(&b)
		}else if choice ==3{
			b.DisplayBalance()
		}
	}
	if account_num == c.UserID{
		fmt.Println("Please select 1. Withdraw")
		fmt.Println("Please select 2. Deposit")
		fmt.Println("Please select 3. Balance")
		fmt.Scanf("%d",&choice)
		if choice == 1 { c.Withdraw(50)//Withdraw(&c)
			
		}else if choice ==2{
			Deposit(&c)
		}else if choice ==3{
			c.DisplayBalance()
		}
	}
	if account_num == d.UserID{
		fmt.Println("Please select 1. Withdraw")
		fmt.Println("Please select 2. Deposit")
		fmt.Println("Please select 3. Balance")
		fmt.Scanf("%d",&choice)
		if choice == 1 { d.Withdraw(400)//Withdraw(&d)
			
		}else if choice ==2{
			Deposit(&d)
		}else if choice ==3{
			d.DisplayBalance()
		}
	}

}
