call get_customers();

call get_customers_by_country("belgium");

set @country = "usa";

call get_customers_by_country(@country);

select @country;

call get_no_customers_by_country(@count, "germany");

select @count;

set @counter =10;
set @inc = 5;

call counter(@counter,@inc);

select @counter,@inc;