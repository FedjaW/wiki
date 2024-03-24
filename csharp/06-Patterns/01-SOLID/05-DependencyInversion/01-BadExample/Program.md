```CSHARP
// The Dependency Inversion Principle states that
// high level components should not depend directly on low-level components
// instead they should depend on an abstraction.

Console.WriteLine("Hello, World!");

public enum Relationship
{
    Parent, Child, Sibling
}

public class Person
{
    public string Name;
}

// low-level
public class Relationships
{
    // From-Person -> Relationship -> To-Person
    private List<(Person, Relationship, Person)> relations 
        = new List<(Person, Relationship, Person)>();

    public void AddParentAndChild(Person parent, Person child)
    {
        relations.Add((parent, Relationship.Parent, child));
        relations.Add((child, Relationship.Child, parent));
    }

    public List<(Person, Relationship, Person)> Relations => relations;
}

public class Research
{
    public Research(Relationships relationships)
    {
        var relations = relationships.Relations;
        foreach (var r in relations.Where(x => x.Item1.Name == "John" && x.Item2 ==  Relationship.Parent))
        {
            System.Console.WriteLine($"John has a child called {r.Item3.Name}");
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