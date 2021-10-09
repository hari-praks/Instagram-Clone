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

type PostInfo struct{
	Id string           `bson:"_id`
	Caption string		`bson:"_Caption`	
	ImgURL string			`bson:"_img`
	Time string		`bson:"_time`
}

func main(){

	tpl := template.Must(template.ParseFiles("post.html"))
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			tpl.Execute(w,nil)
			return 
		}
			u := PostInfo{
			Id: r.FormValue("postid"),
			Caption: r.FormValue("emailid"),
			ImgURL : r.FormValue("Img"),
			Time : r.FormValue("posttime"),
		}
		_ = u
			tpl.Execute(w,struct{success bool} {true})
			dataToDB(u.Id,u.Caption,u.ImgURL,u.Time)
		})
			http.ListenAndServe(":8080",nil)
		
}
func dataToDB(postid,caption,imgurl,timestamp string){
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
		collection := client.Database("godb").Collection("postfromusers")
		insertResult, err := collection.InsertOne(ctx,bson.D{

		        {"Id",postid},
		        {"Caption",caption},
		        {"Image",imgurl},
		        {"Time",timestamp},
		})

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		log.Println("error: ",err)


	

}