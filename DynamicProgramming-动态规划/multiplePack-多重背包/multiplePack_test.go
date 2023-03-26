package multiplePack

import "testing"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/26
 * Time: 20:41
 * Description: No Description
 */

func Test_multiplePack(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		nums      []int
		bagWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				weight:    []int{1, 3, 4},
				value:     []int{15, 20, 30},
				nums:      []int{2, 3, 2},
				bagWeight: 10,
			},
			want: 90,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplePack(tt.args.weight, tt.args.value, tt.args.nums, tt.args.bagWeight); got != tt.want {
				t.Errorf("multiplePack() = %v, want: %v", got, tt.want)
			}
		})
	}
}
