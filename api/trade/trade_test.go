package trade

import (
	"testing"
	"time"

	"gitlab.com/katana-labs/assessment-frontend/api/offer"
)

func TestValidateTrade_WhenGivenNoVolume_ShouldReturnErrVolumeTooSmall(t *testing.T) {
	tm := time.Now()
	of := offer.Rand()
	tr := Trade{
		UserID:  "123",
		OfferID: of.ID,
		Volume:  0,
	}
	ba := 0.0

	err := validateTrade(tm, of, tr, ba)
	if err != errVolumeTooSmall {
		t.Errorf("expected %v, got %v", errVolumeTooSmall, err)
	}
}

func TestValidateTrade_WhenGivenTooMuchVolume_ShouldReturnErrVolumeTooLarge(t *testing.T) {
	tm := time.Now()
	of := offer.Rand()
	of.Volume = 10
	tr := Trade{
		UserID:  "123",
		OfferID: of.ID,
		Volume:  20,
	}
	ba := 0.0

	err := validateTrade(tm, of, tr, ba)
	if err != errVolumeTooLarge {
		t.Errorf("expected %v, got %v", errVolumeTooLarge, err)
	}
}

func TestValidateTrade_WhenGivenExpiredOffer_ShouldReturnErrOfferExpired(t *testing.T) {
	now := time.Now()

	tm := now.Add(-10 * time.Second)
	of := offer.Rand()
	of.Volume = 10
	of.CreatedOn = now.Add(-20 * time.Second)
	tr := Trade{
		UserID:  "123",
		OfferID: of.ID,
		Volume:  5,
	}
	ba := 0.0

	err := validateTrade(tm, of, tr, ba)
	if err != errOfferExpired {
		t.Errorf("expected %v, got %v", errOfferExpired, err)
	}
}

func TestValidateTrade_WhenBalanceIsTooLow_ShouldReturnErrInsufficientFunds(t *testing.T) {
	now := time.Now()

	tm := now.Add(-10 * time.Second)
	of := offer.Rand()
	of.Volume = 10
	of.UnitPrice = 1.0
	of.CreatedOn = now.Add(-5 * time.Second)
	tr := Trade{
		UserID:  "123",
		OfferID: of.ID,
		Volume:  5,
	}
	ba := 4.0

	err := validateTrade(tm, of, tr, ba)
	if err != errInsufficientFunds {
		t.Errorf("expected %v, got %v", errInsufficientFunds, err)
	}
}

func TestValidateTrade_WhenGivenValidTrade_ShouldReturnNil(t *testing.T) {
	now := time.Now()

	tm := now.Add(-10 * time.Second)
	of := offer.Rand()
	of.Volume = 10
	of.UnitPrice = 1.0
	of.CreatedOn = now.Add(-5 * time.Second)
	tr := Trade{
		UserID:  "123",
		OfferID: of.ID,
		Volume:  5,
	}
	ba := 5.0

	err := validateTrade(tm, of, tr, ba)
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}
