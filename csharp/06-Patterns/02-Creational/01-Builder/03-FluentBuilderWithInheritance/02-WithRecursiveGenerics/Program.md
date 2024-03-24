```CSHARP
Console.WriteLine("Code example of a fluent builder with recursive generics");

var p = Person.New.Called("Alex").WorksAsA("developer").Build();
System.Console.WriteLine(p);

public class Person
{
    public string Name;
    public string Position;

    public class Builder : PersonJobBuilder<Builder>
    {

    }

    public static Builder New => new Builder();

    public override string ToString()
    {
        return $"{nameof(Name)}: {Name}, {nameof(Position)}: {Position}";
    }
}

public abstract class PersonJobBuilder
{
    protected Person person = new Person();
    public Person Build() { return person; }
}

// Strange restriction: 
//      where SELF : PersonInfoBuilder<SELF>
// Example for the strange restriction: 
//      class Foo : Bar<Foo>
public class PersonInfoBuilder<SELF> 
    : PersonJobBuilder
    where SELF : PersonInfoBuilder<SELF>
{
    public SELF Called(string name)
    {
        person.Name = name;
        return (SELF)this;
    }
}

public class PersonJobBuilder<SELF> 
    : PersonInfoBuilder<PersonJobBuilder<SELF>>
    where SELF : PersonJobBuilder<SELF>
{
    public SELF WorksAsA(string position)
    {
        person.Position = position;
        return (SELF)this;
    }
}
```