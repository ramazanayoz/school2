alter table employees
add country varchar(100);

alter table employees
add bDate date;



alter table employees 
modify column bDate year; # tip değiştiriyor


alter table employees
drop column country;

truncate table orderdetails;






#select*from company.employees;


#alter table employees 
#modify column BirthDate year;

#delete   truncate

#ddl data definatiom  language 135


#truncate deleteden daha hızlı çünkü log bırakmıyor