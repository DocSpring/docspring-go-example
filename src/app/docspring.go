package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/DocSpring/docspring-go"
	"github.com/antihax/optional"
	"github.com/urfave/cli/v2"
)

var docSpringAPIToken string

func main() {
	// parse flags
	app := &cli.App{
		Name:  "DocSpring Go Example",
		Usage: "DocSpring Go Example - Upload and list templates",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "api_token",
				Usage:       "api_token",
				Destination: &docSpringAPIToken,
				Value:       "<API Token>",
			},
		},
		Action: func(c *cli.Context) error {
			apiToken := c.String("api_token")
			return listFiles(apiToken)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func listFiles(apiToken string) error {
	docSpringCfg := docspring.NewConfiguration()
	// docSpringCfg.BasePath = "http://api.docspring.local:3000/api/v1"

	docSpringCfg.AddDefaultHeader("Authorization",
		fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(apiToken))))
	client := docspring.NewAPIClient(docSpringCfg)

	pdf_file, err := os.Open("./fw8ben.pdf")
	if err != nil {
		panic(err)
	}

	pendingTemplate, _, err := client.PDFApi.CreatePDFTemplate(
		context.Background(), pdf_file, "fw8ben.pdf",
		&docspring.CreatePDFTemplateOpts{TemplateParentFolderId: optional.EmptyString()})

	if err != nil {
		fmt.Println(fmt.Errorf("client.PDFApi.CreatePDFTemplate: got err :%w", err))
		// Cast error as GenericOpenAPIError to print the response body
		openAPIErr := err.(docspring.GenericOpenAPIError)
		fmt.Println(string(openAPIErr.Body()))

		return err
	}

	fmt.Printf("%+v\n", pendingTemplate)

	templates, _, err := client.PDFApi.ListTemplates(context.Background(), &docspring.ListTemplatesOpts{})
	if err != nil {
		fmt.Println(fmt.Errorf("client.PDFApi.ListTemplates: got err :%w", err))
		// Cast error as GenericOpenAPIError to print the response body
		openAPIErr := err.(docspring.GenericOpenAPIError)
		fmt.Println(string(openAPIErr.Body()))
		return err
	}

	fmt.Printf("%+v", templates)
	return nil
}
