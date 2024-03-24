```CSHARP
using System.Text;
using System.Collections.Generic;

Console.WriteLine("Life with a builder");

var builder = new HtmlBuilder("ul");
builder.AddChild("li", "hello").AddChild("li", "world");
System.Console.WriteLine(builder.ToString());

public class HtmlBuilder
{
    public string RootName;
    HtmlElement root = new HtmlElement();

    public HtmlBuilder(string rootName)
    {
        this.RootName = rootName;
        root.Name = rootName;
    }

    // By returning HtmlBuilder instead of void it becomes a fluent builder.
    // Meaning we can chain Methods like so: builder.AddChild("li", "hello").AddChild("li", "world");
    public HtmlBuilder AddChild(string childName, string childText)
    {
        var e = new HtmlElement(childName, childText);
        root.Elements.Add(e);
        return this;
    }

    public override string ToString()
    {
        return root.ToString();
    }

    public void Clear()
    {
        root = new HtmlElement { Name = this.RootName };
    }
}

public class HtmlElement
{
    public string Name, Text;
    public List<HtmlElement> Elements = new List<HtmlElement>();
    private const int indentSize = 2; 

    public HtmlElement() { }
    public HtmlElement(string name, string text)
    {
        this.Name = name;
        this.Text = text;
    }

    private string ToStringImpl(int indent)
    {
        var sb = new StringBuilder();
        var i = new string(' ', indentSize * indent);
        sb.Append($"{i}<{this.Name}>");

        if(!string.IsNullOrWhiteSpace(this.Text))
        {
            sb.Append(new string(' ', indentSize * (indent + 1)));
            sb.AppendLine(this.Text);
        }

        foreach (var element in this.Elements)
        {
            sb.Append(element.ToStringImpl(indent + 1));
        }
        sb.Append($"{i}</{this.Name}>");
        return sb.ToString();
    }

    public override string ToString()
    {
        return this.ToStringImpl(0);
    }
}
```