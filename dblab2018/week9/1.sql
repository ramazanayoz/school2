#use company;
load data local infile "C:\\Users\\eayoz\\Desktop\\week9\\data\\customers.csv"
into table customers
fields terminated by ';'
ignore 1 lines; 





select count(customers.CustomerID) as numberofCostumers, customers.Country
from customers
group by customers.Country
order by count(customers.CustomerID) desc;


select count(products.ProductID), suppliers.SupplierName
from products join suppliers on products.SupplierID = suppliers.SupplierID
group by products.SupplierID
order by count(products.ProductID) desc;
