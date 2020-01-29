package main

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"net/http"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {

	cfg := m.(*Config)

	log.Printf("resourceServerRead: %s, %s", cfg.url, cfg.token)
	client := &http.Client{}

	req, err := http.NewRequest("GET", cfg.url+"/api/ipam/prefixes/", nil)

	if err != nil {
		log.Printf("resourceServerRead err: %s", err)
		return nil
	}

	req.Header.Add("Authorization", "Token "+cfg.token)
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("resourceServerRead err: %s", err)
		return nil
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Printf("resourceServerRead err: %s", err)
	// 	return nil
	// }

	// if body != nil {
	// 	log.Printf("resourceServerRead body: %s", body)
	// }

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
