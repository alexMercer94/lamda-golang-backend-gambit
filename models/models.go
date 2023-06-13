package models

/*
Modelo para recibir las información del Secret de AWS SecretManager
*/
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

/*
Modelo de la info del Usuario
*/
type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
