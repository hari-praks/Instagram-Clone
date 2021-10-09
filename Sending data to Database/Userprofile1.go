package main
import (
	"net/http"
	"html/template"
	"log"
	"fmt"
	"time"
	"context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type UserInfo struct{
	Id string           `bson:"_id`
	Email string		`bson:"_email`	
	Name string			`bson:"_name`
	Password string		`bson:"_pwd`
}

func main(){

	tpl := template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			tpl.Execute(w,nil)
			return 
		}
			u := UserInfo{
			Id: r.FormValue("Signupid"),
			Email: r.FormValue("emailid"),
			Name : r.FormValue("name"),
			Password : r.FormValue("pwd1"),
		}
		
		_ = u
			tpl.Execute(w,struct{success bool} {true})
			senddata(u.Id,u.Email,u.Name,u.Password)
		})
			http.ListenAndServe(":8080",nil)
		
}
func senddata(formid,f_Email,f_name,f_password string){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
		    log.Fatal(err)}

		 ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		    err = client.Connect(ctx)
		    if err != nil {
		        log.Fatal(err)
		    }
		    defer client.Disconnect(ctx)
		fmt.Println("Connected to MongoDB")
		collection := client.Database("godb").Collection("html")
		insertResult, err := collection.InsertOne(ctx,bson.D{

		        {"Id",formid},
		        {"Email",f_Email},
		        {"Name",f_name},
		        {"Password",f_password},
		})

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		log.Println("error: ",err)
	

}
