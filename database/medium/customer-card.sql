SElECT c.id, c.email, ci.title, ci.amount from customer as c left join cart_item as ci on c.id = ci.customer_id

SELECT c.id, c.email, sum(ci.amount*ci.price) as total_cart
from customer as c inner join cart_item as ci
on ci.customer_id = c.id
group by c.id, c.email 
order by total_cart 
desc limit 10;


SELECT c.id, c.email, sum(ci.amount*ci.price) as total_cart
from customer as c inner join cart_item as ci
on ci.customer_id = c.id
WHERE c.contry = "Россия" 
group by c.id, c.email 
HAVING total_cart >= 1000
order by total_cart 
desc limit 10;
