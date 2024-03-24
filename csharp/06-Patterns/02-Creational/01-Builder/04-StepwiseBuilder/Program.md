```CSHARP
Console.WriteLine("Stepwise Builder");

// This is the final usage of a stepwise builder
var car = CarBuilder.Create()   // ISpecifyCarType
    .OfType(CarType.Crossover)  // ISpecifyWheelSize
    .WithWheels(18)             // IBuildCar
    .Build();                   // Car

public enum CarType
{
    Sedan, // Wheelsize of  15 - 17 inches
    Crossover // Wheelsize of 17 - 20 inches

    // -> Builder has to follow an order.
    //    First it has to select a CarType and then the WheelSize.
    // Problem: How to create a builder that creates firt the type and then the wheelsize?
    // Answere: We will use the interface segregation principle (I in SOLID)
    // We will return one interface for configuring the CarType.
    // We will return another interface for configuring the Wheelsize.
    // We will return a third interface for building the Car.
}

public class Car
{
    public CarType Type;
    public int WheelSize;
}

// The trick here is that one interface follows "stepwise" another interface.
// ISpecifyCarType -> ISpecifyWheelSize -> IBuildCar -> Car
public interface ISpecifyCarType
{
    ISpecifyWheelSize OfType(CarType type);
}

public interface ISpecifyWheelSize
{
    IBuildCar WithWheels(int size);
}

public interface IBuildCar
{
    public Car Build();
}

public class CarBuilder
{
    // Instead of a private Impl Class inside the CarBuilder class
    // you can put the builder inside the Car class.
    // The goal with this Impl class is that is is private an never exposed.
    private class Impl:
        ISpecifyCarType,
        ISpecifyWheelSize,
        IBuildCar
    {
        private Car car  = new Car();

        public ISpecifyWheelSize OfType(CarType type)
        {
            car.Type = type;
            return this;
        }

        public IBuildCar WithWheels(int size)
        {
            switch (car.Type)
            {
                case CarType.Crossover when size < 17 || size > 20:
                case CarType.Sedan when size < 15 || size > 17:
                    throw new System.ArgumentException("Wrong size of wheel");
            }

            car.WheelSize = size;
            return this;
        }

        public Car Build()
        {
            return car;
        }
    }

    public static ISpecifyCarType Create()
    {
        return new Impl();
    }
}
```