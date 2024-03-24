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
// op.Scan(d); Throws exception
// op.Fax(d); Throws exception

public class Document
{
    public string Text { get; set; }
}

public interface IMachine
{
    public void Print(Document d);
    public void Scan(Document d);
    public void Fax(Document d);
}

// This Printer is just fine because it can Print, Scan and Fax.
public class MultiFunctionPrinter : IMachine
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

// This Printer can print only, it can not Scan or Fax.
// So what to do, because we have to implement all the interface methods.
public class OldFashionedPrinter : IMachine
{
    public void Print(Document d)
    {
        System.Console.WriteLine("print");
    }

    public void Scan(Document d)
    {
        // This is bad. That is why the interface segregation principle come into play
        System.Console.WriteLine("I can not scan");
        throw new NotImplementedException();
    }

    public void Fax(Document d)
    {
        // This is bad. That is why the interface segregation principle come into play
        System.Console.WriteLine("I can not fax");
        throw new NotImplementedException();
    }
}
```