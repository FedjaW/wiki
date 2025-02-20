# Issue Diversity

The Customer Support team is trying to figure out which countries are experiencing the widest range of issues with the application. They are interested in the number of unique issue types submitted for each country.

## Challenge

Write an SQL statement using multiple joins to return the following:

- The country's `name`
- The country's `country_code`
- The number of unique `issue_types` per country name labeled as `issue_diversity`

Only include records for countries that have support tickets linking back to a user in that country. Sort the records by highest issue diversity first, then by country name in ascending order.

### Solution

```sql
SELECT
    countries.name,
    countries.country_code,
    COUNT(DISTINCT support_tickets.issue_type) AS issue_diversity
FROM countries
INNER JOIN users ON countries.country_code = users.country_code
INNER JOIN support_tickets ON users.id = support_tickets.user_id
GROUP BY countries.name
ORDER BY issue_diversity DESC, countries.name ASC;
```