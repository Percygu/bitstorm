package service

import (
	"bitstorm/internal/pkg/constant"
	"bitstorm/internal/pkg/middlewares/gormcli"
	"bitstorm/internal/repo"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// 需要什么功能，抽象不同的Service

type PrizeService interface {
	GetPrizeList() ([]*ViewPrize, error)
	GetPrize(id uint) (*ViewPrize, error)
	GetPrizeMap() (map[string]*ViewPrize, error)
}

type prizeService struct {
	prizeReop *repo.PrizeRepo
}

func NewPrizeService() PrizeService {
	return &prizeService{
		prizeReop: repo.NewPrizeRepo(),
	}
}

func (p *prizeService) GetPrizeList() ([]*ViewPrize, error) {
	db := gormcli.GetDB()
	list, err := p.prizeReop.GetAll(db)
	if err != nil {
		log.Errorf("GetPrizeList err:%v", err)
		return nil, fmt.Errorf("prizeService|GetAll: %v", err)
	}
	prizeList := make([]*ViewPrize, 0)
	for _, prize := range list {
		if prize.SysStatus != constant.PrizeStatusNormal {
			continue
		}
		prizeList = append(prizeList, &ViewPrize{
			Id:        prize.Id,
			Title:     prize.Title,
			Img:       prize.Img,
			PrizeNum:  prize.PrizeNum,
			LeftNum:   prize.LeftNum,
			PrizeType: prize.PrizeType,
		})
	}
	return prizeList, nil
}

func (p *prizeService) GetPrize(id uint) (*ViewPrize, error) {
	prizeModel, err := p.prizeReop.Get(gormcli.GetDB(), id)
	if err != nil {
		return nil, fmt.Errorf("adminService|GetPrize:%v", err)
	}
	prize := &ViewPrize{
		Id:        prizeModel.Id,
		Title:     prizeModel.Title,
		Img:       prizeModel.Img,
		PrizeNum:  prizeModel.PrizeNum,
		LeftNum:   prizeModel.LeftNum,
		PrizeType: prizeModel.PrizeType,
	}
	return prize, nil
}

func (p *prizeService) GetPrizeMap() (map[string]*ViewPrize, error) {
	db := gormcli.GetDB()
	list, err := p.prizeReop.GetAll(db)
	if err != nil {
		log.Errorf("GetPrizeList err:%v", err)
		return nil, fmt.Errorf("prizeService|GetAll: %v", err)
	}
	prizeMap := make(map[string]*ViewPrize)
	for _, prize := range list {
		if prize.SysStatus == constant.PrizeStatusDelete {
			continue
		}
		prizeMap[prize.Title] = &ViewPrize{
			Id:        prize.Id,
			Title:     prize.Title,
			Img:       prize.Img,
			PrizeNum:  prize.PrizeNum,
			LeftNum:   prize.LeftNum,
			PrizeType: prize.PrizeType,
		}
	}
	return prizeMap, nil
}

func (a *prizeService) AddPrize(viewPrize *ViewPrize) error {
	return nil
}
