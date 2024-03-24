```CSHARP
Console.WriteLine("The Singleton Pattern");

var db = SingletonDatabase.Instance;
string city = "Osaka";
Console.WriteLine($"{city} has population {db.GetPopulation(city)}");

public interface IDatabase
{
    public int GetPopulation(string name);
}

public class SingletonDatabase : IDatabase
{
    private Dictionary<string, int> capitals = new();

    // Constructor is private so no one can create new databases. (*)
    private SingletonDatabase()
    {
        System.Console.WriteLine("init database");
        
        var rows = File.ReadAllLines("capitals.txt");

        // Fill the dictionary with city-name and population
        for (int i = 0; i < rows.Count() - 1; i = i + 2)
        {   
            var city = rows[i];
            var population = rows[i + 1];
            capitals.Add(city, int.Parse(population));
        }
    }

    public int GetPopulation(string name)
    {
        return capitals[name];
    }

    // (*) Here we take care of returning always the same instance of the database.
    // Note: Lazy will only instanciate the database if the instace.Value is called.
    private static Lazy<SingletonDatabase> instance = 
        new Lazy<SingletonDatabase>(() => new SingletonDatabase());
    public static SingletonDatabase Instance = instance.Value;
}
```