# Transaction Insights

CashPal is putting together its end-of-year report for the executive staff. They are particularly interested in the largest transactions that took place over the year. Your task is to create a query that retrieves data on the 5 largest successful transactions.

## Challenge

Using the setup code - write a query that returns the recipient's ID along with the transaction amount and note. Return this information for the top 5 largest successful transactions ordered from highest to lowest.

### Solution

```sql
SELECT recipient_id, amount, note
FROM transactions
WHERE was_successful = true
ORDER BY amount DESC
LIMIT 5;
```