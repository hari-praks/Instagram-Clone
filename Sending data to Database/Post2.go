package main

import( 
    "fmt"
    "context"
    "log"
    "net/http"
    "time"
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
  
)

func Users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        fmt.Fprint(w, "Thank You! The data is sent to the database.\n")
    }
func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        fmt.Fprint(w, "First Page!\n")
    }
func main(){

  
    router := httprouter.New()
    router.POST("/users",Users)
   

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

if err != nil {
    log.Fatal(err)
}


 ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)
fmt.Println("Connected to MongoDB")
collection := client.Database("godb").Collection("instaprofiles")


insertResult, err := collection.InsertMany(ctx,[]interface{}{
    bson.D{
    {"id","1111"},
    {"Email", "ari@gmail.com"},
    {"name","hari"},
    {"password","1234"}, 
    },
     bson.D{
    {"id","2222"},
    {"Email", "ravi@gmail.com"},
    {"name","ravi"},
    {"password","sometext"},
    },
     bson.D{
    {"id","3333"},
    {"Email", "harshit@gmail.com"},
    {"name","harshit"},
    {"password","nopwd"},
    },
     bson.D{
    {"id","4444"},
    {"Email", "ram@gmail.com"},
    {"name","ram"},
    {"password","5678"},
    }, 
})

fmt.Println("Inserted a single document: ", len(insertResult.InsertedID))
log.Println("error: ",err)
http.ListenAndServe(":8080",router)



}
