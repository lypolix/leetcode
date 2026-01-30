SELECT COUNT(DISTINCT c.ID_CUSTOMER)
from customer c
left join purchases p on c.ID_CUSTOMER = p.user_id
where p.creates_at >= '2021-10-10' and p.sru_id = 5



SELECT calls.start_dttm, customers.first_nm, customers.last_nm, customers.middle_nm
from calls
left join customers
    on calls.customer_id = customers.customer_id
where dozv_flag=0 and start_dttm='2019-05-29'