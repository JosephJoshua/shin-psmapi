package models

import (
	"database/sql"
	"fmt"
	"shin-psmapi/db"
	"shin-psmapi/forms"
	"shin-psmapi/utils"
	"time"
)

type Servisan struct {
	NomorNota          int                  `json:"nomor_nota" gorm:"primaryKey;type:integer GENERATED ALWAYS AS IDENTITY (INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1);not null"`
	Tanggal            time.Time            `json:"tanggal" gorm:"->;autoCreateTime;not null;default:CURRENT_TIMESTAMP"`
	NamaPelanggan      string               `json:"nama_pelanggan" gorm:"size:256;not null"`
	NoHp               string               `json:"no_hp" gorm:"size:32"`
	TipeHp             string               `json:"tipe_hp" gorm:"size:128;not null"`
	Imei               string               `json:"imei" gorm:"size:16"`
	KondisiHp          string               `json:"kondisi_hp" gorm:"size:512"`
	Kerusakan          string               `json:"kerusakan" gorm:"size:256;not null"`
	YangBlmDicek       string               `json:"yang_blm_dicek" gorm:"size:128"`
	Kelengkapan        string               `json:"kelengkapan" gorm:"size:128"`
	Warna              string               `json:"warna" gorm:"size:128"`
	KataSandiPola      string               `json:"kata_sandi_pola" gorm:"size:128"`
	IDTeknisi          int                  `json:"id_teknisi" gorm:"type:integer;not null"`
	IDSales            int                  `json:"id_sales" gorm:"type:integer;not null"`
	Teknisi            Teknisi              `json:"-" gorm:"foreignKey:IDTeknisi;constraint:OnDelete:RESTRICT;"`
	Sales              Sales                `json:"-" gorm:"foreignKey:IDSales;constraint:OnDelete:RESTRICT;"`
	Status             utils.StatusServisan `json:"status" gorm:"type:status_servisan;not null"`
	TanggalKonfirmasi  sql.NullTime         `json:"tanggal_konfirmasi"`
	IsiKonfirmasi      string               `json:"isi_konfirmasi" gorm:"size:512"`
	Biaya              float64              `json:"biaya" gorm:"type:double precision;not null;default:0"`
	Diskon             int                  `json:"diskon" gorm:"check:valid_diskon,diskon >= 0 AND diskon <= 100;not null;default:0"`
	DP                 float64              `json:"dp" gorm:"type:double precision;not null;default:0"`
	TambahanBiaya      float64              `json:"tambahan_biaya" gorm:"type:double precision;not null;default:0"`
	TotalBiaya         float64              `json:"total_biaya" gorm:"->;type:double precision GENERATED ALWAYS AS ((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya)) STORED;not null"`
	HargaSparepart     float64              `json:"harga_sparepart" gorm:"->;type:double precision;not null;default:0"`
	Sisa               float64              `json:"sisa" gorm:"->;type:double precision GENERATED ALWAYS AS (((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya) - dp)) STORED;not null"`
	LabaRugi           float64              `json:"laba_rugi" gorm:"->;type:double precision;not null;default:0"`
	TanggalPengambilan sql.NullTime         `json:"tanggal_pengambilan"`
}

func (Servisan) TableName() string {
	return "servisan"
}

type ServisanModel struct{}

func (ServisanModel) All(form forms.GetAllServisanForm) ([]Servisan, error) {
	var (
		servisanList []Servisan
		query        string
		params       []interface{}
	)

	if form.SearchBy != utils.ServisanReturnAll {
		query += "LOWER(?) LIKE '%' || ? || '%'"
		params = append(params, string(form.SearchBy), form.SearchQuery)
	}

	if !form.MinDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "tanggal >= ?"

		// MUST convert to ISO8601/RFC3339 format first before sending it to the postgres db
		params = append(params, utils.ToRFC3339TimeString(form.MinDate))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "tanggal <= ?"

		// MUST convert to ISO8601/RFC3339 format first before sending it to the postgres db
		params = append(params, utils.ToRFC3339TimeString(form.MaxDate))
	}

	if err := db.GetDB().Where(query, params...).Find(&servisanList).Error; err != nil {
		return servisanList, err
	}

	return servisanList, nil
}

func (ServisanModel) ByNomorNota(nomorNota int) (Servisan, error) {
	var servisan Servisan

	res := db.GetDB().Where("nomor_nota = ?", nomorNota).Find(&servisan)

	if res.Error != nil {
		return servisan, res.Error
	}

	if res.RowsAffected < 1 {
		return servisan, fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	return servisan, nil
}

func (ServisanModel) Create(form forms.CreateServisanForm) (nomorNota int, err error) {
	servisan := Servisan{
		NamaPelanggan:      form.NamaPelanggan,
		NoHp:               form.NoHp,
		TipeHp:             form.TipeHp,
		Imei:               form.Imei,
		KondisiHp:          form.KondisiHp,
		Kerusakan:          form.Kerusakan,
		YangBlmDicek:       form.YangBlmDicek,
		Kelengkapan:        form.Kelengkapan,
		Warna:              form.Warna,
		KataSandiPola:      form.KataSandiPola,
		IDTeknisi:          form.IDTeknisi,
		IDSales:            form.IDSales,
		Status:             form.Status,
		TanggalKonfirmasi:  utils.ToNullableTime(form.TanggalKonfirmasi),
		IsiKonfirmasi:      form.IsiKonfirmasi,
		Biaya:              form.Biaya,
		Diskon:             form.Diskon,
		DP:                 form.DP,
		TambahanBiaya:      form.TambahanBiaya,
		TanggalPengambilan: getTanggalPengambilan(form.Status),
	}

	err = db.GetDB().Create(&servisan).Error
	return servisan.NomorNota, err
}

func getTanggalPengambilan(s utils.StatusServisan) sql.NullTime {
	if s == utils.StatusServisanJadiSudahDiambil || s == utils.StatusServisanTdkJadiSudahDiambil {
		return sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	return sql.NullTime{}
}
