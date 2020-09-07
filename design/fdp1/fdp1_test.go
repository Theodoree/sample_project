package fdp1

import (
    "fmt"
    "testing"
)

/*
tip:
    场景:
        大致的思路就是将可变和不可变的变量进行拆分,从而使得不可变的这一块可以被多个地方引用(同一个进程,地址相同,那么数据也相同) */

func TestNewTeacher(t *testing.T) {


    teacherLi :=NewTeacher("李",GetClass("A班"))
    fmt.Printf("班级:%s;姓名:%s;职位:%s \n",teacherLi.Class(),teacherLi.Name(),teacherLi.Position())
    studentMing :=NewStudent("小明",GetClass("A班"))
    fmt.Printf("班级:%s;姓名:%s;职位:%s \n",studentMing.Class(),studentMing.Name(),studentMing.Position())


}

