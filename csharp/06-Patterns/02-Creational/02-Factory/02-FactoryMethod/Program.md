```CSHARP
Console.WriteLine("Factory Method");

var point = Point.NewPolarPoint(1.0, Math.PI / 2);

public class Point
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

    private double x, y;

    // Notice that the constructor is private now.
    private Point(double x, double y)
    {
        this.x = x;
        this.y = y;
    }
}
```