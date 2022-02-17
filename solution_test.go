package square

import "testing"

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen  float64
		sidesNum shapeType
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "triangle",
			args: args{
				sidesNum: SidesTriangle,
				sideLen:  1.0,
			},
			want: 0.4330127018922193,
		},
		{
			name: "triangle 2",
			args: args{
				sidesNum: SidesTriangle,
				sideLen:  2.0,
			},
			want: 1.7320508075688772,
		},
		{
			name: "triangle 3",
			args: args{
				sidesNum: SidesTriangle,
				sideLen:  3.0,
			},
			want: 3.8971143170299736,
		},
		{
			name: "circle 1",
			args: args{
				sidesNum: SidesCircle,
				sideLen:  1.0,
			},
			want: 3.141592653589793,
		},
		{
			name: "circle 4",
			args: args{
				sidesNum: SidesCircle,
				sideLen:  4.0,
			},
			want: 50.26548245743669,
		},
		{
			name: "circle 2.3",
			args: args{
				sidesNum: SidesCircle,
				sideLen:  2.3,
			},
			want: 16.619025137490002,
		},
		{
			name: "square 2.3",
			args: args{
				sidesNum: SidesSquare,
				sideLen:  6.785,
			},
			want: 6.785 * 6.785,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSquare(tt.args.sideLen, tt.args.sidesNum); got != tt.want {
				t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
