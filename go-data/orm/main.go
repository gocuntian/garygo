package main

import (
    "fmt"
)

func main(){
    // member:=new(Member)
    // member.Name="name2"
    // err:= InsertDb(member)
    // if err!=nil{
    //     fmt.Println("insert fail")
    // }else{
    //     fmt.Println("success")
    // }
//     //插入多条数据
//     members:=make([]Member,0)
//     members = append(members,Member{Name:"username2"})
//     members = append(members,Member{Name:"username3"})
//     members = append(members,Member{Name:"username7"})
//     members = append(members,Member{Name:"username6"})
//     members = append(members,Member{Name:"username5"})
//    // fmt.Println(members)
//     err:=InsertBatch(members)
//     if err!=nil{
//         fmt.Println("insert fail")
//     }else{
//         fmt.Println("success")
//     }

//       //使用指针Slice插入多条数据
//         members:=make([]*Member,0)
//         members = append(members,&Member{Name:"username990"})
//         members = append(members,&Member{Name:"username991"})
//         members = append(members,&Member{Name:"username992"})
//         err:=InsertMore(members)
//         if err!=nil{
//             fmt.Println("insert fail")
//         }else{
//             fmt.Println("success")
//         }

        //   var user=Member{Id:5}
        //   _,err:= engine.Get(&user)
        //   if err!=nil{
        //      fmt.Println("fail")
        //   }
        //   fmt.Println(user.Name)
       
        fmt.Println("============================================\r\n")
        // var everyone []Member
        // err:=engine.Find(&everyone)
        // if err!=nil{
        //     fmt.Println("ddddd")
        // }
        // fmt.Println(everyone)

        // members:=make(map[int64]Member)
        // err:=engine.Find(&members)
        // if err!=nil{
        //     fmt.Println("fail")
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.Where("id > ?",3).Limit(10,5).Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.Find(&members,&Member{Salt:"xingcuntian"})
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members) 

        // var members []Member
        // err:=engine.In("id",11,30,40).Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.In("id",[]int{1,44,43}).Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.In("salt",[]string{"xingcuntian","test1"}).Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.Omit("created","updated").Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // var members []Member
        // err:=engine.Cols("user_name","salt").Find(&members)
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(members)

        // total,err:=engine.Where("id > ?",1).Count(&Member{})
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(total)

        // var members []Member
        // err:=engine.Distinct("salt","age").Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // total,err:=engine.Count(&Member{Salt:"xingcuntian"})
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(total)

        // var members []Member
        // err:=engine.Cols("user_name","age").Where("`age` > ?",10).And("`age` < 30").Limit(10,5).Desc("id").Find(&members)
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(members)

        // member:=new(Member)
        // member.Name="john"
        // affected,err:=engine.Id(44).Update(member)
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(affected)
        // fmt.Println(member)

       //  age>0 ====>0  (不可用)
        // member:=new(Member)
        // member.Age=0
        // affected,err:=engine.Id(39).Update(member)
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(affected)
        
        //  member:=new(Member)    ===  &Member{}
        // affected,err:=engine.Id(39).Cols("age").Update(&Member{})
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(affected)  


    //     affected,err:=engine.Table(new(Member)).Id(38).Update(map[string]interface{}{"age":0})
    //    if err!=nil{
    //       fmt.Println(err)
    //     }
    //     fmt.Println(affected)  

    // affected,err:=engine.Table(&Member{}).Id(37).Update(map[string]interface{}{"age":0})
    //    if err!=nil{
    //       fmt.Println(err)
    //     }
    //     fmt.Println(affected)  


     //  var member Member
    //    has,err:=engine.Id(47).Get(&member)
    //    if err!=nil{
    //       fmt.Println(err)
    //     }
    //     fmt.Println(has)
    //     fmt.Println(member) 

        // affected,err:=engine.Id(1).Update(&Member{Age:30,Version:1})
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(affected)  

        // affected,err:=engine.Id(46).Delete(&Member{})
        
        // affected,err:=engine.Delete(&Member{Age:30})
        
        // if err!=nil{
        //   fmt.Println(err)
        // }
        // fmt.Println(affected) 


        // sql:="select*from user"
        // result,err:=engine.Query(sql)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(result)


        //Get方法
        // member:=new(Member)
        // has,err:=engine.Id(45).Get(member)
        // //has,err:=engine.Id(xorm.PK{1,2}).Get(member)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(has)
        // fmt.Println(member)

        // member:=&Member{}
        // has,err:=engine.Id(45).Get(member)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(has)
        // fmt.Println(member)

        // member:=new(Member)
        // has,err:=engine.Where("user_name=?","john").Get(member)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(has)
        // fmt.Println(member)

        // member:=&Member{Id:45}
        // has,err:=engine.Get(member)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(has)
        // fmt.Println(member)

        // member:=&Member{Name:"john"}
        // has,err:=engine.Get(member)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(has)
        // fmt.Println(member)

        //Find方法使用
        //1)传入Slice用于返回数据
        // members:=make([]Member,0)
        // err:=engine.Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        //指针
        // pmembers:=make([]*Member,0)
        // err:=engine.Find(&pmembers)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(pmembers[0].Name)

        //2)传入Map用户返回数据，map必须为map[int64]Userinfo的形式，map的key为id，因此对于复合主键无法使用这种方式。

        // members:=make(map[int64]Member)
        // err:=engine.Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // pmembers:=make(map[int64]*Member)
        // err:=engine.Find(&pmembers)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(pmembers)
        // fmt.Println(pmembers[38].Name)

        // members:=make([]Member,0)
        // err:=engine.Where("age > ? or user_name=?",10,"john").Limit(2,2).Find(&members)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(members)

        // Join用法

        // users:=make([]UserGroup,0)
        // err:=engine.Join("INNER","prefix_group","prefix_group.id = prefix_user.group_id").Find(&users)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(users)

        // users:=make([]UserGroup,0)
        // err:=engine.Sql("select prefix_user.*,prefix_group.name from prefix_user,prefix_group where prefix_user.group_id=prefix_group.id").Find(&users)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(users)


        // users:=make([]UserGroupType,0)
        // err:=engine.Table("prefix_user").Join("INNER","prefix_group","prefix_group.id=prefix_user.group_id").Join("INNER","prefix_type","prefix_type.id=prefix_user.type_id").Find(&users)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(users)

        // Where and like 使用
        // users:=make([]UserGroupType,0)
        // err:=engine.Table("prefix_user").Join("INNER","prefix_group","prefix_group.id = prefix_user.group_id").Join("INNER","prefix_type","prefix_type.id = prefix_user.type_id").Where("prefix_user.name like ?","%gary%").Find(&users)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // fmt.Println(users)

        //Iterate方法
        //terate方法提供逐条执行查询到的记录的方法，他所能使用的条件和Find方法完全相同
        err:=engine.Where("age > ? or salt=?",0,"test").Iterate(new(Member),func(i int,bean interface{})error{
            member:=bean.(*Member)
            fmt.Println(member)
           return nil
        })
        if err!=nil{
            fmt.Println(err)
        }
       // fmt.Println(member)
        






















        
 }