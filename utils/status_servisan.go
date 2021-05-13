package utils

import "database/sql/driver"

type StatusServisan string

const (
	StatusServisanSedangDikerjakan    StatusServisan = "Sedang dikerjakan"
	StatusServisanPending             StatusServisan = "Pending"
	StatusServisanJadiBelumDiambil    StatusServisan = "Jadi (Belum diambil)"
	StatusServisanJadiSudahDiambil    StatusServisan = "Jadi (Sudah diambil)"
	StatusServisanTdkJadiBelumDiambil StatusServisan = "Tidak jadi (Belum diambil)"
	StatusServisanTdkJadiSudahDiambil StatusServisan = "Tidak jadi (Sudah diambil)"
)

func IsValidServisanStatus(s string) bool {
	return s == string(StatusServisanSedangDikerjakan) || s == string(StatusServisanPending) ||
		s == string(StatusServisanJadiBelumDiambil) || s == string(StatusServisanJadiSudahDiambil) ||
		s == string(StatusServisanTdkJadiBelumDiambil) || s == string(StatusServisanTdkJadiSudahDiambil)
}

func (s *StatusServisan) Scan(value interface{}) error {
	*s = StatusServisan(value.(string))
	return nil
}

func (s StatusServisan) Value() (driver.Value, error) {
	return string(s), nil
}
