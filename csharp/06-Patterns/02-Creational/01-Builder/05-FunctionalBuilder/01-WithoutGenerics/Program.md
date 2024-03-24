```CSHARP
using System;
using System.Linq;
using System.Collections.Generic;

Console.WriteLine("Functional Builder");

var person = new PersonBuilder()
    .Called("Sarag")
    .WorksAs("Developer")
    .Build();

public class Person
{
    public string Name, Position;
}

// Sealed prevents other classes from inheriting from it.
public sealed class PersonBuilder
{
    private readonly List<Func<Person, Person>> actions
        = new List<Func<Person, Person>>();
    
    public PersonBuilder Called(string name) => Do(p => p.Name = name);

    public PersonBuilder Do(Action<Person> action) => AddAction(action);

    public Person Build() => actions.Aggregate(new Person(), (p, f) => f(p));

    private PersonBuilder AddAction(Action<Person> action)
    {
        // Taking an action and turning it into a func.
        actions.Add(p => { action(p);
            return p;
        });

        return this;
    }
}

public static class PersonBuilderExtensions
{
    public static PersonBuilder WorksAs(
        this PersonBuilder builder, 
        string position) 
        => builder.Do(p => p.Position = position);
}
```