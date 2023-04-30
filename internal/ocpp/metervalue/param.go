package metervalue

import (
	"strconv"
	"strings"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/util"
)

func createMeterValueParams(chargePointID int64, connectorID int32, transactionID *int64, meteredValue MeteredValue, sampledValue SampledValue) db.CreateMeterValueParams {
	format, rawValue, signedValue := defaultFormat(sampledValue.Format, sampledValue.Value)
	measurand := defaultMeasurand(sampledValue.Measurand)

	return db.CreateMeterValueParams{
		ChargePointID:   chargePointID,
		ConnectorID:     connectorID,
		TransactionID:   util.SqlNullInt64(transactionID),
		Format:          format,
		Context:         defaultContext(sampledValue.Context),
		Measurand:       measurand,
		Phase:           defaultPhase(sampledValue.Phase),
		Location:        defaultLocation(sampledValue.Location),
		Unit:            defaultUnit(measurand, sampledValue.Unit),
		RawValue:        util.SqlNullFloat64(rawValue),
		SignedDataValue: util.SqlNullString(signedValue),
		Timestamp:       meteredValue.Timestamp.Time(),
		CreatedAt:       time.Now(),
	}
}

func defaultContext(context *db.MeterReadingContext) db.MeterReadingContext {
	if context == nil {
		return db.MeterReadingContextSamplePeriodic
	}

	return *context
}

func defaultFormat(format *db.MeterValueFormat, value string) (db.MeterValueFormat, *float64, *string) {
	rawValue, err := strconv.ParseFloat(value, 64)

	if (format == nil || *format == db.MeterValueFormatRaw) && err == nil {
		return db.MeterValueFormatRaw, &rawValue, nil
	}

	return db.MeterValueFormatRaw, nil, &value
}

func defaultLocation(location *db.MeterLocation) db.MeterLocation {
	if location == nil {
		return db.MeterLocationOutlet
	}

	return *location
}

func defaultMeasurand(measurand *db.MeterMeasurand) db.MeterMeasurand {
	if measurand == nil {
		return db.MeterMeasurandEnergyActiveImportRegister
	}

	return *measurand
}

func defaultPhase(phase *db.MeterPhase) db.NullMeterPhase {
	nullPhase := db.NullMeterPhase{}

	if phase == nil {
		nullPhase.Scan(nil)
	} else {
		nullPhase.Scan(phase)
	}

	return nullPhase
}

func defaultUnit(measurand db.MeterMeasurand, unit *db.MeterUnitOfMeasure) db.NullMeterUnitOfMeasure {
	nullUnit := db.NullMeterUnitOfMeasure{}

	if unit == nil {
		if strings.Contains(string(measurand), "Energy") {
			nullUnit.Scan(string(db.MeterUnitOfMeasureWh))
		} else {
			nullUnit.Scan(nil)
		}
	} else {
		nullUnit.Scan(unit)
	}

	return nullUnit
}
