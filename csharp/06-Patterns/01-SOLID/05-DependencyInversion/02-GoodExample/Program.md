```CSHARP
using System.Linq;

Console.WriteLine("Hello, World!");

public enum Relationship
{
    Parent, Child, Sibling
}

public class Person
{
    public string Name;
}

public interface IRelationshipBrowser
{
    public IEnumerable<Person> FindAllChildrenOf(string name);
}

// low-level
public class Relationships : IRelationshipBrowser
{
    // From-Person -> Relationship -> To-Person
    private List<(Person, Relationship, Person)> relations 
        = new List<(Person, Relationship, Person)>();

    public void AddParentAndChild(Person parent, Person child)
    {
        relations.Add((parent, Relationship.Parent, child));
        relations.Add((child, Relationship.Child, parent));
    }

    public IEnumerable<Person> FindAllChildrenOf(string name)
    {
        return relations.Where(x => x.Item1.Name == name && x.Item2 == Relationship.Parent).Select(r => r.Item3);
    }
}

public class Research
{
    public Research(IRelationshipBrowser browser)
    {
        foreach (var p in browser.FindAllChildrenOf("John"))
        {
            System.Console.WriteLine($"John has a child called {p.Name}");
        }
    }

    static void Main(string[] args)
    {
        var parent = new Person { Name = "John" };
        var child1 = new Person { Name = "John" };
        var child2 = new Person { Name = "John" };

        var relationships = new Relationships();
        relationships.AddParentAndChild(parent, child1);
        relationships.AddParentAndChild(parent, child2);

        new Research(relationships);
    }
}
```