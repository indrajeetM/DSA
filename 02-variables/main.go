package main

import "fmt"

const Profession string = "Software engineer";

func main(){
fmt.Println("Variables");

var myName string = "Indrajeet";
fmt.Println(myName);
fmt.Printf("Name is of type %T \n",myName);

var myAge uint8 = 31;
fmt.Println(myAge);
fmt.Printf("Age is of type %T \n",myAge);


var isMarried bool = false;
fmt.Println(isMarried)
fmt.Printf("Is married is of type %T \n", isMarried);


location := "Belgaum"
fmt.Println(location);
fmt.Printf("location is of type %T \n",location);


fmt.Println(Profession);
fmt.Printf("Profession is of type %T \n",Profession);


}

