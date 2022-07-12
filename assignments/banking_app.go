//Khareen Proverbs
//Activity 6
//Banking System

package assignments

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//---------------------------------------------------------------------------------
// Structures
type account struct{
	Type string
	//Balance map[string]float64
	TotalBalance float64
   	CheckingBalance float64
   	SavingBalance float64
   	InvestBalance float64
   	FirstName string
   	LastName string
   	UserID int
	Interest float64
   	CreatedAt time.Time
	Owner entity
	
}

type entity struct{
	ID int
	Address string
	EntityType string
}

type wallet struct{
	WalletID string
	Accounts account
	WalletOwner entity
}

//---------------------------------------------------------------------------------
//Welcome Screen
func Welcome(){
	fmt.Println("**************************************")
	fmt.Println("***             WELCOME            ***")
	fmt.Println("******* PLEASE SELECT NUMBER *********")
	fmt.Println("************   TO START   ************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
}
//---------------------------------------------------------------------------------
//Create account
func CreateAcc() {
	//acc:=map([],)
	var acc_entity, account_type string
	fmt.Println("Please enter an entity, Individual or business")
	fmt.Scanln(&acc_entity)
	fmt.Println("Please enter type of account: Checkings, Savings, Investment")
	fmt.Scanln(&account_type)

	//return acc

}

//---------------------------------------------------------------------------------
//Withdraw
func(x *account) Withdraw(value float64) (bool,error) {

	var amount float64
	var account_type string
	
	fmt.Println("Please enter account type:")
	fmt.Println("Checkings:")
	fmt.Println("Savings:")
	fmt.Println("Investment:")
	fmt.Scanf("%f",&account_type)

	withdrawal:
	fmt.Println("Please enter amount to Withdraw, divisible by 10: ")
	fmt.Scanf("%f",&amount)

	for int(amount) % 10 ==0{
		switch {

		case account_type =="checkings":
			if amount < x.CheckingBalance{
				x.CheckingBalance-= amount
				fmt.Printf("You have made a total withdrawl of: %.2f\n",amount)
				fmt.Printf("Your new balance is: %.2f",x.CheckingBalance)
			}else{
				fmt.Println("Sorry insufficient balance. Please try again")
				goto withdrawal
			}
	
		case account_type =="savings":
			if amount < x.SavingBalance{
				x.SavingBalance-=amount
				fmt.Printf("You have made a total deposit of: %.2f\n",amount)
				fmt.Printf("Your new balance is: %.2f",x.SavingBalance)
			}else{
				fmt.Println("Sorry insufficient balance. Please try again")
				goto withdrawal
			}
	
		case account_type =="investment":
			if amount < x.InvestBalance{
				x.InvestBalance-= amount
				fmt.Printf("You have made a total deposit of: %.2f\n",amount)
				fmt.Printf("Your new balance is: %.2f",x.InvestBalance)
			}else{
				fmt.Println("Sorry insufficient balance. Please try again")
				goto withdrawal
			}
		}
	} 
 return false, errors.New("You cannot withdraw that amount from this account")
}

//---------------------------------------------------------------------------------
//Deposit
func Deposit(x *account) (string, float64){
	var amount float64
	var account_type string
	
	fmt.Println("Please select account type: ")
	fmt.Scanf("%f",&account_type)

	switch {

	case account_type =="checkings":
		fmt.Println("Please enter amount to deposit: ")
		fmt.Scanf("%f",&amount)
		x.CheckingBalance+= amount
		fmt.Printf("You have made a total deposit of: %.2f\n",amount)
		fmt.Printf("Your new balance is: %.2f",x.CheckingBalance)

	case account_type =="savings":
		fmt.Println("Please enter amount to deposit: ")
		fmt.Scanf("%f",&amount)
		x.SavingBalance+=amount
		fmt.Printf("You have made a total deposit of: %.2f\n",amount)
		fmt.Printf("Your new balance is: %.2f",x.SavingBalance)

	case account_type =="investment":
		fmt.Println("Please enter amount to deposit: ")
		fmt.Scanf("%f",&amount)
		x.InvestBalance+= amount
		fmt.Printf("You have made a total deposit of: %.2f\n",amount)
		fmt.Printf("Your new balance is: %.2f",x.InvestBalance)

	}
	x.TotalBalance= x.CheckingBalance + x.SavingBalance + x.InvestBalance
	return  x.Type, x.TotalBalance
}
//---------------------------------------------------------------------------------
//Transfer between accounts
func InternalTransfer(x *account) (string, float64){
	var amount float64
	var account1,account2 string
	
	fmt.Println("Please select the account to transfer from: ")
	fmt.Scanf("%s",&account1)
	fmt.Println("Please select the account to transfer to: ")
	fmt.Scanf("%s",&account2)

	fmt.Println("Please enter the amount you wish to transfer:")
	fmt.Scanf("%f",&amount)

	switch {
		//Transfer from checkings to savings
		case account1 =="checkings" && account2 == "savings":
			x.CheckingBalance-=amount
			x.SavingBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.CheckingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.SavingBalance)

		//Transfer from checkings to investment
		case account1 =="checkings" && account2 == "investment":
			x.CheckingBalance-=amount
			x.InvestBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.CheckingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.InvestBalance)

		//Transfer from savings to checkings
		case account1 =="savings" && account2 == "checkings":
			x.SavingBalance-=amount
			x.CheckingBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.SavingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.CheckingBalance)

		//Transfer from savings to investment
		case account1 =="savings" && account2 == "investment":
			x.SavingBalance-=amount
			x.InvestBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.SavingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.InvestBalance)

		//Transfer from investment to checkings
		case account1 =="investment" && account2 == "checkings":
			x.InvestBalance-=amount
			x.CheckingBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.CheckingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.InvestBalance)

	
		//Transfer from investment to savings
		case account1 =="investment" && account2 == "savings":
			x.InvestBalance-=amount
			x.SavingBalance+=amount
			fmt.Printf("You have made a transfer of %.2f from %s to %s\n", amount,account1,account2)
			fmt.Printf("Your new balance in %s is: %2.f\n",account1, x.SavingBalance)
			fmt.Printf("Your new balance in %s is: %2.f\n",account2,x.InvestBalance)

	

	}
	x.TotalBalance= x.CheckingBalance + x.SavingBalance + x.InvestBalance
	return  x.Type, x.TotalBalance
}
//---------------------------------------------------------------------------------
//External Transfer
//Transfer between different users

