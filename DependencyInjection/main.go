package main

import "fmt"

// The Service interface defines the behavior of a service
// that can be injected into other components.
type Service interface {
    	DoSomething() string
}

// The RealService struct is an 
// implementation of the Service interface.
type RealService struct{}

// DoSomething is a method of the RealService struct.
// It returns a string indicating that the real service is being used.
func (s *RealService) DoSomething() string {
    	return "Using real service"
}

// The FakeService struct is another 
// implementation of the Service interface.
// It can be used as a fake service for testing purposes.
type FakeService struct{}

// DoSomething is a method of the FakeService struct.
// It returns a string indicating that the fake service is being used.
func (s *FakeService) DoSomething() string {
    	return "Using fake service"
}

// The Controller struct depends on the Service interface.
// It has a field of type Service that will be 
// injected with a concrete implementation.
type Controller struct {
    	Service Service
}

// The DoSomething method of the Controller 
// struct uses the Service interface
// to perform some operation and return 
// the result as a string.
func (c *Controller) DoSomething() string {
    	return c.Service.DoSomething()
}

func main() {
	// Create a new instance of the Controller struct.
	// Inject the RealService as the service 
	// to be used by the controller.
	controller := &Controller{Service: &RealService{}}

	// Call the DoSomething method of the Controller struct.
	// It will use the injected RealService to 
	// perform the operation.
	result := controller.DoSomething()
	fmt.Println(result) // Output: "Using real service"

	// Inject the FakeService as the service 
	// to be used by the controller.
	controller.Service = &FakeService{}

	// Call the DoSomething method of the Controller struct again.
	// This time, it will use the injected FakeService 
	// to perform the operation.
	result = controller.DoSomething()
	fmt.Println(result) // Using fake service
}
