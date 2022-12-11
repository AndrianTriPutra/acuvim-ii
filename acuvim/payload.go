package acuvim

type Payload struct {
	Device_ID  string
	Timestamp  string
	Powermeter Powermeter
}

type Powermeter struct {
	PM_ID uint8
	Data  Data
}

type Data struct {
	Frequency   float32
	Voltage     Voltage
	Current     Current
	Power       Power
	PowerFactor Phase2
	Unbalance   Params
	PowerDemand Demands
}

type Voltage struct {
	Phase2Netral Phase1
	Phase2Phase  Phase2Phase
}

type Phase2Phase struct {
	L1_L2 float32
	L2_L3 float32
	L3_L1 float32
	Avg   float32
}

type Phase1 struct {
	Phase_1 float32
	Phase_2 float32
	Phase_3 float32
	Avg     float32
}

type Phase2 struct {
	Phase_1 float32
	Phase_2 float32
	Phase_3 float32
	Total   float32
}

type Current struct {
	Phase_1 float32
	Phase_2 float32
	Phase_3 float32
	Avg     float32
	Netral  float32
}

type Power struct {
	Active   Phase2
	Reactive Phase2
	Apparent Phase2
}

type Params struct {
	Voltage float32
	Current float32
}

type Demands struct {
	Active   float32
	Reactive float32
	Apparent float32
}
