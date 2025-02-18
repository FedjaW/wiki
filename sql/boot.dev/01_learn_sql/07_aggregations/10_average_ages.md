# Average Ages

The CashPal marketing team was able to optimize their advertising with the help of the user data we pulled. However, they now want you to query the average age of users for each country that the business operates in, not just the United States.

## Challenge

Write an SQL statement that returns two columns, the `country_code` and the average age of users for records with that `country_code`. The marketing team has asked that we round the average to the nearest whole number and rename the column that contains the average age to `average_age`.

### Solution

```sql
SELECT country_code, round(avg(age)) as average_age
  FROM users
  GROUP BY country_code;
```