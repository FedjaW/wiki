```CSHARP
Console.WriteLine("Faceted Builder");

var pb = new PersonBuilder();

// By writing Person instead of var the implicit converter from PersonBuilder is used.
Person person = pb
    .Lives.At("132 London Road").In("London").WithPostcode("132")
    .Works.At("Fabrikam").AsA("Engineer").Earning(123000);

System.Console.WriteLine($"{person.StreetAddress} : {person.AnnualIncome}");

public class Person
{
    // employment
    public string CompanyName, Position;
    public int AnnualIncome;

    // address
    public string StreetAddress, Postcode, City;
}

// The PersonBuilder will not build Persons itself,
// but it will hold a reference to it.
// It allows access to the sub-builders.
public class PersonBuilder // facade for other builders
{
    // Must be a reference! (struct will cause problems. Not sure why...)
    protected Person person = new Person();

    // employment
    public PersonJobBuilder Works => new PersonJobBuilder(person);

    // address
    public PersonAddressBuilder Lives => new PersonAddressBuilder(person);

    // implicit person converter
    public static implicit operator Person(PersonBuilder pb)
    {
        return pb.person;
    }
}

// The PersonJobBuilder covers the employment information about a person only.
public class PersonJobBuilder : PersonBuilder
{
    // Might not work with a value type!
    public PersonJobBuilder(Person person)
    {
        this.person = person;
    }

    public PersonJobBuilder At(string companyName)
    {
        person.CompanyName = companyName;
        return this;
    }

    public PersonJobBuilder AsA(string position)
    {
        person.Position = position;
        return this;
    }

    public PersonJobBuilder Earning(int amount)
    {
        person.AnnualIncome = amount;
        return this;
    }
}

// The PersonJobBuilder covers the address information about a person only.
public class PersonAddressBuilder : PersonBuilder
{
    // Might not work with a value type!
    public PersonAddressBuilder(Person person)
    {
        this.person = person;
    }

    public PersonAddressBuilder At(string address)
    {
        person.StreetAddress = address;
        return this;
    }

    public PersonAddressBuilder WithPostcode(string postcode)
    {
        person.Postcode = postcode;
        return this;
    }

    public PersonAddressBuilder In(string city)
    {
        person.City = city;
        return this;
    }
}
```