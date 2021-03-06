package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type thermostat struct {
	*Accessory

	thermostat         *service.Thermostat
	onTargetTempChange func(float64)
	onTargetModeChange func(model.HeatCoolModeType)
}

// NewThermostat returns a thermostat which implements model.Thermostat.
func NewThermostat(info model.Info, temp, min, max, steps float64) *thermostat {
	accessory := New(info)
	t := service.NewThermostat(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	ts := thermostat{accessory, t, nil, nil}

	t.TargetTemp.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		if ts.onTargetTempChange != nil {
			ts.onTargetTempChange(t.TargetTemp.Temperature())
		}
	})

	t.TargetMode.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		if ts.onTargetModeChange != nil {
			ts.onTargetModeChange(t.TargetMode.HeatingCoolingMode())
		}
	})

	return &ts
}

func (t *thermostat) Temperature() float64 {
	return t.thermostat.Temp.Temperature()
}

func (t *thermostat) SetTemperature(value float64) {
	t.thermostat.Temp.SetTemperature(value)
}

func (t *thermostat) Unit() model.TempUnit {
	return t.thermostat.Unit.Unit()
}

func (t *thermostat) SetTargetTemperature(value float64) {
	t.thermostat.TargetTemp.SetTemperature(value)
}

func (t *thermostat) TargetTemperature() float64 {
	return t.thermostat.TargetTemp.Temperature()
}

func (t *thermostat) SetMode(value model.HeatCoolModeType) {
	if value != model.HeatCoolModeAuto {
		t.thermostat.Mode.SetHeatingCoolingMode(value)
	}
}

func (t *thermostat) Mode() model.HeatCoolModeType {
	return t.thermostat.Mode.HeatingCoolingMode()
}

func (t *thermostat) SetTargetMode(value model.HeatCoolModeType) {
	t.thermostat.TargetMode.SetHeatingCoolingMode(value)
}

func (t *thermostat) TargetMode() model.HeatCoolModeType {
	return t.thermostat.TargetMode.HeatingCoolingMode()
}

func (t *thermostat) OnTargetTempChange(fn func(float64)) {
	t.onTargetTempChange = fn
}

func (t *thermostat) OnTargetModeChange(fn func(model.HeatCoolModeType)) {
	t.onTargetModeChange = fn
}
