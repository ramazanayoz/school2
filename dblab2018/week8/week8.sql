update customers
set ContactName="Gul Sude Demircan"
where CustomerID=3 ;


update customers
set ContactName="Gul Sude Demircan" , City="Mugla"
where Country="Argentina" ;

select OrgerID
from orders
where CustomerID=3 ;


delete from orderdetails
where OrderId = 10365;


delete from orders
where CustomerID = 3 ;


delete from customers
where CustomerID = 3 ;


select city
from customers
union
select city
from suppliers;
