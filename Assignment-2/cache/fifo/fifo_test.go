package cache

import (
	"testing"
)

// test function for fifo insert
func TestFifoCache_Insert(t *testing.T) {
	type fields struct {
		cap int
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				cap: 3,
			},
			args: []args{
				{
					key:   10,
					value: 13,
				},
				{
					key:   8,
					value: 5,
				},
				{
					key:   9,
					value: 10,
				},
				{
					key:   2,
					value: 19,
				},
				{
					key:   17,
					value: 5,
				},
				{
					key:   9,
					value: 25,
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				cap: -1,
			},
			args: []args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewFifoCache(tt.fields.cap)
			for _, ele := range tt.args {
				c.Insert(ele.key, ele.key)
			}
		})
	}
}

//test function for fifo Get

func TestFifoCache_Get(t *testing.T) {
	type fields struct {
		cap int
	}
	type args struct {
		key   int
		value int
	}

	tests := []struct {
		name   string
		fields fields
		args   []args
		want   []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				cap: 3,
			},
			args: []args{
				{
					key:   10,
					value: 13,
				},
				{
					key:   8,
					value: 5,
				},
				{
					key:   9,
					value: 10,
				},
				{
					key:   2,
					value: 19,
				},
				{
					key:   17,
					value: 5,
				},
				{
					key:   9,
					value: 25,
				},
			},
			want: []int{
				-1, -1, 25, 19, 5, 25,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewFifoCache(tt.fields.cap)

			//insert
			for _, ele := range tt.args {
				c.Insert(ele.key, ele.value)
			}

			// Get and verify code
			for i, ele := range tt.args {
				if got := c.Get(ele.key); got != tt.want[i] {
					t.Errorf("FifoCache.get() = %v, want %v", got, tt.want[i])
				}
			}

		})
	}
}
