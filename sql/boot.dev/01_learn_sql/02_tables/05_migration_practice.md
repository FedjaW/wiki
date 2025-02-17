# Migration Practice

When writing reversible migrations, we use the terms "up" and "down" migrations. An "up" migration is simply the set of changes you want to make, like altering/removing/adding/editing a table in some way. A "down" migration includes the changes that would revert any of the "up" migration's changes.

## Assignment

Add additional columns to the `transactions` table. We want to know whether or not the transaction was successfully completed between two users. We also want our database to track the type of transaction.

Our `transactions` table looks like this at the moment:

![img](05.png)

Complete the following up migration:

- Add the BOOLEAN was_successful column to the transactions table.
- Add the TEXT transaction_type column to the transactions table.

Note: `BOOL` is valid, but the assignment expects BOOLEAN, so use `BOOLEAN` instead of `BOOL` to pass this assignment.

### Solution

```sql
ALTER TABLE transactions 
ADD COLUMN was_successful BOOLEAN;

ALTER TABLE transactions 
ADD COLUMN transaction_type TEXT;

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('transactions');
```
