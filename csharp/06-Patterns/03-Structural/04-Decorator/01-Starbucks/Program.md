```CSHARP
using StructuralPatterns.Decorator.CoffeeExample;

Console.WriteLine("Decorator Pattern");

var coffee = new BaseCoffee();
System.Console.WriteLine("coffee: " + " cost: " + coffee.Cost.ToString() + " temp: " + coffee.Temperatur.ToString());

var cappuccino = new CappuccinoDecorator(coffee);
System.Console.WriteLine("cappuccino: " + " cost: " + cappuccino.Cost.ToString() + " temp: " + cappuccino.Temperatur.ToString());

var doubleCappuccino = new CappuccinoDecorator(cappuccino);
System.Console.WriteLine("double cappuccino: " + " cost: " + doubleCappuccino.Cost.ToString() + " temp: " + doubleCappuccino.Temperatur.ToString());

// Take note that the original base coffe was not changed due to some decorators. (nice!)
System.Console.WriteLine("coffee: " + " cost: " + coffee.Cost.ToString() + " temp: " + coffee.Temperatur.ToString());
```