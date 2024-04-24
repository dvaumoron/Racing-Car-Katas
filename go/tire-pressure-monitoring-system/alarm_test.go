package tpms

import "testing"

type sensorMock struct {
	value int
}

func (s sensorMock) popNextPressurePsiValue() int {
	return s.value
}

func TestAlarmLow(t *testing.T) {
	a := NewAlarm(sensorMock{value: 16})

	a.check()

	if !a.alarmOn {
		t.Fatal("should indicate that the value is too low")
	}
}

func TestAlarmLowDouble(t *testing.T) {
	a := NewAlarm(sensorMock{value: 16})

	a.check()
	a.sensor = sensorMock{value: 18}
	a.check()

	if !a.alarmOn {
		t.Fatal("should not change")
	}
}

func TestAlarmNoAlarm(t *testing.T) {
	a := NewAlarm(sensorMock{value: 19})

	a.check()

	if a.alarmOn {
		t.Fatal("should not trigger")
	}
}

func TestAlarmHigh(t *testing.T) {
	a := NewAlarm(sensorMock{value: 22})

	a.check()

	if !a.alarmOn {
		t.Fatal("should indicate that the value is too high")
	}
}
