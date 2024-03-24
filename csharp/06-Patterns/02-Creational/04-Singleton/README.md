# Singleton

A component which is instantiated only once.

## Quote

When discussing which patterns to drop, we fount that we sill love them all. (Not really - I`m in favor of dropping Sinleton. Its use is almost always a design smell.) - Erich Gamma

## Motivation

For some components it only makes sense to have one in the system.

- Database repository
- Object factory

E.g., the constructor call is expensive

- We only do it once
- We provide everypne with the same instance

Want to prevent anyone creating additional copies

Need to take care of lazy instantiation and thread safety
