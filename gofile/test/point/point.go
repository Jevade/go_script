package point
import
	(
		"math"
	)
type Point2D struct{
	X float64
	Y float64
}

type Point3D struct{
	Z float64
    Point2D
}
func (T *Point2D) Get(x,y float64)(float64,float64){
	return T.X,T.Y
}
func (T *Point2D) Set(x,y float64){
	T.X = x
	T.Y = y
}

func (T *Point2D) ABS2D()float64{
	return math.Sqrt(T.X * T.X + T.Y * T.Y)
}

func (T *Point2D) Dis(T2 *Point2D)float64{
	return math.Sqrt((T.X - T2.X)*(T.X - T2.X) + (T.Y - T2.Y)*(T.Y - T2.Y)) 
}

func (T *Point2D) Scale(scale float64){
	T.X *= scale
	T.Y *= scale
}

func(T *Point3D)ABS()float64{
	return math.Sqrt(T.X*T.X+T.Y*T.Y+T.Z*T.Z)
}