```CSHARP
// Read Bad Example first.

// Solution:
// Use the override and virtual keywords.

static int Area(Rectangle r) => r.Width * r.Height;

Rectangle rc = new Rectangle(2, 3);
System.Console.WriteLine($"{rc} has area {Area(rc)}");

// This works perfectly fine, and ...
Square sq = new Square();
sq.Width = 4;
System.Console.WriteLine($"{sq} has area {Area(sq)}");

// ... this one too, now.
Rectangle sq2 = new Square();
sq2.Width = 4;
System.Console.WriteLine($"{sq2} has area {Area(sq2)}");

public class Rectangle
{
    public virtual int Width { get; set; }
    public virtual int Height { get; set; }

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
    public override int Width
    {
        set { base.Width = base.Height = value; }
    }

    public override int Height
    {
        set { base.Width = base.Height = value; }
    }
}
```