// func ExternalTransfer(acc1, acc2 wallet){
// 	fmt.Println("Please ")

// }


//---------------------------------------------------------------------------------
//Display balance
func (x *account) DisplayBalance(){
	var choice int
	fmt.Println("Please select account balance you wish to display")
	fmt.Println("1. Savings")
	fmt.Println("1. Checkings")
	fmt.Println("1. Investment")
	fmt.Scanf("%s\n",&choice)

	switch {
	case choice == 1:
		fmt.Printf("Your Savings account balance is: %.2f\n", x.SavingBalance)
	case choice == 2:
		fmt.Printf("Your Checkings account balance is: %.2f\n", x.CheckingBalance)
	case choice == 3:
		fmt.Printf("Your Investment account balance is: %.2f\n", x.InvestBalance)
	}	
}

//---------------------------------------------------------------------------------
//interest added to accounts individual and business

func InterestIndividual(x *account){

	account_type := x.Type	
	interest_c:= 0.01
	interest_i:= 0.02
	interest_s:= 0.05

	for account_type == "individual"{
		switch{
		case  account_type=="checkings":
			x.CheckingBalance = (x.CheckingBalance* interest_c) + x.CheckingBalance
	
		case  account_type=="savings":
			x.CheckingBalance = (x.CheckingBalance* interest_s) + x.CheckingBalance
	
		case  account_type=="investment":
			x.CheckingBalance = (x.CheckingBalance* interest_i) + x.CheckingBalance
		}
	}
}
func InterestBusiness(x *account){

	account_type := x.Type	
	interest_c:= 0.005
	interest_i:= 0.01
	interest_s:= 0.02

	for account_type == "business"{
		switch{
		case  account_type=="checkings":
			x.CheckingBalance = (x.CheckingBalance* interest_c) + x.CheckingBalance
	
		case  account_type=="savings":
			x.CheckingBalance = (x.CheckingBalance* interest_s) + x.CheckingBalance
	
		case  account_type=="investment":
			x.CheckingBalance = (x.CheckingBalance* interest_i) + x.CheckingBalance
		}
	}
}

//---------------------------------------------------------------------------------
//Start application
func Bank(){
	Welcome()
	min := 10
    max := 99999
	
	var time_ *time.Time
	var account_num, choice int
	currentTime := time.Now()
	time_= &currentTime
	
	//a:= new (account
	//a := account{}
	//a := new(account)
	// a:=account{Name: "Mark" , Balance:140,UserID:(rand.Intn(max - min) + min)}
	a:=account{FirstName: "Mark" ,LastName: "Vine" , CheckingBalance:140.00,UserID:(rand.Intn(max - min) + min),Type: "checkings"}
	b:=account{FirstName: "John" ,LastName: "Doe" , CheckingBalance:500.00,UserID:(rand.Intn(max - min) + min),Type:"investment"}
	c:=account{FirstName: "Jane" ,LastName: "Doe" , CheckingBalance:1000.00,UserID:(rand.Intn(max - min) + min),Type: "savings"}
	d:=account{FirstName: "Sarah" ,LastName: "Gill" , CheckingBalance:8500.00,UserID:(rand.Intn(max - min) + min),Type: "savings"}
	
	fmt.Println(d.UserID)
	fmt.Println(time_)




	fmt.Println("Please enter your account number:")
	fmt.Scanf("%d",&account_num)

	if account_num == a.UserID{
		fmt.Println("Please select 1. Withdraw")
		fmt.Println("Please select 2. Deposit")
		fmt.Println("Please select 3. Balance")
		fmt.Println("Please select 4. Transfer between accounts")
		fmt.Scanf("%d",&choice)
		if choice == 1 { a.Withdraw(200)//Withdraw(&a)
			
		}else if choice ==2{
			Deposit(&a)
		}else if choice ==3{
			a.DisplayBalance()
		}else if choice ==4 {
			InternalTransfer(&a)
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
		}else if choice ==4 {
			InternalTransfer(&b)
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
		}else if choice ==4 {
			InternalTransfer(&c)
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
		}else if choice ==4 {
			InternalTransfer(&d)
		}
	}

}