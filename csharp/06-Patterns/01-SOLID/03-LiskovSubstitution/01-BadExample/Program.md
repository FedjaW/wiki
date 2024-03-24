```CSHARP
// The Liskov Substitution Principle states that
// it should always be possible to "upcast" to a base-type
// and the operations should always work as expected.
// Meaning the Square should behave as a Square, even is you are getting
// a reference to a Reactangle.

static int Area(Rectangle r) => r.Width * r.Height;

Rectangle rc = new Rectangle(2, 3);
System.Console.WriteLine($"{rc} has area {Area(rc)}");

// This works perfectly fine, but ...
Square sq = new Square();
sq.Width = 4;
System.Console.WriteLine($"{sq} has area {Area(sq)}");

// ... if I change Square to Reactangle, (that is doable, because Square inherits from Reactangle)
// it will not calculate the correct result.
// This happens, because if you are setting the width you are ONLY setting the width.
Rectangle sq2 = new Square();
sq2.Width = 4;
System.Console.WriteLine($"{sq2} has area {Area(sq2)}");

public class Rectangle
{
    public int Width { get; set; }
    public int Height { get; set; }

    public Rectangle() {}

    public Rectangle(int width, int height)
    {
        this.Width = width;
        this.Height = height;
    }

    public override string ToString()
    {
        return $"{nameof(this.Width)}: {this.Width}, {nameof(this.Height)}: {this.Height}";
    }
}

public class Square : Rectangle
{
    // Sidenote: Use the new keyword to hide the inherited member
    // see: https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/using-properties#hidden-property-example
    public new int Width
    {
        set { base.Width = base.Height = value; }
    }

    public new int Height
    {
        set { base.Width = base.Height = value; }
    }
}
```