provider "netbox" {
	token = "7ae6f0bff5561988c167a3344975edd3e8ddd05e"
	url = "https://ipam.ent.wixpress.com"
}

resource "netbox_server" "my-server" {
	address = "1.2.3.4"
}
