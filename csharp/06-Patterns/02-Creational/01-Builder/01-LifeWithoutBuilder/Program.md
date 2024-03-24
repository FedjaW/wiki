```CSHARP
using System.Text;

Console.WriteLine("Life without a builder");
// Forget the fact that we're using a StringBuilder.
// The focus is that we build up an HTML text without any special builder pattern.
// The next examples will show how to do better by using a builder pattern to create HTML text.

var hello = "hello";
var sb = new StringBuilder();
sb.Append("<p>");
sb.Append(hello);
sb.Append("</p>");
System.Console.WriteLine(sb);

var words = new[] { "hello", "world" };
sb.Clear();
sb.Append("<ul>");
foreach (var word in words)
{
    sb.AppendFormat("<li>{0}</li>", word);
}
sb.Append("</ul>");
System.Console.WriteLine(sb);
```
