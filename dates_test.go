package dateinterval

import "testing"

func Test_LeapYear(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1900",
			args: args{1900},
			want: false,
		},
		{
			name: "1992",
			args: args{1992},
			want: true,
		},
		{
			name: "2000",
			args: args{2000},
			want: true,
		},
		{
			name: "2400",
			args: args{2400},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeapYear(tt.args.y); got != tt.want {
				t.Errorf("leapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DaysSoFar(t *testing.T) {
	type args struct {
		d int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1/1",
			args: args{1, 1},
			want: 1,
		},
		{
			name: "1/2",
			args: args{1, 2},
			want: 32,
		},
		{
			name: "1/3",
			args: args{1, 3},
			want: 60,
		},
		{
			name: "31/12",
			args: args{31, 12},
			want: 365,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysSoFar(tt.args.d, tt.args.m); got != tt.want {
				t.Errorf("daysSoFar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LeapYears(t *testing.T) {
	type args struct {
		ys int
		ye int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "same year, not leap year",
			args: args{1900, 1900},
			want: 0,
		},
		{
			name: "same year, leap year",
			args: args{2008, 2008},
			want: 1,
		},
		{
			name: "two leap years",
			args: args{2004, 2008},
			want: 2,
		},
		{
			name: "two leap years in the range",
			args: args{2003, 2009},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeapYears(tt.args.ys, tt.args.ye); got != tt.want {
				t.Errorf("leapYears() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Interval(t *testing.T) {
	type fields struct {
		Year  int
		Month int
		Day   int
	}
	type args struct {
		o Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "same dates",
			fields: fields{1983, 6, 2},
			args:   args{Date{1983, 6, 2}},
			want:   0,
		},
		{
			name:   "same month",
			fields: fields{1983, 6, 2},
			args:   args{Date{1983, 6, 22}},
			want:   19,
		},
		{
			name:   "same month, out of order",
			fields: fields{1983, 6, 22},
			args:   args{Date{1983, 6, 2}},
			want:   19,
		},
		{
			name:   "same year",
			fields: fields{1984, 7, 4},
			args:   args{Date{1984, 12, 25}},
			want:   173,
		},
		{
			name:   "same year, out of order",
			fields: fields{1984, 7, 4},
			args:   args{Date{1984, 12, 25}},
			want:   173,
		},
		{
			name:   "in a leap year",
			fields: fields{1984, 1, 31},
			args:   args{Date{1984, 3, 1}},
			want:   29,
		},
		{
			name:   "in a leap year, out of order",
			fields: fields{1984, 3, 1},
			args:   args{Date{1984, 1, 31}},
			want:   29,
		},
		{
			name:   "same day, different years",
			fields: fields{1984, 7, 4},
			args:   args{Date{1985, 7, 4}},
			want:   364,
		},
		{
			name:   "same day, different years, with a leap year",
			fields: fields{1983, 7, 4},
			args:   args{Date{1985, 7, 4}},
			want:   730,
		},
		{
			name:   "out of order, the same day and month, one year appart, leap year",
			fields: fields{1984, 8, 3},
			args:   args{Date{1983, 8, 3}},
			want:   365,
		},
		{
			name:   "out of order, multiple year apart",
			fields: fields{1989, 1, 3},
			args:   args{Date{1983, 8, 3}},
			want:   1979,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.Interval(tt.args.o); got != tt.want {
				t.Errorf("Date.Interval() = %v, want %v", got, tt.want)
			}
		})
	}
}
