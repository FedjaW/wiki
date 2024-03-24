```CSHARP
Console.WriteLine("Copy Constructor");

var john = new Person(new[] {"John", "Smith"}, new Address("London Road", 123));

var jane = new Person(john); // copy
jane.Names[0] = "Jane"; 
jane.Address.HouseNumber = 999;

Console.WriteLine(john);
Console.WriteLine(jane);

public class Person
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
}

public class Address
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
}
```