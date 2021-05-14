package models

import (
	"fmt"
	"reflect"
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
	TanggalKonfirmasi  utils.NullTime       `json:"tanggal_konfirmasi" gorm:"type:timestamp with time zone"`
	IsiKonfirmasi      string               `json:"isi_konfirmasi" gorm:"size:512"`
	Biaya              float64              `json:"biaya" gorm:"type:double precision;not null;default:0"`
	Diskon             int                  `json:"diskon" gorm:"check:valid_diskon,diskon >= 0 AND diskon <= 100;not null;default:0"`
	DP                 float64              `json:"dp" gorm:"type:double precision;not null;default:0"`
	TambahanBiaya      float64              `json:"tambahan_biaya" gorm:"type:double precision;not null;default:0"`
	TotalBiaya         float64              `json:"total_biaya" gorm:"->;type:double precision GENERATED ALWAYS AS ((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya)) STORED;not null"`
	HargaSparepart     float64              `json:"harga_sparepart" gorm:"->;type:double precision;not null;default:0"`
	Sisa               float64              `json:"sisa" gorm:"->;type:double precision GENERATED ALWAYS AS (((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya) - dp)) STORED;not null"`
	LabaRugi           float64              `json:"laba_rugi" gorm:"->;type:double precision;not null;default:0"`
	TanggalPengambilan utils.NullTime       `json:"tanggal_pengambilan" gorm:"type:timestamp with timezone"`
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
		if form.SearchBy == utils.ServisanSearchByNomorNota {
			query += "nomor_nota::TEXT LIKE '%' || ? || '%'"
		} else if form.SearchBy == utils.ServisanSearchByStatus {
			query += "LOWER(status::TEXT) LIKE '%' || LOWER(?) || '%'"
		} else {
			// form.SearchBy has been validated beforehand so this
			// shouldn't be vulnerable to SQL injection, AFAIK.
			query += "LOWER(" + string(form.SearchBy) + ") LIKE '%' || LOWER(?) || '%'"
		}

		params = append(params, form.SearchQuery)
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

func (ServisanModel) ByNomorNota(nomorNota int) (*Servisan, error) {
	var servisan *Servisan

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

func (ServisanModel) Update(nomorNota int, form forms.UpdateServisanForm) error {
	newServisan := make(map[string]interface{}, 26)

	insertToMapIfExists(form.NamaPelanggan, "nama_pelanggan", &newServisan)
	insertToMapIfExists(form.NoHp, "no_hp", &newServisan)
	insertToMapIfExists(form.TipeHp, "tipe_hp", &newServisan)
	insertToMapIfExists(form.Imei, "imei", &newServisan)
	insertToMapIfExists(form.KondisiHp, "kondisi_hp", &newServisan)
	insertToMapIfExists(form.Kerusakan, "kerusakan", &newServisan)
	insertToMapIfExists(form.YangBlmDicek, "yang_blm_dicek", &newServisan)
	insertToMapIfExists(form.Kelengkapan, "kelengkapan", &newServisan)
	insertToMapIfExists(form.Warna, "warna", &newServisan)
	insertToMapIfExists(form.KataSandiPola, "kata_sandi_pola", &newServisan)
	insertToMapIfExists(form.IDTeknisi, "id_teknisi", &newServisan)
	insertToMapIfExists(form.IDSales, "id_sales", &newServisan)
	insertToMapIfExists(form.Status, "status", &newServisan)
	insertToMapIfExists(form.IsiKonfirmasi, "isi_konfirmasi", &newServisan)
	insertToMapIfExists(form.Biaya, "biaya", &newServisan)
	insertToMapIfExists(form.Diskon, "diskon", &newServisan)
	insertToMapIfExists(form.DP, "dp", &newServisan)
	insertToMapIfExists(form.TambahanBiaya, "tambahan_biaya", &newServisan)

	newServisan["tanggal_konfirmasi"] = utils.ToNullableTime(form.TanggalKonfirmasi)

	if form.Status != nil {
		newServisan["tanggal_pengambilan"] = getTanggalPengambilan(*form.Status)
	}

	res := db.GetDB().Model(&Servisan{NomorNota: nomorNota}).Updates(newServisan)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 && len(newServisan) > 0 {
		return fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	return nil
}

func (ServisanModel) Delete(nomorNota int) error {
	res := db.GetDB().Delete(&Servisan{}, nomorNota)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	return nil
}

func insertToMapIfExists(formVal interface{}, key string, m *map[string]interface{}) {
	if formVal == nil || (reflect.ValueOf(formVal).Kind() == reflect.Ptr && reflect.ValueOf(formVal).IsNil()) {
		return
	}

	(*m)[key] = formVal
}

func getTanggalPengambilan(s utils.StatusServisan) utils.NullTime {
	if s == utils.StatusServisanJadiSudahDiambil || s == utils.StatusServisanTdkJadiSudahDiambil {
		return utils.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	return utils.NullTime{}
}
