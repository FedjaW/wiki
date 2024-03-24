```CSHARP
Console.WriteLine("Deep Copy Interface");

// Note: Problem with this approach is that it does not hold for a lot of classes.

var john = new Person(new[] {"John", "Smith"}, new Address("London Road", 123));

var jane = john.DeepCopy();
jane.Names[0] = "Jane"; 
jane.Address.HouseNumber = 999;

Console.WriteLine(john);
Console.WriteLine(jane);

public interface IPrototype<T>
{
    public T DeepCopy();
}

public class Person : IPrototype<Person>
{
    public string[] Names;
    public Address Address;
    
    public Person(string[] names, Address address)
    {
        Names = names;
        Address = address;
    }

    public override string ToString()
    {
        return $"{nameof(Names)}: {string.Join(" ", Names)}, {nameof(Address)}: {Address}";
    }

    public Person(Person other)
    {
        // strings sind immutable value types.
        Names = other.Names[..]; // C#8 - range operator -> anfans- und endindex -> subset array -> creates new array
        Address = new Address(other.Address);
    }

    public Person DeepCopy()
    {
        return new Person(Names[..], Address.DeepCopy());
    }
}

public class Address : IPrototype<Address>
{
    public string StreetName;
    public int HouseNumber;

    public Address(string streetName, int houseNumber)
    {
        StreetName = streetName;
        HouseNumber = houseNumber;
    }

    public Address(Address other)
    {
        StreetName = other.StreetName;
        HouseNumber = other.HouseNumber;
    }

    public override string ToString()
    {
        return $"{nameof(StreetName)}: {StreetName}, {nameof(HouseNumber)}: {HouseNumber}";
    }

    public Address DeepCopy()
    {
        return new Address(StreetName, HouseNumber);
    }
}
```