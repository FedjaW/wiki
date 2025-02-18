# Invoice Subquery

Certain customers have been using their personal CashPal accounts for business expenses. CashPal is trying to contact these customers so they can upsell business accounts.

## Challenge

Using a subquery, write an SQL statement that retrieves full user records for every user who matches the sender_id in a transaction with invoice or tax mentioned anywhere in the transaction note, and who is not an admin.

### Solution

```sql
SELECT * FROM users
  WHERE id IN (
    SELECT sender_id
    FROM transactions
    WHERE note LIKE '%invoice%'
    OR note LIKE '%tax%'
  )
  AND is_admin = false;
```