
```CSHARP
namespace StructuralPatterns.Decorator.CoffeeExample;

public class CappuccinoDecorator : ICoffee
{
    // I think here is the trick: 
    // This object is implementing ICoffee and will also take a reference to an ICoffee object instance.
    private readonly ICoffee _decoratedCoffee;

    // Then this instance will be manipulated aka decorated.
    public CappuccinoDecorator(ICoffee decoratedCoffee)
    {
        // Take also note that this reference to the given object is not manipulated here... 
        // but it could be done, then the given object will change its props of course to that change.
        // _decoratedCoffee.Cost = 1000; // would lead to change of cost for original coffee that is decorated as cappuccino
        _decoratedCoffee = decoratedCoffee;

        Cost = _decoratedCoffee.Cost + 0.50;

        _decoratedCoffee.Ingredients.ForEach(x => Ingredients.Add(x));

        Ingredients.Add("Espresso");

        Temperatur = _decoratedCoffee.Temperatur - 9;
    }

    public double Cost { get; set; }

    public List<string> Ingredients { get; set; }
        = new List<string>();

    public int Temperatur { get; set; }
}
```