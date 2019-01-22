package main


import ("fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	)
	

	
type Restaurant struct {
	DBID         bson.ObjectId     `json:"id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"address line 2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"type_of_food" bson:"typeOfFood"`
}	
	

	
func main(){
	
	
	url := "localhost"			
	data := Restaurant{}
	
	fmt.Println("Enter the filename:")	//get file name from stdin
	read:= bufio.NewReader(os.Stdin)
	fname,err:=read.ReadString('\n')
	if err!=nil{		
		log.Fatal(err)
	}
	
	fname = strings.Trim(fname,"\n")	//remove trailing spaces from fname
	
	f,err := os.Open(fname)
	if err!=nil{		
		log.Fatal(err)
	}
	
	
	filep := bufio.NewReader(f)		//create a file reader
	
	
	
	session ,err := mgo.Dial(url)		//connect to the mongo server
	if err != nil{
		log.Fatal(err)
	}
	defer session.Close()
	
	c  := session.DB("restaurant_db").C("restaurant")	//create an pointer to collection
	
	for {						//json file read linewise
	buff, err := filep.ReadBytes('\n')
	if err != nil {
		break
	}
	
	json.Unmarshal(buff,&data)			//unmarshall in data 
	
	data.DBID = bson.ObjectIdHex(bson.NewObjectId().Hex())	//create a new objectid for each mongo doc
	
	err = c.Insert(&data)			//insert into database	
	if err!=nil{
	log.Fatal(err)
	}
	
	} 

}	
