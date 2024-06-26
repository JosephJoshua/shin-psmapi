package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/utils"
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
	NamaTeknisi        string               `json:"nama_teknisi" gorm:"->"`
	NamaSales          string               `json:"nama_sales" gorm:"->"`
	TanggalKonfirmasi  utils.NullTime       `json:"tanggal_konfirmasi" gorm:"type:timestamptz"`
	IsiKonfirmasi      string               `json:"isi_konfirmasi" gorm:"size:512"`
	Biaya              float64              `json:"biaya" gorm:"type:double precision;not null;default:0"`
	Diskon             int                  `json:"diskon" gorm:"check:valid_diskon,diskon >= 0 AND diskon <= 100;not null;default:0"`
	DP                 float64              `json:"dp" gorm:"type:double precision;not null;default:0"`
	TambahanBiaya      float64              `json:"tambahan_biaya" gorm:"type:double precision;not null;default:0"`
	TotalBiaya         float64              `json:"total_biaya" gorm:"->;type:double precision GENERATED ALWAYS AS ((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya)) STORED;not null"`
	HargaSparepart     float64              `json:"harga_sparepart" gorm:"type:double precision;not null;default:0"`
	Sisa               float64              `json:"sisa" gorm:"->;type:double precision GENERATED ALWAYS AS (((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya) - dp)) STORED;not null"`
	LabaRugi           float64              `json:"laba_rugi" gorm:"->;type:double precision GENERATED ALWAYS AS ((((((100.0)::double precision - (diskon)::double precision) / (100.0)::double precision) * biaya) + tambahan_biaya) - harga_sparepart) STORED;not null"`
	TanggalPengambilan utils.NullTime       `json:"tanggal_pengambilan" gorm:"type:timestamptz"`
}

type LabaRugiReportItem struct {
	NomorNota 		    	int						`json:"nomor_nota"`
	TanggalPengambilan  time.Time			`json:"tanggal_pengambilan"`
	TipeHp 			    		string				`json:"tipe_hp"`
	Kerusakan 		    	string				`json:"kerusakan"`
	Biaya 			    		float64				`json:"biaya"`
	HargaSparepart 	    float64				`json:"harga_sparepart"`
	LabaRugi 		    		float64				`json:"laba_rugi"`
}

type SisaReportItem struct {
	NomorNota 		    	int						`json:"nomor_nota"`
	TanggalPengambilan  time.Time			`json:"tanggal_pengambilan"`
	TipeHp 			    		string				`json:"tipe_hp"`
	Kerusakan 		    	string				`json:"kerusakan"`
	Biaya 			    		float64				`json:"biaya"`
	DP									float64       `json:"dp"`
	Sisa                float64       `json:"sisa"`
}

type TeknisiReportItem struct {
	NomorNota 		    	int						`json:"nomor_nota"`
	TanggalPengambilan  time.Time			`json:"tanggal_pengambilan"`
	TipeHp 			    		string				`json:"tipe_hp"`
	Kerusakan 		  	  string				`json:"kerusakan"`
	Biaya 			    		float64				`json:"biaya"`
	HargaSparepart 	    float64				`json:"harga_sparepart"`
	LabaRugi 		    		float64				`json:"laba_rugi"`
	Teknisi							string				`json:"teknisi"`
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

	// Search by nama pelanggan by default if a 'search by' column wasn't provided.
	if form.SearchQuery != "" && form.SearchBy == utils.ServisanReturnAll {
		form.SearchBy = utils.ServisanSearchByNama
	}

	if form.SearchBy != utils.ServisanReturnAll && form.SearchQuery != "" {
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
	} else {
		// We only want to filter based on date range when the search query is empty.
		if !form.MinDate.IsZero() {
			if len(params) > 0 {
				query += " AND "
			}

			query += "DATE(tanggal) >= ?"

			params = append(params, form.MinDate.Format(utils.DateOnly))
		}

		if !form.MaxDate.IsZero() {
			if len(params) > 0 {
				query += " AND "
			}

			query += "DATE(tanggal) <= ?"

			params = append(params, form.MaxDate.Format(utils.DateOnly))
		}
	}

	err := db.GetDB().
		Model(&Servisan{}).
		Select("servisan.*, teknisi.nama as nama_teknisi, sales.nama as nama_sales").
		Joins("LEFT JOIN teknisi ON servisan.id_teknisi = teknisi.id").
		Joins("LEFT JOIN sales ON servisan.id_sales = sales.id").
		Where(query, params...).
		Order("nomor_nota ASC").Find(&servisanList).Error

	if err != nil {
		return servisanList, err
	}

	return servisanList, nil
}

