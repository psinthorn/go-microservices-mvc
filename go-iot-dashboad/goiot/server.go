package main

type server struct{}

func (*server) TempSensor(req *sensorpb.SensorRequest, stream sensorpb.TempSensor) {

}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer(
	sensorpb.RegisterSensorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
