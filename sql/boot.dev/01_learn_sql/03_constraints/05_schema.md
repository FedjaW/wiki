# Schema

We've used the word schema a few times now, let's talk about what that word means. A database's schema describes how data is organized within it.

Data types, table names, field names, constraints, and the relationships between all of those entities are part of a database's schema.

## There Is No Perfect Way to Architect a Database Schema

When designing a database schema there typically isn't a "correct" solution. We do our best to choose a sane set of tables, fields, constraints, etc that will accomplish our project's goals. Like many things in programming, different schema designs come with different tradeoffs.

## How Do We Decide on a Sane Schema Architecture?

Let's use CashPal as an example. One very important decision that needs to be made is to decide which table will store a user's balance! As you can imagine, ensuring our data is accurate when dealing with money is super important. We want to be able to:

- Keep track of a user's current balance
- See the historical balance at any point in the past
- See a log of which transactions changed the balance over time

There are many ways to approach this problem. For our first attempt, let's try the simplest schema that fulfills our project's needs.

## Assignment

The architecture team at CashPal has decided on a single `transactions` table. The `transactions` table stores individual `transactions`, and we can keep track of the "current balance" on each transaction record. If we want the current balance, we can just look at the most recent transaction!

Create the `transactions` table with the following fields and constraints:

- `id` - `INTEGER` `PRIMARY KEY`
- `sender_id` - `INTEGER`
- `recipient_id` - `INTEGER`
- `memo` - `TEXT` - `NOT NULL`
- `amount` - `INTEGER` - `NOT NULL`
- `balance` - `INTEGER` - `NOT NULL`

### Solution

```sql
CREATE TABLE transactions (
  id INTEGER PRIMARY KEY,
  sender_id INTEGER,
  recipient_id INTEGER,
  memo TEXT NOT NULL,
  amount INTEGER NOT NULL,
  balance INTEGER NOT NULL
);

/* -- TEST SUITE. DON'T TOUCH BELOW THIS LINE -- */

PRAGMA TABLE_INFO('transactions');

```