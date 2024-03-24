```CSHARP
// Open-Closed-Principle means that a class should be open 
// for extension but closed for modification

// How can we achive this? The answere is inheritance and interfaces.
// We will implement the Specification Pattern. See -> (#)
// But first read the bad example

var apple = new Product("Apple", Color.Green, Size.Small);
var tree = new Product("Tree", Color.Green, Size.Large);
var house = new Product("House", Color.Blue, Size.Large);

Product[] products = { apple, tree, house };

var pf = new ProductFilter();
System.Console.WriteLine("Green products (old):"); 
// we have only filtering by size implemented
// in order to filter by color we have to modify implementation
// this is breaking the open-closed-principle!
// see -> (*)
foreach (var p in pf.FilterByColor(products, Color.Green))
{
    System.Console.WriteLine($" -  {p.Name} is green");
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

// Bad Example starts here
public class ProductFilter
{
    // Step 1
    public IEnumerable<Product> FilterBySize(IEnumerable<Product> products, Size size)
    {
        foreach (var p in products)
        {
            if (p.Size == size)
             {
                yield return p;
            }
        }
    }

    // Step 2
    // -> (*)
    public IEnumerable<Product> FilterByColor(IEnumerable<Product> products, Color color)
    {
        foreach (var p in products)
        {
            if (p.Color == color)
             {
                yield return p;
            }
        }
    }
}
```