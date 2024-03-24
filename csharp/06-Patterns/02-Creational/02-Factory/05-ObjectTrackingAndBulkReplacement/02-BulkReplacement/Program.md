```CSHARP
using System.Text;

Console.WriteLine("Bulk Replacement");

var factory1 = new ReplaceableThemeFactory();
var magicTheme = factory1.CreateTheme(true);
var magicTheme2 = factory1.CreateTheme(true);
Console.WriteLine(magicTheme.Value.BgrColor);
Console.WriteLine(magicTheme2.Value.BgrColor);

factory1.ReplaceTheme(false); // Will replace all themes to white
Console.WriteLine(magicTheme.Value.BgrColor);
Console.WriteLine(magicTheme2.Value.BgrColor);

public interface ITheme
{
    string TextColor { get; }
    string BgrColor { get; }
}

public class LightTheme : ITheme
{
    public string TextColor => "black";
    public string BgrColor => "white";
}

public class DarkTheme : ITheme
{
    public string TextColor => "white";
    public string BgrColor => "black";
}

public class ReplaceableThemeFactory
{
    private readonly List<WeakReference<Ref<ITheme>>> themes = new();
    private ITheme createThemeImpl(bool dark)
    {
        return dark ? new DarkTheme() : new LightTheme();
    }
    public Ref<ITheme> CreateTheme(bool dark)
    {
        var r = new Ref<ITheme>(createThemeImpl(dark));
        themes.Add(new(r));
        return r;
    }
    public void ReplaceTheme(bool dark)
    {
        foreach (var wr in themes)
        {
            if (wr.TryGetTarget(out var reference))
            {
                reference.Value  = createThemeImpl(dark);
            }
        }
    }
}

// A weak reference -> see: https://learn.microsoft.com/en-us/dotnet/standard/garbage-collection/weak-references
public class Ref<T> where T : class
{
    public T Value;

    public Ref(T value)
    {
        Value = value;
    }
}
```