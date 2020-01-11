select *
from customers
order by CustomerName;

create index abc on customers(CustomerName);

select *
from customers
order by CustomerName;

explain select *
from customers
order by CustomerName;


alter table customers drop index abc ;

create view tmp as 
select * 
from customers
order by CustomerName ;


create view tmp as 
select * 
from customers
order by CustomerName ;


insert into tmp
values (99,"asdfhg","aad","aaa","aa","aa","a");


