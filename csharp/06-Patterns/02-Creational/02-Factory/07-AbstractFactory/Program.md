```CSHARP
using System;
using System.Collections.Generic;

Console.WriteLine("Abstract Factory");

var machine = new HotDrinkMachine();
var drink = machine.MakeDrink(HotDrinkMachine.AvailableDrink.Tea, 100);
drink.Consume();

public interface IHotDrink
{
    void Consume();
}

internal class Tea : IHotDrink
{
    public void Consume()
    {
        Console.WriteLine("This tea is nice but I had prefer it with milk.");
    }
}

internal class Coffee : IHotDrink
{
    public void Consume()
    {
        Console.WriteLine("This coffee is sensational!.");
    }
}

public interface IHotDrinkFactory
{
    public IHotDrink Prepare(int amount);
}

internal class TeaFactory : IHotDrinkFactory
{
    public IHotDrink Prepare(int amount)
    {
        Console.WriteLine($"Put in a tea bag, boil water, pour {amount} ml, add lemon, enjoy!");
        return new Tea();
    }
}

internal class CoffeeFactory : IHotDrinkFactory
{
    public IHotDrink Prepare(int amount)
    {
        Console.WriteLine($"Grind beans, boil water, pour {amount} ml, add cream, enjoy!");
        return new Coffee();
    }
}

public class HotDrinkMachine
{
    public enum AvailableDrink
    {
        Coffee, Tea
    }

    private Dictionary<AvailableDrink, IHotDrinkFactory> factories = new();

    public HotDrinkMachine()
    {
        foreach (AvailableDrink drink in Enum.GetValues(typeof(AvailableDrink)))
        {
            var factory = (IHotDrinkFactory) Activator.CreateInstance(
                Type.GetType("07-AbstractFactory." + Enum.GetName(typeof(AvailableDrink), drink) + "Factory"));
            factories.Add(drink, factory);
        }
    }
    public IHotDrink MakeDrink(AvailableDrink drink, int amount)
    {
        return factories[drink].Prepare(amount);
    }
}
```