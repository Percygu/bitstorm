package repo

import (
	"bitstorm/internal/model"
	"bitstorm/internal/pkg/middlewares/cache"
	"bitstorm/internal/pkg/middlewares/log"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type PrizeRepo struct {
}

func NewPrizeRepo() *PrizeRepo {
	return &PrizeRepo{}
}

func (r *PrizeRepo) Get(db *gorm.DB, id uint) (*model.Prize, error) {
	// 优先从缓存获取
	prize, err := r.GetFromCache(id)
	if err == nil && prize != nil {
		return prize, nil
	}
	prize = &model.Prize{
		Id: id,
	}
	err = db.Model(&model.Prize{}).First(prize).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("PrizeRepo|Get:%v", err)
	}
	return prize, nil
}

func (r *PrizeRepo) GetAll(db *gorm.DB) ([]*model.Prize, error) {
	var prizes []*model.Prize
	err := db.Model(&model.Prize{}).Where("").Order("sys_updated desc").Find(&prizes).Error
	if err != nil {
		return nil, fmt.Errorf("PrizeRepo|GetAll:%v", err)
	}
	return prizes, nil
}

func (r *PrizeRepo) CountAll(db *gorm.DB) (int64, error) {
	var num int64
	err := db.Model(&model.Prize{}).Count(&num).Error
	if err != nil {
		return 0, fmt.Errorf("PrizeRepo|CountAll:%v", err)
	}
	return num, nil
}

func (r *PrizeRepo) Create(db *gorm.DB, prize *model.Prize) error {
	err := db.Model(&model.Prize{}).Create(prize).Error
	if err != nil {
		return fmt.Errorf("PrizeRepo|Create:%v", err)
	}
	return nil
}

func (r *PrizeRepo) Delete(db *gorm.DB, id uint) error {
	prize := &model.Prize{Id: id}
	if err := db.Model(&model.Prize{}).Delete(prize).Error; err != nil {
		return fmt.Errorf("PrizeRepo|Delete:%v")
	}
	return nil
}

func (r *PrizeRepo) Update(db *gorm.DB, prize *model.Prize, cols ...string) error {
	var err error
	if len(cols) == 0 {
		err = db.Model(&model.Prize{}).Updates(prize).Error
	} else {
		err = db.Model(&model.Prize{}).Select(cols).Updates(prize).Error
	}
	if err != nil {
		return fmt.Errorf("PrizeRepo|Update:%v", err)
	}
	return nil
}

// GetFromCache 根据id从缓存获取奖品
func (r *PrizeRepo) GetFromCache(id uint) (*model.Prize, error) {
	redisCli := cache.GetRedisCli()
	idStr := strconv.FormatUint(uint64(id), 10)
	ret, exist, err := redisCli.Get(context.Background(), idStr)
	if err != nil {
		log.Errorf("PrizeRepo|GetFromCache:" + err.Error())
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	prize := model.Prize{}
	json.Unmarshal([]byte(ret), &model.Prize{})

	return &prize, nil
}
