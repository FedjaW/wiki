# Decorators

Python decorators are just syntactic sugar for higher-order functions.

Example:

```py
def vowel_counter(func_to_decorate):
    vowel_count = 0
    def wrapper(doc):
        nonlocal vowel_count
        vowels = "aeiou"
        for char in doc:
            if char in vowels:
                vowel_count += 1
        print(f"Vowel count: {vowel_count}")
        return func_to_decorate(doc)
    return wrapper

@vowel_counter
def process_doc(doc):
    print(f"Document: {doc}")

process_doc("What")
# Vowel count: 1
# Document: What

process_doc("a wonderful")
# Vowel count: 5
# Document: a wonderful

process_doc("world")
# Vowel count: 6
# Document: world
```

The @vowel_counter line is "decorating" the process_doc function with the vowel_counter function. vowel_counter is called once when process_doc is defined with the @ syntax, but the wrapper function that it returns is called every time process_doc is called. That's why vowel_count is preserved and printed after each time.

## With decorator

```py
@vowel_counter
def process_doc(doc):
    print(f"Document: {doc}")

process_doc("Something wicked this way comes")
```

## Without decorator

```py
def process(doc):
    print(f"Document: {doc}")

process_doc = vowel_counter(process)
process_doc("Something wicked this way comes")
```
