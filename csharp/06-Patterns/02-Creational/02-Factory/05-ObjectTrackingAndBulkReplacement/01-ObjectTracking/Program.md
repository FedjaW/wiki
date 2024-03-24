```CSHARP
using System.Text;

Console.WriteLine("Object Tracking");

var factory = new TrackingThemeFactory();
var theme1 = factory.CreateTheme(false);
var theme2 = factory.CreateTheme(true);

Console.WriteLine(factory.Info);

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

public class TrackingThemeFactory
{
    private readonly List<WeakReference<ITheme>> themes = new();

    public ITheme CreateTheme(bool dark)
    {
        ITheme theme = dark ? new DarkTheme() : new LightTheme();
        themes.Add(new WeakReference<ITheme>(theme));
        return theme;
    }

    public string Info 
    {
        get 
        {
            var sb = new StringBuilder();
            foreach (var reference in themes)
            {
                if (reference.TryGetTarget(out var theme))
                {
                    bool dark = theme is DarkTheme;
                    sb.Append(dark? "Dark" : "Light")
                      .AppendLine(" theme");
                }
            }

            return sb.ToString();
        }
    }
}
```