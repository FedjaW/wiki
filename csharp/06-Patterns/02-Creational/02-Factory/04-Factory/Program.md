```CSHARP
Console.WriteLine("The real factory");

// Instead of having factory methods we will create a factory class.
// This is fulfilling the s in solid (seperation of concerns).
// The big drawback of this approach ist that the Point constructor is public.
// Solution to this problem, see InnerFactory

var point = PointFactory.NewPolarPoint(1.0, Math.PI / 2);

public static class PointFactory
{
    // Factory method 1
    public static Point NewCartesianPoint(double x, double y)
    {
        return new Point(x, y);
    }

    // Factory method 2
    public static Point NewPolarPoint(double rho, double theta)
    {
        return new Point(rho * Math.Cos(theta), rho* Math.Sin(theta));
    }
}

public class Point
{
    private double x, y;

    // Notice that the constructor is public now. 
    // Before it was private, see factoy method examples.
    // Maybe this is not what you want. Will tackle this later
    // For now this is the frog we eat.
    public Point(double x, double y)
    {
        this.x = x;
        this.y = y;
    }
}
```