SELECT user_id, count(price_total) as order_count 
FROM orders 
WHERE price_total >= 1000
GROUP BY user_id
ORDER BY order_count 
DESC