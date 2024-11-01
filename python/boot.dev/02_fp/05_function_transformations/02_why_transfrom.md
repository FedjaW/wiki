# Why Transform?
You might be wondering:

- "When would I use function transformations in the real world?"
- "Isn't it simpler to just define functions at the top level of the code, and call them as needed?"

Good questions. To be clear, we don't just transform functions at runtime for the fun of it! We only use advanced techniques like function transformations when they make our code simpler than it would otherwise be.

## Code Reusability

Creating variations of the same function dynamically can make it a lot easier to share common functionality. Take a look at this formatter function. It accepts a "pattern" and returns a new function that formats text according to that pattern:

```py
def formatter(pattern):
    def inner_func(text):
        result = ""
        i = 0
        while i < len(pattern):
            if pattern[i:i+2] == '{}':
                result += text
                i += 2
            else:
                result += pattern[i]
                i += 1
        return result
    return inner_func
```

Now we can create new formatters easily:

```py
bold_formatter = formatter("**{}**")
italic_formatter = formatter("*{}*")
bullet_point_formatter = formatter("* {}")
```
And use them like this:

```py
print(bold_formatter("Hello"))
# **Hello**
print(italic_formatter("Hello"))
# *Hello*
print(bullet_point_formatter("Hello"))
# * Hello
```