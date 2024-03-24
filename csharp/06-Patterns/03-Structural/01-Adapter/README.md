# Adapter Pattern

A construct which dapts an existing interface X to conform to the required interface Y.

## When to use it?

- Allow a system to use classes of another system that is incompatible with it.

- Allow communication between a new and already existing system which are independent of each other

- Ado.Net SqlAdapter, OracleAdapter, MySqlAdapter are the best example of Adapter Pattern.

## Who is what (in code)?

- The classes, interfaces, and objects can be identified as follows:

- ITraget - Target interface

- Employee Adapter - Adapter Class

- HR System - Adaptee Class

- ThirdPartyBillingSystem - Client
