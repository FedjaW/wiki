```CSHARP
using System.IO;
using System.Xml.Serialization;

Console.WriteLine("Copy Through Serialization");

// This approach shown here is in fact the way how the prototype pattern is implemented in the real world.

Foo foo = new Foo { Stuff = 42, Whatever = "abc" };

Foo foo2 = foo.DeepCopyXml();

foo2.Whatever = "xyz";
Console.WriteLine(foo);
Console.WriteLine(foo2);

public static class ExtensionMethods
{
    public static T DeepCopyXml<T>(this T self)
    {
        using (var ms = new MemoryStream())
        {
        XmlSerializer s = new XmlSerializer(typeof(T));
        s.Serialize(ms, self);
        ms.Position = 0;
        return (T) s.Deserialize(ms);
        }
    }
}

public class Foo
{
    public uint Stuff;
    public string Whatever;

    public override string ToString()
    {
        return $"{nameof(Stuff)}: {Stuff}, {nameof(Whatever)}: {Whatever}";
    }
}
```