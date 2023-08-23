package forms

import (
	"time"

	"github.com/JosephJoshua/shin-psmapi/utils"
)

type GetAllServisanForm struct {
	SearchQuery 	 string                 `form:"q"`
	SearchBy    	 utils.ServisanSearchBy `form:"by" binding:"servisan_search_by_col"`
	MinDate     	 time.Time              `form:"min_date"`
	MaxDate     	 time.Time              `form:"max_date"`
}

type ServisanLabaRugiReportForm struct {
	MinDate time.Time `form:"min_date"`
	MaxDate time.Time `form:"max_date"`
}

type ServisanSisaReportForm struct {
	MinDate time.Time `form:"min_date"`
	MaxDate time.Time `form:"max_date"`
}

type ServisanTeknisiReportForm struct {
	MinDate    time.Time `form:"min_date"`
	MaxDate    time.Time `form:"max_date"`
	IDTeknisi  int 	  	 `form:"id_teknisi"`
}

type CreateServisanForm struct {
	NamaPelanggan     string               `json:"nama_pelanggan" binding:"required,max=256"`
	NoHp              string               `json:"no_hp" binding:"max=32"`
	TipeHp            string               `json:"tipe_hp" binding:"required,max=128"`
	Imei              string               `json:"imei" binding:"max=16"`
	KondisiHp         string               `json:"kondisi_hp" binding:"max=512"`
	Kerusakan         string               `json:"kerusakan" binding:"required,max=256"`
	YangBlmDicek      string               `json:"yang_blm_dicek" binding:"max=128"`
	Kelengkapan       string               `json:"kelengkapan" binding:"max=128"`
	Warna             string               `json:"warna" binding:"max=128"`
	KataSandiPola     string               `json:"kata_sandi_pola" binding:"max=128"`
	IDTeknisi         int                  `json:"id_teknisi" binding:"required"`
	IDSales           int                  `json:"id_sales" binding:"required"`
	Status            utils.StatusServisan `json:"status" binding:"required,status_servisan"`
	TanggalKonfirmasi time.Time            `json:"tanggal_konfirmasi"`
	IsiKonfirmasi     string               `json:"isi_konfirmasi" binding:"max=512"`
	Biaya             float64              `json:"biaya"`
	Diskon            int                  `json:"diskon" binding:"min=0,max=100"`
	DP                float64              `json:"dp"`
	TambahanBiaya     float64              `json:"tambahan_biaya"`
}

type UpdateServisanForm struct {
	NamaPelanggan     *string               `json:"nama_pelanggan" binding:"omitempty,max=256"`
	NoHp              *string               `json:"no_hp" binding:"omitempty,max=32"`
	TipeHp            *string               `json:"tipe_hp" binding:"omitempty,max=128"`
	Imei              *string               `json:"imei" binding:"omitempty,max=16"`
	KondisiHp         *string               `json:"kondisi_hp" binding:"omitempty,max=512"`
	Kerusakan         *string               `json:"kerusakan" binding:"omitempty,max=256"`
	YangBlmDicek      *string               `json:"yang_blm_dicek" binding:"omitempty,max=128"`
	Kelengkapan       *string               `json:"kelengkapan" binding:"omitempty,max=128"`
	Warna             *string               `json:"warna" binding:"omitempty,max=128"`
	KataSandiPola     *string               `json:"kata_sandi_pola" binding:"omitempty,max=128"`
	IDTeknisi         *int                  `json:"id_teknisi" binding:"omitempty"`
	IDSales           *int                  `json:"id_sales" binding:"omitempty"`
	Status            *utils.StatusServisan `json:"status" binding:"omitempty,status_servisan"`
	TanggalKonfirmasi time.Time             `json:"tanggal_konfirmasi"`
	IsiKonfirmasi     *string               `json:"isi_konfirmasi" binding:"omitempty,max=512"`
	Biaya             *float64              `json:"biaya" binding:"omitempty"`
	Diskon            *int                  `json:"diskon" binding:"omitempty,min=0,max=100"`
	DP                *float64              `json:"dp" binding:"omitempty"`
	TambahanBiaya     *float64              `json:"tambahan_biaya" binding:"omitempty"`
}
