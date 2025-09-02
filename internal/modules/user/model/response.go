package user_model

type UserResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	IDRole      int    `json:"id_role"`
	IDPuskesmas int    `json:"id_puskesmas"`
	IDPegawai   int    `json:"id_pegawai"`
	IDPoli      *int   `json:"id_poli"`
	IDPustu     *int   `json:"id_pustu"`
}
