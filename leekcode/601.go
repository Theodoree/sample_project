package main

/* 601. 体育馆的人流量
SELECT
	tmp.id,
	tmp.date,
	tmp.people
FROM (
SELECT
	id AS id,
	date AS date,
	people AS people,
FROM
	stadium
WHERE
	people > 100
)

*/

/*
select distinct Num as ConsecutiveNums from
(select
   Num,  //这里使用的还是一个用户变量,按照主键排序,如果当前Num等于上一个num cnt++
	@cnt:=if(@pre=num,@cnt:=@cnt+1,@cnt:=1) as n,
 @pre:=Num from Logs,
 (select @pre := -1,@cnt := 0)as init
)as t
where t.n>=3;
*/