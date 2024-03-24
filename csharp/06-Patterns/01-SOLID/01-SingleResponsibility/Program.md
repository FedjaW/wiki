```CSHARP
var j = new Journal();
j.AddEntry("I like to debug");
j.AddEntry("I kill bugs");
System.Console.WriteLine(j); // will call ToString()

var p = new Persistence();
var filename = @"c:\temp\journal.txt";
p.SaveToFile(j, filename);

public class Journal
{
    private readonly List<string> entries = new List<string>();
    private static int count = 0;

    public int AddEntry(string text)
    {
        entries.Add($"{++count}: {text}");
        return count; // memento pattern
    }

    public void RemoveEntry(int index)
    {
        entries.RemoveAt(index);
    }

    public override string ToString()
    {
        return string.Join(Environment.NewLine, entries);
    }
    
    // Violating the Single Responsibility Principle 
    // You want to extract this into an own class (Persistence) to have a separation of concerns
    public void Save(string filename)
    {
        File.WriteAllText(filename, this.ToString());
    }
}

public class Persistence
{
    public void SaveToFile(Journal j, string filename)
    {
        File.WriteAllText(filename, j.ToString());
    }
}
```
