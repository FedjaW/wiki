```CSHARP
Console.WriteLine("ICloneable is bad!");
// Why is ICloneable bad?
// The reason is that ICloneable do not force you to clone correctly
// What is meant by "correctly". The example below shows that it possible to shallow clone an object 
// which leads to changing the referenced obejct instead of a new objects values.

var john = new Person(new[] {"John", "Smith"}, new Address("London Road", 123));

// This will not work of course, because we copied the reference to john
// and will change the name of the person and not clone the object.
// So one approach is to implement the ICloneable.
var jane = john;
jane.Names[0] = "Jane"; // changed the referenced objects name from john to jane.

Console.WriteLine(jane);
Console.WriteLine(john);

// Here we make use of the ICloneable.
var jessy = (Person)john.Clone();
jessy.Address.HouseNumber = 999;

Console.WriteLine(jane);
Console.WriteLine(john);
Console.WriteLine(jessy); // change the referenced objects address from 123 to 999.

// Note: Implementing the ICloneable for class Address (*) so that the address is deep copied then it will work,
// but you can see the struggle.


public class Person : ICloneable
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

    // Q: How to we clone?
    // A shallow copy, meaning references only
    // or a deep copy? 
    // -> This is the issue with ICloneable, 
    // you can implement a shallow copy which is not solving the problem and leads to errors.
    public object Clone()
    {
        // This is a shallow copy!
        // Won't solve the problem.
        return new Person(Names, Address);
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

    public override string ToString()
    {
        return $"{nameof(StreetName)}: {StreetName}, {nameof(HouseNumber)}: {HouseNumber}";
    }
    // (*)
    // public object Clone()
    // {
    //    return new Address(StreetName, HouseNumber);
    // }
}
```