func (ServisanModel) ByNomorNota(nomorNota int) (*Servisan, error) {
	var servisan *Servisan

	res := db.GetDB().
		Model(&Servisan{}).
		Select("servisan.*, teknisi.nama as nama_teknisi, sales.nama as nama_sales").
		Joins("LEFT JOIN teknisi ON servisan.id_teknisi = teknisi.id").
		Joins("LEFT JOIN sales on servisan.id_sales = sales.id").
		Where("nomor_nota = ?", nomorNota).Find(&servisan)

	if res.Error != nil {
		return servisan, res.Error
	}

	if res.RowsAffected < 1 {
		return servisan, fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	return servisan, nil
}

func (ServisanModel) LabaRugiReport(form forms.ServisanLabaRugiReportForm) ([]LabaRugiReportItem, error) {
	var (
		labaRugiList []LabaRugiReportItem
		query        string
		params       []interface{}
	)

	if !form.MinDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pengambilan) >= ?"

		params = append(params, form.MinDate.Format(utils.DateOnly))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pengambilan) <= ?"

		params = append(params, form.MaxDate.Format(utils.DateOnly))
	}

	err := db.GetDB().
		Model(&Servisan{}).
		Select("servisan.nomor_nota, servisan.tanggal_pengambilan, servisan.tipe_hp, servisan.kerusakan, servisan.total_biaya as biaya, servisan.harga_sparepart, servisan.laba_rugi").
		Where("status::TEXT LIKE '%Sudah diambil%'").
		Where(query, params...).
		Order("nomor_nota ASC").Find(&labaRugiList).Error

	if err != nil {
		return labaRugiList, err
	}

	return labaRugiList, nil
}

func (ServisanModel) SisaReport(form forms.ServisanSisaReportForm) ([]SisaReportItem, error) {
	var (
		sisaList 		[]SisaReportItem
		query       string
		params      []interface{}
	)

	if !form.MinDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pengambilan) >= ?"

		params = append(params, form.MinDate.Format(utils.DateOnly))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pengambilan) <= ?"

		params = append(params, form.MaxDate.Format(utils.DateOnly))
	}

	err := db.GetDB().
		Model(&Servisan{}).
		Select("servisan.nomor_nota, servisan.tanggal_pengambilan, servisan.tipe_hp, servisan.kerusakan, servisan.total_biaya as biaya, servisan.dp, servisan.sisa").
		Where("status::TEXT LIKE '%Sudah diambil%'").
		Where(query, params...).
		Order("nomor_nota ASC").Find(&sisaList).Error
		
	return sisaList, err
}

func (ServisanModel) TeknisiReport(form forms.ServisanTeknisiReportForm) ([]TeknisiReportItem, error) {
	var (
		teknisiReportItemList []TeknisiReportItem
		query        		  string
		params       		  []interface{}
	)

	if !form.MinDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pengambilan) >= ?"

		params = append(params, form.MinDate.Format(utils.DateOnly))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}
		
		query += "DATE(tanggal_pengambilan) <= ?"

		params = append(params, form.MaxDate.Format(utils.DateOnly))
	}

	err := db.GetDB().
		Model(&Servisan{}).
		Select("servisan.nomor_nota, servisan.tanggal_pengambilan, servisan.tipe_hp, servisan.kerusakan, servisan.total_biaya as biaya, servisan.harga_sparepart, servisan.laba_rugi, teknisi.nama as teknisi").
		Where("status::TEXT LIKE '%Sudah diambil%'").
		Where("id_teknisi = ?", form.IDTeknisi).
		Where(query, params...).
		Joins("LEFT JOIN teknisi ON servisan.id_teknisi = teknisi.id").
		Order("nomor_nota ASC").Find(&teknisiReportItemList).Error

	if err != nil {
		return teknisiReportItemList, err
	}

	return teknisiReportItemList, nil
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

	var oldStatus []utils.StatusServisan

	res := db.GetDB().Model(&Servisan{}).Where("nomor_nota = ?", nomorNota).Pluck("status", &oldStatus)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	if form.Status != nil && oldStatus[0] != *form.Status {
		newServisan["tanggal_pengambilan"] = getTanggalPengambilan(*form.Status)
	}

	res = db.GetDB().Model(&Servisan{NomorNota: nomorNota}).Updates(newServisan)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 && len(newServisan) > 0 {
		return fmt.Errorf("tidak ada servisan dengan nomor nota %d", nomorNota)
	}

	return nil
}

func (ServisanModel) Delete(nomorNota int) error {
	db := db.GetDB()

	if err := db.Delete(&Sparepart{}, "nomor_nota = ?", nomorNota).Error; err != nil {
		return err
	}

	res := db.Delete(&Servisan{}, nomorNota)

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

func getTanggalPengambilan(status utils.StatusServisan) utils.NullTime {
	if status == utils.StatusServisanJadiSudahDiambil || status == utils.StatusServisanTdkJadiSudahDiambil {
		return utils.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	return utils.NullTime{}
}
