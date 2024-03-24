```CSHARP
using System;
using System.Collections.Generic;

Console.WriteLine("Non Open Closed Principle breaking abstract dactory");
// The other implementation with an enum breaks the open closed principle.
// This approach is not violation the open closed principle.
// This approach uses reflection.

var machine = new HotDrinkMachine();
var drink = machine.MakeDrink();
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
    private List<Tuple<string, IHotDrinkFactory>> factories = new List<Tuple<string, IHotDrinkFactory>>();
    public HotDrinkMachine()
    {
        foreach (var t in typeof(HotDrinkMachine).Assembly.GetTypes())
        {
            if (typeof(IHotDrinkFactory).IsAssignableFrom(t) && !t.IsInterface)
            {
                factories.Add(Tuple.Create(
                    t.Name.Replace("Factory", string.Empty),
                    (IHotDrinkFactory)Activator.CreateInstance(t)
                ));
            }
        }
    }
    public IHotDrink MakeDrink()
    {
        System.Console.WriteLine("available drinks:");
        for (var i = 0; i < factories.Count; i++)
        {
            Tuple<string, IHotDrinkFactory> tuple = factories[i];
            System.Console.WriteLine($"{i}: {tuple.Item1}");
        }
        while (true)
        {
            string s;
            if ((s = Console.ReadLine()) != null && int.TryParse(s, out int i) && i >= 0 && i< factories.Count)
            {
                System.Console.WriteLine("Specify amount: ");
                s = Console.ReadLine();

                if (s != null && int.TryParse(s, out var amount) && amount > 0)
                {
                    return factories[i].Item2.Prepare(amount);
                }
            }
        }
    }
}
```