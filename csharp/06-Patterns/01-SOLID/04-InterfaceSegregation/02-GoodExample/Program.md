```CSHARP
// The Interface Segregation Principle states that
// interface should not be too large but split up into smaller ones.

var d = new Document();
d.Text = "TPS report";

var mp = new MultiFunctionPrinter();
mp.Print(d);
mp.Scan(d);
mp.Fax(d);

var op = new OldFashionedPrinter();
op.Print(d);

public class Document
{
    public string Text { get; set; }
}

// I seperated the big interface into 3 smaller
public interface IPrinter
{
    public void Print(Document d);
}

public interface IFaxer
{
    public void Fax(Document d);
}

public interface IScanner
{
    public void Scan(Document d);
}

public interface IMultiFunctionDevice : IPrinter, IScanner, IFaxer
{
}


// This Printer is just fine because it can Print, Scan and Fax.
public class MultiFunctionPrinter : IMultiFunctionDevice
{
    public void Print(Document d)
    {
        System.Console.WriteLine("print");
    }

    public void Scan(Document d)
    {
        System.Console.WriteLine("scan");
    }

    public void Fax(Document d)
    {
        System.Console.WriteLine("fax");
    }
}

public class OldFashionedPrinter : IPrinter
{
    public void Print(Document d)
    {
        System.Console.WriteLine("print");
    }
}

public class Photocopier : IPrinter, IScanner
{
    public void Print(Document d)
    {
        System.Console.WriteLine("print");
    }

    public void Scan(Document d)
    {
        System.Console.WriteLine("scan");
    }
}
```