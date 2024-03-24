# Motivation

- Want to augment an object with additinal functionality
- Do not want to rewrite or alter existing code (OCP)
- Want to keep new functionality seperate (SRP)
- Nee to be able to interact with existing structures
- Two options:
  - Inherit from required object if possible; problem: some objects are selaed
  - Solution: Build a Decorator, which simply references the decorated object(s)

## Definition

Facilitates the addition of behaviors to individual obejcts without inheriting from them.

Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors. (Source: <https://refactoring.guru/design-patterns/decorator>)

- In object-oriented programming, the decorator pattern is a design pattern that allows behavior to be added to an individual object, dynamically, without affecting the behavior of other objects from the same class.
- The decorator pattern is often useful for adhering to the Single Responsibility Principle, as it allows functionality to be divided between classes with unique areas of concern as well as to the Open-Closed Principle, by allowing the functionality of a class to be extended without being modified.
- Decorator use can be more efficient than subclassing, because an object's behavior can be augmented without defining an entirely new object. (Source: <https://en.wikipedia.org/wiki/Decorator_pattern>)
