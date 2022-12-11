package acuvim

import (
	"errors"
	"time"

	"github.com/simonvetter/modbus"
)

func Polling(id uint8, port string) (payload Payload, err error) {
	url := "rtu://" + port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    9600,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  1 * time.Second,
	})
	if err != nil {
		err = errors.New("E0")
		return payload, err
	}

	err = client.Open()
	if err != nil {
		err := errors.New("E1")
		return payload, err
	}
	defer client.Close()

	client.SetUnitId(id)
	payload.Powermeter.PM_ID = id

	datas, err := client.ReadFloat32s(16384, 36, modbus.HOLDING_REGISTER)
	datas, err = client.ReadFloat32s(16384, 36, modbus.HOLDING_REGISTER)
	if err != nil {
		err = errors.New("Timeout")
		return payload, err
	}

	for i, data := range datas {
		//log.Printf("[%d] %f", i, data)
		switch i {
		case 0:
			payload.Powermeter.Data.Frequency = data

		case 1:
			payload.Powermeter.Data.Voltage.Phase2Netral.Phase_1 = data
		case 2:
			payload.Powermeter.Data.Voltage.Phase2Netral.Phase_2 = data
		case 3:
			payload.Powermeter.Data.Voltage.Phase2Netral.Phase_3 = data
		case 4:
			payload.Powermeter.Data.Voltage.Phase2Netral.Avg = data

		case 5:
			payload.Powermeter.Data.Voltage.Phase2Phase.L1_L2 = data
		case 6:
			payload.Powermeter.Data.Voltage.Phase2Phase.L2_L3 = data
		case 7:
			payload.Powermeter.Data.Voltage.Phase2Phase.L3_L1 = data
		case 8:
			payload.Powermeter.Data.Voltage.Phase2Phase.Avg = data

		case 9:
			payload.Powermeter.Data.Current.Phase_1 = data
		case 10:
			payload.Powermeter.Data.Current.Phase_2 = data
		case 11:
			payload.Powermeter.Data.Current.Phase_3 = data
		case 12:
			payload.Powermeter.Data.Current.Avg = data
		case 13:
			payload.Powermeter.Data.Current.Netral = data

		case 14:
			payload.Powermeter.Data.Power.Active.Phase_1 = data
		case 15:
			payload.Powermeter.Data.Power.Active.Phase_2 = data
		case 16:
			payload.Powermeter.Data.Power.Active.Phase_3 = data
		case 17:
			payload.Powermeter.Data.Power.Active.Total = data

		case 18:
			payload.Powermeter.Data.Power.Reactive.Phase_1 = data
		case 19:
			payload.Powermeter.Data.Power.Reactive.Phase_2 = data
		case 20:
			payload.Powermeter.Data.Power.Reactive.Phase_3 = data
		case 21:
			payload.Powermeter.Data.Power.Reactive.Total = data

		case 22:
			payload.Powermeter.Data.Power.Apparent.Phase_1 = data
		case 23:
			payload.Powermeter.Data.Power.Apparent.Phase_2 = data
		case 24:
			payload.Powermeter.Data.Power.Apparent.Phase_3 = data
		case 25:
			payload.Powermeter.Data.Power.Apparent.Total = data

		case 26:
			payload.Powermeter.Data.PowerFactor.Phase_1 = data
		case 27:
			payload.Powermeter.Data.PowerFactor.Phase_2 = data
		case 28:
			payload.Powermeter.Data.PowerFactor.Phase_3 = data
		case 29:
			payload.Powermeter.Data.PowerFactor.Total = data

		case 30:
			payload.Powermeter.Data.Unbalance.Voltage = data
		case 31:
			payload.Powermeter.Data.Unbalance.Current = data

		case 33:
			payload.Powermeter.Data.PowerDemand.Active = data
		case 34:
			payload.Powermeter.Data.PowerDemand.Reactive = data
		case 35:
			payload.Powermeter.Data.PowerDemand.Apparent = data

		}
	}

	return payload, nil
}
