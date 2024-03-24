```CSHARP
// Life without a factory
Console.WriteLine("Point Example");

public enum CoordinateSystem
{
    Cartesian, Polar 
}

public class Point
{
    private double x, y;

    // Pretty bad example on how to create a point in different coordinate systems.
    public Point(
        double a,
        double b,
        CoordinateSystem system = CoordinateSystem.Cartesian)
    {
        switch (system)
        {
            case CoordinateSystem.Cartesian:
                x = a;
                y = b;
                break;
            case CoordinateSystem.Polar:
                x = a * Math.Cos(b);
                y = b * Math.Sin(b);
                break;
            default:
                throw new ArgumentOutOfRangeException();
        }
    }

    // Problem:
    // I can not make a second constructor with the same arguments
    // public Point(double rho, double theta) { }
    // Therefor the CoordinateSystem was indroduced, that makes life not easier!
    // This is a bad example of a life without the factory pattern.
}
```