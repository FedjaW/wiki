```CSHARP
namespace StructuralPatterns.Decorator.CoffeeExample;

public class BaseCoffee : ICoffee
{
    public BaseCoffee()
    {
        Ingredients.Add("coffee");
        Cost = 3;
        Temperatur = 99;
    }

    public double Cost { get; set; }

    public List<string> Ingredients { get; set; }
        = new List<string>();

    public int Temperatur { get; set; }
}
```