# Sum Types

A "sum" type is the opposite of a "product" type. This Python object is an example of a product type:

```py
man.studies_finance = True
man.has_trust_fund = False
```

As opposed to product types, which can have many (often infinite) combinations, sum types have a fixed number of possible values. To be clear: Python doesn't really support sum types. We have to use a workaround and invent our own little system and enforce it ourselves.
