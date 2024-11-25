# Iterating

Even though iterating with linked lists kinda sucks compared to the simplicity of arrays (normal lists), we've got to do it. Although the implementation is more complex and slow, we can still make it easy for users of our class by implementing the `__iter__` method.

## The yield keyword

The `yield` keyword in Python returns a value, kind of like `return`. However, it's used to turn the function into a generator function.

A generator function creates a new function object. When that function is called, it executes the code in the generator function until it hits a `yield` statement. At that point, the function pauses and returns the value of the `yield` statement. The next time the function is called, it picks up right where it left off.

```py
def create_message_generator():
    yield "hi"
    yield "there"
    yield "friend"

gen = create_message_generator()
first = next(gen)
print(first)  # prints: hi
second = next(gen)
print(second)  # prints: there
third = next(gen)
print(third)  # prints: friend
```

Every time you call `create_message_generator()`, it creates a new `generator` instance. To continue from where you left off, you need to assign this generator to a variable (like gen in the example above). This way, when you use `next()` or loop over the generator, you’re continuing with the same instance rather than starting a new one.