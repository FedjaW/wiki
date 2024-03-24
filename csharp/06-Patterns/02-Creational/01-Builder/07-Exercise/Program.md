```CSHARP
Console.WriteLine("Coding Builder Exercise");

var cb = new CodeBuilder("Person").AddField("Name", "string").AddField("Age", "int");
Console.WriteLine(cb);

public class CodeBuilder
{
    private string ClassName;
    private List<(string, string)> Properties = new List<(string, string)>();
    
    public CodeBuilder(string className)
    {
        className = className;
    }
    
    public CodeBuilder AddField(string attribute, string type)
    {
        Properties.Add((type, attribute));
        return this;
    }

    // TODO: override .toString()
}
```