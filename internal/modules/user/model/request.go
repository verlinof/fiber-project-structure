package user_model

type CreateUserRequest struct {
	Username    string `json:"username" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	IDRole      int    `json:"id_role" validate:"number,min=1,required"`
	IDPuskesmas int    `json:"id_puskesmas" validate:"number,min=1,required"`
	IDPegawai   int    `json:"id_pegawai"`
	IDPoli      *int   `json:"id_poli"`
	IDPustu     *int   `json:"id_pustu"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	IDRole      int    `json:"id_role"`
	IDPuskesmas int    `json:"id_puskesmas"`
	IDPegawai   int    `json:"id_pegawai"`
	IDPoli      *int   `json:"id_poli"`
	IDPustu     *int   `json:"id_pustu"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}
