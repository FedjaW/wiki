```CSHARP
Console.WriteLine("Code example of a fluent builder without recursive generics");

var builder = new PersonJobBuilder();
builder.Called("Alex");
    //.WorksAsA() NOTE: I would like to call .WorksAsA() on builder.Called("Alex")
    // but I can't, because Called() will return a PersonInfoBuilder and not an PersonJobBuilder.
    // PersonJobBuilder does not know anything about WorksAsA(), because it is not part of its inheritance hirachy.

    // What we need to do is replacing the return type of WorksAsA() and Called() with recursive generics.
    // To learn more about recursive generics see next example.

public class Person
{
    public string Name;
    public string Position;

    public override string ToString()
    {
        return $"{nameof(Name)}: {Name}, {nameof(Position)}: {Position}";
    }
}

// Let's say we need to adapt this builder in future. See (*)
// Because we follow SOLID we think of the open closed principle.
// Therefore we decide to use inheritance (spoiler: bad idea). 
public class PersonInfoBuilder
{
    // Protected because we will use inheritance.
    protected Person person = new Person();

    public PersonInfoBuilder Called(string name)
    {
        person.Name = name;
        return this;
    }
}

// (*)
public class PersonJobBuilder : PersonInfoBuilder
{
    protected PersonJobBuilder WorksAsA(string position)
    {
        person.Position = position;
        return this;
    }
    
}
```
