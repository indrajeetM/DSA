# DSA
this is the DSA syllubus for go lang


## Beginners guide video
1. [Hitesh Choudhary](https://www.youtube.com/watch?v=JoUSa8jtadE&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=13)


## Printing
1. Println : Prints to new line
2. Printf : Prints the formatted string

## Print formated values
1. %T : To print the type
    ```
        var myName string = "john"
        fmt.Printf("Data type is %T\n:",myName)
    ```

2. %v : To print the value
3. %+v : To print the detailed view
    ```
        johnDoe := User{"John Doe", 22, "jd@gmail.com", "Software engineer"}

        //below line prints the values without the keys
        fmt.Println(johnDoe)

        //below line prints the value with the keys
        fmt.Printf("The details are %+v\n", johnDoe)
    ```
