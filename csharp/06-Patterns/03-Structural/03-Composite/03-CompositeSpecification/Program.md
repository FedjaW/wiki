```CSHARP
Console.WriteLine("Composite Specification");

// Note: This code might need rework. Did not check. Serves as an rough example.

// Important Note: Most of this code is copied from the GoodExample of the Open-Closed-Principle (see folder).
// The new code representing the Composite Specification is marked.
// The Composite Specification is an generalization of the specifiation.

var apple = new Product("Apple", Color.Green, Size.Small);
var tree = new Product("Tree", Color.Green, Size.Large);
var house = new Product("House", Color.Blue, Size.Large);

Product[] products = { apple, tree, house };

var bf = new BetterFilter();
System.Console.WriteLine("Green products:"); 
foreach (var p in bf.Filter(products, new ColorSpecification(Color.Green)))
{
    System.Console.WriteLine($" -  {p.Name} is green");
}

System.Console.WriteLine("Large blue products:"); 
foreach (var p in bf.Filter(
    products, 
    new AndSpecification<Product>(
        new ColorSpecification(Color.Blue), 
        new SizeSpecification(Size.Large))))
{
    System.Console.WriteLine($" -  {p.Name} is large and blue");
}

public enum Color
{
    Red, Green, Blue
}

public enum Size 
{
    Small, Medium, Large, Yuge
}

public class Product
{
    // Sidenote: public fields are pretty bad practice!
    public string Name;
    public Color Color;
    public Size Size;

    public Product(string name, Color color, Size size)
    {
        this.Name = name;
        this.Color = color;
        this.Size = size;
    }
}

public interface ISpecification<T>
{
    bool IsSatisfied(T t);
}

public interface IFilter<T>
{
    IEnumerable<T> Filter(IEnumerable<T> items, ISpecification<T> spec);
}

public class ColorSpecification : ISpecification<Product>
{
    private Color Color;

    public ColorSpecification(Color color)
    {
        this.Color = color;    
    }

    public bool IsSatisfied(Product p)
    {
        return p.Color == this.Color;
    }
}

public class SizeSpecification : ISpecification<Product>
{
    private Size Size;

    public SizeSpecification(Size size)
    {
        this.Size = size;    
    }

    public bool IsSatisfied(Product p)
    {
        return p.Size == this.Size;
    }
}

// (*) added a new base class
public abstract class CompositeSpecification<T> : ISpecification<T>
{
    protected readonly ISpecification<T>[] items;
    public CompositeSpecification(params ISpecification<T>[] items)
    {
        this.items = items;
    }
}

// (*) This is really a combinator of specifications
public class AndSpecification<T> : CompositeSpecification<T>
{
    public AndSpecification(params ISpecification<T>[] items) : base(items)
    {
    }

    public bool IsSatisfied(T t)
    {
        // Any instead of All -> OrSpecification
        return items.All(i => i.IsSatisfied(t));
    }
}

public class BetterFilter : IFilter<Product>
{
    public IEnumerable<Product> Filter(IEnumerable<Product> products, ISpecification<Product> colorSpec)
    {
        foreach (var p in products)
        {
            if (colorSpec.IsSatisfied(p))
            {
                yield return p;
            }
        }
    }
}
```