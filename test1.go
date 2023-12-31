package main

import "fmt"

func main() {
     fmt.Println("\nHello go language")
     fmt.Println("Hello World\n")
     
  var x float64 = 20.0
  y := "a"
  z := 30 
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)
   fmt.Println(y)
   fmt.Printf("y is of type %T\n", y)
   fmt.Println(z)
   fmt.Printf("z is of type %T\n", z)
   
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("\nvalue of area : %d", area)
   
   // if statment example

   var b int = 10
 
   /* check the boolean condition using if statement */
   if( b < 20 ) {
      /* if condition is true then print the following */
      fmt.Printf("\n\nb is less than 20\n" )
   }
   fmt.Printf("value of b is : %d\n", b)
	
// if... else statment example

 /* local variable definition */
   var a int = 100;
 
   /* check the boolean condition */
   if( a < 20 ) {
      /* if condition is true then print the following */
      fmt.Printf("\n\na is less than 20\n" );
   } else {
      /* if condition is false then print the following */
      fmt.Printf("\n\na is not less than 20\n" );
   }
   fmt.Printf("value of a is : %d\n", a);

// nested if statement example

/* local variable definition */
   var c int = 100
   var d int = 200
 
   /* check the boolean condition */
   if( c == 100 ) {
      /* if condition is true then check the following */
      if( d == 200 )  {
         /* if condition is true then print the following */
         fmt.Printf("\n\nValue of c is 100 and d is 200\n" );
      }
   }
   fmt.Printf("Exact value of c is : %d\n", c );
   fmt.Printf("Exact value of d is : %d\n", d );
}