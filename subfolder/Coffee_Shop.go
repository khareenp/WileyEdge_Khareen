//Khareen Francis Proverbs
//Coffee Shop
//June 30th 2022
/*
Write a program that will calculate the cost of a custom cup of coffee
at a gourmet coffee shop, based on the size of the cup, 
the type of coffee selected, and flavors that can be added to the 
coffee. It should complete the following steps:

1.	Ask the user what size cup they want, 
	choosing between small, medium, and large.

2.	Ask the user what kind of coffee they want, 
	choosing between brewed, espresso, and cold brew.

3.	Ask the user what flavoring they want, if any. 
	Choices include hazelnut, vanilla, and caramel.

4.	Calculate the price of the cup using the following values:
	Size:
		small: $2
		medium: $3
		large: $4
	Type:
		brewed: no additional cost
		espresso: 50 cents
		cold brew: $1
	Flavoring:
		None: no additional cost
		All other options: 50 cents

5.	Display a statement that summarizes what the user ordered.

6.	Display the total cost of the cup of coffee as well as the 
	cost with a 15% tip, in phrases that explain the values to 
	the user. Round the cost with tip to two decimal places.
		For example, if the user asks for a medium-sized 
		espresso with hazelnut flavoring, the total should be $4; 
		the total with a tip should be $4.60.


*/
package subfolder

import ("fmt"
	"strings")


// CALCULATE SMALL DRINK
func priceSmall(size,ctype,flavor string) float32{
	var total float32
	price_cup:=2.00
	//SMALL,EXPRESSO, WITH FLAVOR
	if size == "SMALL" && ctype == "EXPRESSO" && flavor !="NONE"{
		total:= price_cup + .50 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//SMALL,EXPRESSO, NO FLAVOR
		} else if size == "SMALL" && ctype == "EXPRESSO" && flavor == "NONE"{
			total:= price_cup + .50
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//SMALL,COLDBREWED, WITH FLAVOR
	if size == "SMALL" && ctype == "COLD" && flavor !="NONE"{
		total:= price_cup + 1.00 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//SMALL COLD BREWED, NO FLAVOR
		} else if size == "SMALL" && ctype == "COLD" && flavor == "NONE"{
			total:= price_cup +1
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//SMALL,BREWED, WITH FLAVOR
	if size == "SMALL" && ctype == "BREWED" && flavor !="NONE"{
		total:= price_cup + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//SMALL BREWED, NO FLAVOR
		} else if size == "SMALL" && ctype == "BREWED" && flavor == "NONE"{
			tip:= float32(price_cup * .15)
			total:= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		


	return total

}

//CALCULATE MEDIUM COFFEE
func priceMedium(size,ctype,flavor string) float32{
	var total float32
	price_cup:=3.00
	//SMALL,EXPRESSO, WITH FLAVOR
	if size == "MEDIUM" && ctype == "EXPRESSO" && flavor !="NONE"{
		total:= price_cup + .50 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//MEDIUM,EXPRESSO, NO FLAVOR
		} else if size == "MEDIUM" && ctype == "EXPRESSO" && flavor == "NONE"{
			total:= price_cup + .50
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//MEDIUM,COLDBREWED, WITH FLAVOR
	if size == "MEDIUM" && ctype == "COLD" && flavor !="NONE"{
		total:= price_cup + 1.00 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//MEDIUM COLD BREWED, NO FLAVOR
		} else if size == "MEDIUM" && ctype == "COLD" && flavor == "NONE"{
			total:= price_cup +1
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//MEDIUM,BREWED, WITH FLAVOR
	if size == "MEDIUM" && ctype == "BREWED" && flavor !="NONE"{
		total:= price_cup + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//MEDIUM BREWED, NO FLAVOR
		} else if size == "MEDIUM" && ctype == "BREWED" && flavor == "NONE"{
			tip:= float32(price_cup * .15)
			total:= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		


	return total

}

//CALCULATE LARGE COFFEE
func priceLarge(size,ctype,flavor string) float32{
	var total float32
	//LARGE,EXPRESSO, WITH FLAVOR
	if size == "LARGE" && ctype == "EXPRESSO" && flavor !="NONE"{
		price_cup:= 4.00
		total:= price_cup + .50 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//LARGE,EXPRESSO, NO FLAVOR
		} else if size == "LARGE" && ctype == "EXPRESSO" && flavor == "NONE"{
			price_cup:= 4.00
			total:= price_cup + .50
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//LARGE,COLDBREWED, WITH FLAVOR
	if size == "LARGE" && ctype == "COLD" && flavor !="NONE"{
		price_cup:= 4.00
		total:= price_cup + 1.00 + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//LARGE COLD BREWED, NO FLAVOR
		} else if size == "LARGE" && ctype == "COLD" && flavor == "NONE"{
			price_cup:= 4.00
			total:= price_cup +1
			tip:= total * .15
			total= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		//LARGE,BREWED, WITH FLAVOR
	if size == "LARGE" && ctype == "BREWED" && flavor !="NONE"{
		price_cup:= 4.00
		total:= price_cup + .50
		tip:= total * .15
		total= tip+total

		fmt.Printf("Your total is: %2f",total)

		//LARGE BREWED, NO FLAVOR
		} else if size == "LARGE" && ctype == "BREWED" && flavor == "NONE"{
			price_cup:= 4.00
			tip:= float32(price_cup * .15)
			total:= tip+total
			fmt.Printf("Your total is: %2f",total)

		}
		


	return total

}


func Coffee() {	
	var choice_cup,choice_type,flavor string 
	var mode int

	//welcoming message for coffe shop program
	coffee:
	fmt.Println("Welcome to the COFFEE SHOP. Please enter 1 to contiue or 2 to quit.")
	fmt.Scanf("%d",&mode)

	//error handling to make sure customer makes an appropriate choice
	if mode == 1{
		goto start
	}else if mode ==2{
		goto end
	} else{
		goto MESSAGE1
	}

	


	//prompt customer to enter cup size
	start:

	fmt.Println("Please select a cup of coffee.")
	fmt.Println("Your choices are: small, medium and large")
	fmt.Scanf("%s",&choice_cup)
	choice_cup= strings.ToUpper(choice_cup)


	fmt.Printf("\nYou have selected: %s\n\n",choice_cup)

	//prompt user to selet type of coffee
	fmt.Println("Please select a type of coffee.")
	fmt.Println("Your choices are: Brewed, Expresso and Cold Brew")
	fmt.Scanf("%s",&choice_type)
	choice_type= strings.ToUpper(choice_type)

	fmt.Printf("\nYou have selected: %s\n\n",choice_type)

	//prompt user to select flavour
	fmt.Println("Please select flavouring if you require:")
	fmt.Println("Your choices are: Hazelnut, Vanilla, Caramel or None:")
	fmt.Scanf("%s",&flavor)
	flavor= strings.ToUpper(flavor)

	fmt.Printf("\nYou have selected: %s\n\n",flavor)

	priceSmall(choice_cup,choice_type,flavor)
	priceMedium(choice_cup,choice_type,flavor)
	priceLarge(choice_cup,choice_type,flavor)


	
	//if user has not selected 1 or 2 to commence or quit program
	//this message will be displayed
	MESSAGE1:
	fmt.Println("\nSorry. You have made an incorrect choice. Please try again.")

	goto coffee

	//after all selections. this message will be displayed and program will terminate.
	end: 
	fmt.Println("\nThankyou have a nice day!!")
 
	

	
  

}