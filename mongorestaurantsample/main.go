package main

import (
	"fmt"
	"os"
	"domain"
	"strings"
	"log"
	"errors"
	"strconv"
	dbrepo "dbrepository"
	mongoutils "utils"
)


//getData gets comma seperated string from the input and splits it to get the field values  
func getData(data string)(domain.Restaurant){
	
	values := domain.Restaurant{}
	//split the string by "," 
	lst := strings.SplitN(data,",",-1)
	
	//if the number of values is more than the fields 
	if len(lst) > 8{
		err := errors.New("TOO MANY VALUES")
		log.Fatal(err)
	}else if len(lst) < 8{
		err := errors.New("NOT ENOUGH VALUES(provide all 8 values, blank for no value)")
		log.Fatal(err)
	}
	fmt.Print(len(lst))
	
	values.Name = lst[0]
	values.Address = lst[1]	
	values.AddressLine2 = lst[2]
	values.URL = lst[3]
	values.Outcode = lst[4]
	values.Postcode = lst[5]
	rating,_ := strconv.ParseFloat(lst[6],32)
	values.Rating = float32(rating)
	values.TypeOfFood = lst[7]
			
	return values		
}

func print_command_info(){
	fmt.Println("\n---Command info : go run <filename> [flag] [input]---\n")
	fmt.Println("\n-i for command info\n-s <search pattern> for search\n-g <id> to get record by id\n-ga to get all the records")
	fmt.Println("\n-st <record> to insert a record\n    eg:go run file.go -st value1,value2,value3,.. \n\n-d <id> to delete by id\n-fn <name> to find by name\n-ff <foodType> to find by type of food\n-fp <postcode> to find by postcode")
	fmt.Println("\n\n---Enter the flag and info to get or put data---\n")
	os.Exit(0)
}


func main() {
	
	args := os.Args
	_len := len(args)
	
	if _len == 1{
		print_command_info()
	}else if _len > 3{
	
		err := errors.New("TOO MANY ARGUMENTS")
		log.Fatal(err)
	}else{
	
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant_db"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)

	//Run sample commands
	
	switch args[1]{
	
	case "-i" :
		if _len == 2{ 
		print_command_info()
		}else{
		err := errors.New("TOO MANY ARGUMENTS")
		log.Fatal(err)
		}
		
	case "-s" :
		value,err := repoaccess.Search(args[2])
		if err!=nil{
			log.Fatal(err)
		}
		for _,v := range value{
			fmt.Println(v)
		}		
	case "-st" : 
		values := getData(args[2])
		ans,err:= repoaccess.Store(&values)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(ans)
	case "-g" :
		data,err :=repoaccess.Get(domain.ID(args[2]))
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(data)
	
	case "-ga" :
		
		if _len == 2{
		result,err := repoaccess.GetAll()
		if err!=nil{
			log.Fatal(err)
		}
	
		for _,v := range result{
			fmt.Printf("%v\n\n",v)
		}	
		}else{
			err := errors.New("TOO MANY ARGUMENTS")
			log.Fatal(err)
		}
	case "-d" :
		err := repoaccess.Delete(domain.ID(args[2]))
		if err != nil{
			log.Fatal(err)
		}

	case "-ff" :

		value,err := repoaccess.FindByTypeOfFood(args[2])
		if err!=nil{
			log.Fatal(err)
		}
		for _,v := range value{
			fmt.Println(v)
		}	
	

	case "-fn" : 
		value,err := repoaccess.FindByName(args[2])
		if err!=nil{
			log.Fatal(err)
		}
		for _,v := range value{
			fmt.Println(v)
		}	

	case "-fp" :
		value,err := repoaccess.FindByTypeOfPostCode(args[2])
		if err!=nil{
			log.Fatal(err)
		}
		for _,v := range value{
			fmt.Println(v)
		}	

	}
	
	}	

}	
