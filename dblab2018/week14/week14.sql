CREATE TABLE `employees_backup`(
`id` INT AUTO_INCREMENT,
`Employee_ID` int(11) NOT NULL,
`LastName` varchar(45) NOT NULL,
`FirstName` varchar(45) NOT NULL,
`BirthDate` varchar(45) NOT NULL,
`changedat` DATETIME DEFAULT NULL,
`action` VARCHAR(50) DEFAULT NULL,
PRIMARY KEY (`id`)
);

select * from employees_backup;

update Employees set LastName = "KORKMAZ" where EmployeeID = 1;

update Employees set BirthDate = "1997-05-21" where EmployeeID = 10;
 
insert into Employees (EmployeeID,LastName,FirstName,BirthDate,Photo,Notes)
		values ("11","KORKMAZ","BERKAY","1997-05-21","EmpID11.pic","Intern");
        

delete from Employees where EmployeeID = 11;