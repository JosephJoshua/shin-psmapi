package utils

import "database/sql/driver"

type ServisanSearchBy string

const (
	ServisanReturnAll         ServisanSearchBy = ""
	ServisanSearchByNama      ServisanSearchBy = "nama_pelanggan"
	ServisanSearchByNomorNota ServisanSearchBy = "nomor_nota"
	ServisanSearchByStatus    ServisanSearchBy = "status"
	ServisanSearchByNomorHp   ServisanSearchBy = "no_hp"
	ServisanSearchByTipeHp    ServisanSearchBy = "tipe_hp"
)

func IsValidServisanSearchByColumn(s string) bool {
	return s == string(ServisanReturnAll) || s == string(ServisanSearchByNama) ||
		s == string(ServisanSearchByNomorNota) || s == string(ServisanSearchByStatus) ||
		s == string(ServisanSearchByNomorHp) || s == string(ServisanSearchByTipeHp)
}

func (s *ServisanSearchBy) Scan(value interface{}) error {
	*s = ServisanSearchBy(value.(string))
	return nil
}

func (s ServisanSearchBy) Value() (driver.Value, error) {
	return string(s), nil
}
