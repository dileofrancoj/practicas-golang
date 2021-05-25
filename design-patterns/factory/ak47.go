package main

type ak47 struct {
	gun
}

func newMusket() iGun {
	return &ak47{
		gun : gun{
			name : "Ak47 gun",
			power : 1,
		},
	}
}