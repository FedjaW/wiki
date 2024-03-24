```CSHARP
Console.WriteLine("Monostate");

// Q: If you want a singleton, why dont make the whole thing static?
// A: A static class with static members is terrible because you do not even have
//    a constructor. So you have to refering to it by name and can not use dependency injection.

// There is a variation of the singleton pattern. Is is called the Monostate Pattern.

var ceo = new CEO();
ceo.Name = "Adam Smith";
ceo.Age = 55;

var ceo2 = new CEO(); // will create a new isntance but will refere to the same data as of ceo1
Console.WriteLine(ceo2);

// There is only one CEO.
public class CEO
{
    private static string name;
    private static int age;

    public string Name
    {
      get => name;
      set => name = value;
    }

    public int Age
    {
      get => age;
      set => age = value;
    }

    public override string ToString()
    {
      return $"{nameof(Name)}: {Name}, {nameof(Age)}: {Age}";
    }
}
```