```CSHARP
Console.WriteLine("Inner Factory");

var point = Point.Factory.NewPolarPoint(1.0, Math.PI / 2);

public class Point
{
    private double x, y;

    // This time we keep the constructor private!
    private Point(double x, double y)
    {
        this.x = x;
        this.y = y;
    }

    // Now, the Factory is an inner class of Point.
    public static class Factory
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
}
```