# User Count

CashPal has a dashboard on its website that shows statistics on where users are located. One of your QA team members is concerned because the number of users located in the US looks pretty small.

## Challenge

Write an SQL query that returns the count of every user record that has their `country_code` equal to `US`.

Remember, we want to know the number of total records so we can use the wildcard `(*)` in our `COUNT(*)`. While you could also use the `id` in your `COUNT(id)`, for this challenge stick with the `(*)` wildcard.

### Solution

```sql
SELECT COUNT(*) FROM users WHERE country_code = 'US';
```

# Update Country Codes

Using our previous query, we've noticed that some of the records were incorrectly saved with a country_code value of USA instead of US.

## Challenge

Write an SQL statement to update the country_code value from 'USA' to 'US' for all user records, ensuring that only those records initially marked with 'USA' are changed.

### Solution

```sql
UPDATE users
SET country_code = 'US'
WHERE country_code = 'USA';
```