   package main
   import (
             "fmt"
             "github.com/julienschmidt/httprouter"
             "log"
             "net/http"
             "context"
             "time"
               "go.mongodb.org/mongo-driver/bson"
               "go.mongodb.org/mongo-driver/mongo"
               "go.mongodb.org/mongo-driver/mongo/options"
               //"go.mongodb.org/mongo-driver/mongo/readpref"
    )
type simple struct{
    dummyId string
}
      func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        uid := ps.ByName("uid")
        fmt.Fprintf(w, "<your Post if of ID:  %s", uid)
        var p = simple{dummyId:uid}
        retreivefromDB(p.dummyId)
    }     
     func main() {
          router := httprouter.New()
           router.GET("/posts/:uid", getuser)
           log.Fatal(http.ListenAndServe(":8080", router))
          
         }


func retreivefromDB(copyId string){

          client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
           ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
              err = client.Connect(ctx)
              if err != nil {
                  log.Fatal(err)
              }
              defer client.Disconnect(ctx)
          fmt.Println("Connected to MongoDB")
          
          quickstartdb := client.Database("godb")
          usersCollection :=quickstartdb.Collection("postfromusers")        
          filterCursor,err := usersCollection.Find(ctx,bson.M{"Id":copyId})
          if err!=nil{
               log.Fatal(err)
          }
          var FilteredPosts []bson.M
          if err=filterCursor.All(ctx,&FilteredPosts); 
          err!=nil{
               log.Fatal(err)
          }
          fmt.Println(FilteredPosts)
         
}
