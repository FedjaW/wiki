# Inheritance

The holy grail of object-oriented programming: inheritance. Non-OOP languages like Go and Rust allow for encapsulation and abstraction features as nearly every language does. Inheritance, on the other hand, tends to be unique to class-based languages like Python, Java, and Ruby.

## What is inheritance?
Inheritance allows one class, the "child" class, to inherit the properties and methods of another class, the "parent" class.
This powerful language feature helps us avoid writing a lot of the same code twice. It allows us to DRY (don't repeat yourself) up our code


Here Cow is a "child" class that inherits from the "parent" class Animal:

```python
class Animal:
    # parent "Animal" class

class Cow(Animal):
    # child class "Cow" inherits "Animal"
```

The Cow class can reuse the Animal class's constructor with the super() method:

```python
class Animal:
    def __init__(self, num_legs):
        self.num_legs = num_legs

class Cow(Animal):
    def __init__(self):
        # call the parent constructor to
        # give the cow some legs
        super().__init__(4)
```

Very nice implementation of a Square and a Rectangular:

```python
class Rectangle:
    def __init__(self, length, width):
        self.__length = length
        self.__width = width

    def get_area(self):
        return self.__length * self.__width

    def get_perimeter(self):
        return 2 * (self.__length + self.__width)


class Square(Rectangle):
    def __init__(self, length):
        super().__init__(length, length)
```