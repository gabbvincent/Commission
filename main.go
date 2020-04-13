//Name: Vincent G.
//Date: 4/13/2020
//Descrition: Commission

package main

import "fmt"
//Create function to determine commission based on sales price
func Commission(a int)(string, int){
  if a <= 50000{ //if price is 50k or less, then 4% commission
    return "your commission is $", (a*4)/100
  } else if a >50000 && a<250000{ // if sales price is more than 50k and less than or equal to 250k, then 6% commission.
    return "your commission is $", (a*6)/100
  } else { //anyother price higher than 250k gets 6.75% commission
    return "your commission is $", (a*675/100)/100
  }
}

func main() {
  //Declare variable for a as int
  var a int
  //ask user for sales price and scan for a 
  fmt.Println("Enter in the seeling price of the item you sold")
  fmt.Scanln(&a)
  //call commission(a)
  fmt.Println(Commission(a))
}
