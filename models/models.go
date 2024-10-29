package models

type SecretRDSJson struct {
	Username            string `json:"username"` // Alt izq + 96
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

type Category struct {
	CategID   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type Secretaria struct {
	SecreID          int    `json:"SecreID"`
	SecreNombre      string `json:"nombre"`
	SecreDescripcion string `json:"descripcion"`
	SecreActivo      bool   `json:"activo"`
	SecreTelefono    string `json:"telefono"`
	SecreCorreo      string `json:"correo"`
}

type Direccion struct {
	DirecID          int    `json:"DirecID"`
	SecreID          int    `json:"SecreID"`
	DirecNombre      string `json:"nombre"`
	DirecDescripcion string `json:"descripcion"`
	DirecActivo      bool   `json:"activo"`
	DirecTelefono    string `json:"telefono"`
	DirecCorreo      string `json:"correo"`
}

type Jefatura struct {
	JefaID          int    `json:"JefaID"`
	DirecID         int    `json:"DirecID"`
	JefaNombre      string `json:"nombre"`
	JefaDescripcion string `json:"descripcion"`
	JefaActivo      bool   `json:"activo"`
	JefaTelefono    string `json:"telefono"`
	JefaCorreo      string `json:"correo"`
}
