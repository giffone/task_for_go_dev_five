package service

import (
	"nbrates/internal/domain"
	"testing"
	"time"
)

func TestDTO(t *testing.T) {
	date := time.Date(2021, 04, 15, 0, 0, 0, 0, time.Local)
	// test 0 - null
	items := []domain.Item{}
	wantInt := 0

	_, gotInt := dto(items, date)

	if gotInt != wantInt {
		t.Fatalf("test 0: want %d but got %d", wantInt, gotInt)
	}

	// test 1 - SUCCESS

	items = []domain.Item{
		{
			Fullname:    "АВСТРАЛИЙСКИЙ ДОЛЛАР",
			Title:       "AUD",
			Description: "331.35",
		},
		{
			Fullname:    "АЗЕРБАЙДЖАНСКИЙ МАНАТ",
			Title:       "AZN",
			Description: "254.57",
		},
	}

	wantInt = 2
	wantDTO := []domain.ItemDTO{
		{
			Title: "АВСТРАЛИЙСКИЙ ДОЛЛАР",
			Code:  "AUD",
			Value: 331.35,
			Date:  date,
		},
		{
			Title: "АЗЕРБАЙДЖАНСКИЙ МАНАТ",
			Code:  "AZN",
			Value: 254.57,
			Date:  date,
		},
	}

	gotDTO, gotInt := dto(items, date)

	if gotInt != wantInt {
		t.Fatalf("test 1: want %d but got %d", wantInt, gotInt)
	}

	if len(gotDTO) != len(wantDTO) {
		t.Fatalf("test 1: want  length %d but got length %d", wantInt, gotInt)
	}

	for i, v := range gotDTO {
		if v.Code != wantDTO[i].Code || v.Value != wantDTO[i].Value || v.Date != wantDTO[i].Date {
			t.Fatalf("test 1: want %v but got %v", wantDTO, gotDTO)
		}
	}

	// test 2 - 1 wrong value

	items = []domain.Item{
		{
			Fullname:    "АВСТРАЛИЙСКИЙ ДОЛЛАР",
			Title:       "AUD",
			Description: "331.35",
		},
		{
			Fullname:    "АЗЕРБАЙДЖАНСКИЙ МАНАТ",
			Title:       "AZN",
			Description: "254.57",
		},
		{
			Fullname:    "АРМЯНСКИЙ ДРАМ",
			Title:       "AMD",
			Description: "8.3a", // WRONG VALUE
		},
		{
			Fullname:    "БЕЛОРУССКИЙ РУБЛЬ",
			Title:       "BYN",
			Description: "165.34",
		},
	}

	wantInt = 3
	wantDTO = []domain.ItemDTO{
		{
			Title: "АВСТРАЛИЙСКИЙ ДОЛЛАР",
			Code:  "AUD",
			Value: 331.35,
			Date:  date,
		},
		{
			Title: "АЗЕРБАЙДЖАНСКИЙ МАНАТ",
			Code:  "AZN",
			Value: 254.57,
			Date:  date,
		},
		{
			Title: "БЕЛОРУССКИЙ РУБЛЬ",
			Code:  "BYN",
			Value: 165.34,
			Date:  date,
		},
	}

	gotDTO, gotInt = dto(items, date)

	if gotInt != wantInt {
		t.Fatalf("test 1: want %d but got %d", wantInt, gotInt)
	}

	if len(gotDTO) != len(wantDTO) {
		t.Fatalf("test 1: want  length %d but got length %d", wantInt, gotInt)
	}

	for i, v := range gotDTO {
		if v.Code != wantDTO[i].Code || v.Value != wantDTO[i].Value || v.Date != wantDTO[i].Date {
			t.Fatalf("test 1: want %v but got %v", wantDTO, gotDTO)
		}
	}
}
