package main

import "testing"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f ", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("SALAH! Seharusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f ", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH! Seharusnya %.2f", luasSeharusnya)
	}
}
