package singleton

/*Singleton interface*/
type Singleton interface{
	AddOne() int	
}

type singleton struct {
	count int
}

var (
	instance *singleton
)

func init(){
	instance = &singleton{ count: 100 }
}

/*Singleton func return object*/
func GetInstance() Singleton {
	return instance
}

func (s * singleton) AddOne() int {
	s.count++
	return s.count
}