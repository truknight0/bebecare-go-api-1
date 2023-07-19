package firebase

import (
	"bebecare-go-api-1/utils"
	"bebecare-go-api-1/utils/log"
	"context"
	firebsego "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

var app *firebsego.App
var ctx context.Context

func init() {
	serviceHome := os.Getenv("BEBECARE_GO_API_1_HOME")
	if serviceHome == "" {
		serviceHome = ".."
	}
	serviceKeyFile := fmt.Sprintf("%s/config/google-service-account.json", serviceHome)

	ctx = context.Background()
	conf := &firebsego.Config{
		DatabaseURL: "https://parking-cone.firebaseio.com",
	}
	opt := option.WithCredentialsFile(serviceKeyFile)

	var err error
	app, err = firebsego.NewApp(ctx, conf, opt)
	if err != nil {
		log.ERROR(err.Error())
	}
}

func UpdateReserveStatus(receiptNo string, status int) {
	client, err := app.Database(ctx)
	if err != nil {
		log.ERROR(fmt.Sprintf("%s : %s", "Error initializing database client:", err.Error()))
		return
	}
	node := ""
	if utils.IsDev() {
		node = "dev"
	} else {
		node = "real"
	}

	path := fmt.Sprintf("%s/parking/receiptNo/%s/matching", node, receiptNo)
	ref := client.NewRef(path)
	err = ref.Set(ctx, map[string]string{
		"reserveStatus": fmt.Sprintf("%d", status),
	})
	if err != nil {
		log.ERROR(fmt.Sprintf("%s : %s", "Error setting value:", err.Error()))
	}
}
