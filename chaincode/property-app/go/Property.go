package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PropertyTransferSmartContract는 부동산 거래 트랜잭션을 처리하는 스마트 계약입니다.
type PropertyTransferSmartContract struct {
	contractapi.Contract
}

// Property는 부동산 정보를 나타내는 구조체입니다.
type Property struct {
	ID        string `json:"id"`        // 부동산 ID
	Name      string `json:"name"`      // 부동산 이름
	Area      int    `json:"area"`      // 부동산 면적
	OwnerName string `json:"ownerName"` // 현재 소유주 이름
	Value     int    `json:"value"`     // 부동산 가치
}

// AddProperty 함수는 새로운 부동산 정보를 추가하는 메서드입니다.
func (pc *PropertyTransferSmartContract) AddProperty(ctx contractapi.TransactionContextInterface, id string, name string, area int, ownerName string, value int) error {
	// 월드 스테이트에서 부동산 데이터 조회
	propertyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("월드 스테이트에서 데이터를 읽어오지 못했습니다: %s", err)
	}

	// 이미 해당 ID의 부동산이 존재하는 경우 오류 반환
	if propertyJSON != nil {
		return fmt.Errorf("%s 부동산은 이미 존재합니다", id)
	}

	// 새로운 부동산 객체 생성
	prop := Property{
		ID:        id,
		Name:      name,
		Area:      area,
		OwnerName: ownerName,
		Value:     value,
	}

	// 부동산 객체를 JSON으로 직렬화
	propertyBytes, err := json.Marshal(prop)
	if err != nil {
		return err
	}

	// 월드 스테이트에 새로운 부동산 데이터 저장
	err = ctx.GetStub().PutState(id, propertyBytes)
	if err != nil {
		return fmt.Errorf("부동산 데이터를 월드 스테이트에 저장하지 못했습니다: %s", err)
	}

	return nil
}
