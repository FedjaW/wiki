# User Support Tickets

CashPal has a system in which users can submit support tickets to get help with their accounts. This data is stored in a table named `support_tickets`. You may or may not have written some bad code that led to an influx of "Account Access" related tickets. You've fixed the bug but the new influx of tickets is overshadowing other types of user concerns.

The Customer Support department wants to triage the situation by first focusing on customers who have experienced multiple issues. They've asked you to write a query that will help remove the new "Account Access" issues to filter through the noise.

## Challenge

Write an SQL statement that includes an `INNER JOIN` and returns the following:

- The user's `name`
- The user's `username`
- The count of support tickets attributed to that user labeled as `support_ticket_count`

Exclude any tickets that have `"Account Access"` as the `issue_type` and only return records for users who have more than 1 support ticket of another type. Finally, sort the records by users with the most support tickets at the top.

### Solution

```sql
```