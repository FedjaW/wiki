
```CSHARP
namespace StructuralPatterns.Decorator.CoffeeExample;

public interface ICoffee
{
    double Cost { get; set; }
    List<string> Ingredients { get; set; }
    int Temperatur { get; set; }
}
```