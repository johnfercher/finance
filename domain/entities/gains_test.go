package entities_test

import (
	"finance/m/v2/domain/consts/gain"
	"finance/m/v2/domain/consts/month"
	"finance/m/v2/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGain(t *testing.T) {
	// Arrange
	recurrence := []month.Month{"January", "March"}

	// Act
	createdGain := entities.NewGain("label", 1000.0, gain.Taxable, recurrence)

	// Assert
	assert.Equal(t, "label", createdGain.Label)
	assert.Equal(t, 1000.0, createdGain.Value)
	assert.Equal(t, gain.Taxable, createdGain.Type)
	assert.Equal(t, month.BuildRecurrence(recurrence), createdGain.Recurrence)
}

func TestGains_AddTaxableGain_WhenThereIsOnlyThisType(t *testing.T) {
	// Arrange
	recurrence := []month.Month{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	gains := entities.NewGains()

	// Act
	gains.AddTaxableGain("label", 1000.0, recurrence)

	// Assert
	for _, month := range recurrence {
		assert.Equal(t, 1000.0, gains.GetTaxableTotal(month))
		assert.Equal(t, 0.0, gains.GetNonTaxableTotal(month))
		assert.Equal(t, 725.0, gains.GetTaxableLiquid(month))
		assert.Equal(t, 725.0, gains.GetTotalLiquid(month))
	}
}

func TestGains_AddNonTaxableGain_WhenThereIsOnlyThisType(t *testing.T) {
	// Arrange
	recurrence := []month.Month{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	gains := entities.NewGains()

	// Act
	gains.AddNonTaxableGain("label", 1000.0, recurrence)

	// Assert
	for _, month := range recurrence {
		assert.Equal(t, 1000.0, gains.GetNonTaxableTotal(month))
		assert.Equal(t, 0.0, gains.GetTaxableTotal(month))
		assert.Equal(t, 0.0, gains.GetTaxableLiquid(month))
		assert.Equal(t, 1000.0, gains.GetTotalLiquid(month))
	}
}

func TestGains_GetTotalLiquid(t *testing.T) {
	// Arrange
	recurrenceTaxable := []month.Month{"February", "March", "April", "May", "June", "July"}
	recurrenceNonTaxable := []month.Month{"June", "July", "August", "September", "October", "November"}

	gains := entities.NewGains()

	// Act
	gains.AddTaxableGain("label", 1000.0, recurrenceTaxable)
	gains.AddNonTaxableGain("label", 500.0, recurrenceNonTaxable)

	// Assert
	assert.Equal(t, 0.0, gains.GetTotalLiquid(month.January))
	assert.Equal(t, 725.0, gains.GetTotalLiquid(month.February))
	assert.Equal(t, 725.0, gains.GetTotalLiquid(month.March))
	assert.Equal(t, 725.0, gains.GetTotalLiquid(month.April))
	assert.Equal(t, 725.0, gains.GetTotalLiquid(month.May))
	assert.Equal(t, 1225.0, gains.GetTotalLiquid(month.June))
	assert.Equal(t, 1225.0, gains.GetTotalLiquid(month.July))
	assert.Equal(t, 500.0, gains.GetTotalLiquid(month.August))
	assert.Equal(t, 500.0, gains.GetTotalLiquid(month.September))
	assert.Equal(t, 500.0, gains.GetTotalLiquid(month.October))
	assert.Equal(t, 500.0, gains.GetTotalLiquid(month.November))
	assert.Equal(t, 0.0, gains.GetTotalLiquid(month.December))
}
