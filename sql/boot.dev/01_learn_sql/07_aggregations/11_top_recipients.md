# Top Recipients

To better analyze user engagement on the platform, the CashPal finance team is interested in finding the users who have been the recipients of the greatest number of successful transactions.

## Challenge

Craft an SQL statement that identifies the top 2 recipients based on the number of successful transactions received. This statement should return two columns: recipient_id and transactions_received, where transactions_received is the count of successful transactions for each recipient.

The finance team has requested the following specifics:

- Only include successful transactions with a recipient.
- Group the data by recipient_id.
- Count the number of transactions received by each recipient and label this column as transactions_received.
- Sort the results so that recipients with the most transactions received are at the top.
- Limit the results to the top 2 recipients.

### Solution

```sql
SELECT recipient_id, count(*) as transactions_received
    FROM transactions
    WHERE recipient_id IS NOT NULL
    AND was_successful = true
    GROUP BY recipient_id
    ORDER BY transactions_received DESC
    LIMIT 2;
```