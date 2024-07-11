package user

import (
	"context"
	"time"

	"open-user/internal/app/helper"

	"go.elastic.co/apm"
)

const (
	tblOnboarding = "tbl_onboarding"
)

type TblOnBoarding struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	UUID              string          `gorm:"column:uuid" json:"uuid"`
	DeviceID          string          `gorm:"column:device_id" json:"device_id"`
	MerchantType      string          `gorm:"column:merchant_type" json:"merchant_type"`
	Status            string          `gorm:"column:status" json:"status"`
	SchedullerStep    int             `gorm:"column:scheduller_step" json:"scheduller_step"`
	ISSchedullerRetry int             `gorm:"column:is_scheduller_retry" json:"is_scheduller_retry"`
	CreatedAt         helper.WrapTime `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         helper.WrapTime `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`
}

func (u *TblOnBoarding) TableName() string {
	return tblOnboarding
}

func (r *UserRepository) InquiryDataOnboardingByUUIDDayNow(ctx context.Context, uuid, device_id string) (TblOnBoarding, error) {
	var result TblOnBoarding
	// start span apm
	span, _ := apm.StartSpan(ctx, "InquiryDataOnboardingByUUIDDayNow", CtxRequestStr)
	defer span.End()

	err := r.Db.Model(&result).Where("uuid = ? AND device_id = ? AND (DATE(created_at) = ? OR status IN (?)) AND status IN (?)", uuid, device_id, time.Now().Format("2006-01-02"), []string{"7", "8"}, []string{"1", "2", "3", "4", "5", "6", "7", "8"}).Order("created_at desc").First(&result).Error

	return result, err
}

func (r *UserRepository) InquiryDataOnboardingByUUIDAndDeviceID(ctx context.Context, uuid, device_id string) (TblOnBoarding, error) {
	var result TblOnBoarding
	// start span apm
	span, _ := apm.StartSpan(ctx, "InquiryDataOnboarding", CtxRequestStr)
	defer span.End()
	err := r.Db.Model(&result).Where("uuid = ? AND device_id = ?", uuid, device_id).Order("created_at desc").First(&result).Error

	return result, err
}

func (r *UserRepository) InquiryDataOnboardingByUUID(ctx context.Context, uuid string) (TblOnBoarding, error) {
	var result TblOnBoarding
	// start span apm
	span, _ := apm.StartSpan(ctx, "InquiryDataOnboardingByUUID", CtxRequestStr)
	defer span.End()
	err := r.Db.Model(&result).Where("uuid = ? ", uuid).Order("created_at desc").First(&result).Error

	return result, err
}

func (r *UserRepository) InsertDataOnboarding(ctx context.Context, param TblOnBoarding) error {
	// start span apm
	span, _ := apm.StartSpan(ctx, "InsertDataOnboarding", CtxRequestStr)
	defer span.End()
	err := r.Db.Create(&param).Error

	return err
}

func (r *UserRepository) UpdateDataOnboarding(ctx context.Context, param TblOnBoarding) error {
	// start span apm
	span, _ := apm.StartSpan(ctx, "UpdateDataOnboarding", CtxRequestStr)
	defer span.End()
	err := r.Db.Save(&param).Error

	return err
}

func (r *UserRepository) UpdateDataOnboardingStatusCompleted(ctx context.Context, nik string) {
	// start span apm
	span, _ := apm.StartSpan(ctx, "UpdateDataOnboardingStatusCompleted", CtxRequestStr)
	defer span.End()
	_ = r.Db.Exec("UPDATE tbl_onboarding INNER JOIN tbl_onboarding_customers ON tbl_onboarding.uuid = tbl_onboarding_customers.uuid SET tbl_onboarding.status = '9' WHERE tbl_onboarding_customers.nik = ? and tbl_onboarding.status = ?", nik, "8")

}
