package cache

import (
	"testing"
)

func TestLRUCache_Insert(t *testing.T) {
	type fields struct {
		cap int
	}
	type args struct {
		s int
		t int
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
					s: 10,
					t: 13,
				},
				{
					s: 8,
					t: 5,
				},
				{
					s: 9,
					t: 10,
				},
				{
					s: 2,
					t: 19,
				},
				{
					s: 17,
					t: 5,
				},
				{
					s: 9,
					t: 25,
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				cap: -1,
			},
			args: []args{
				// {
				// 	s: 10,
				// 	t: 13,
				// },
				// {
				// 	s: 8,
				// 	t: 5,
				// },
				// {
				// 	s: 9,
				// 	t: 10,
				// },
				// {
				// 	s: 2,
				// 	t: 19,
				// },
				// {
				// 	s: 17,
				// 	t: 5,
				// },
				// {
				// 	s: 9,
				// 	t: 25,
				// },
			},
		},
		// {
		// 	name: "test2",
		// 	fields: fields{
		// 		cap: 6,
		// 	},
		// 	args: args{

		// 	},
		// },
		// {
		// 	name: "test3",
		// 	fields: fields{
		// 		cap: 6,
		// 	},
		// 	args: args{
		// 		s: 8,
		// 		t: 5,
		// 	},
		// },
		// {
		// 	name: "test4",
		// 	fields: fields{
		// 		cap: 6,
		// 	},
		// 	args: args{
		// 		s: 9,
		// 		t: 10,
		// 	},
		// },
		// {
		// 	name: "test5",
		// 	fields: fields{
		// 		cap: 6,
		// 	},
		// 	args: args{
		// 		s: 2,
		// 		t: 19,
		// 	},
		// },
		// {
		// 	name: "test6",
		// 	fields: fields{
		// 		cap: 6,
		// 	},
		// 	args: args{
		// 		s: 17,
		// 		t: 5,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewLRUCache(tt.fields.cap)
			for _, ele := range tt.args {
				c.Insert(ele.s, ele.t)
			}
		})
	}
}
func TestLRUCache_Get(t *testing.T) {
	type fields struct {
		cap int
	}
	type args struct {
		s int
	}

	tests := []struct {
		name   string
		fields fields
		args   []args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				cap: 3,
			},
			args: []args{
				{
					s: 10,
					t: 13,
				},
				{
					s: 8,
					t: 5,
				},
				{
					s: 9,
					t: 10,
				},
				{
					s: 2,
					t: 19,
				},
				{
					s: 17,
					t: 5,
				},
				{
					s: 9,
					t: 25,
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				cap: 6,
			},
			args: args{
				s: 9,
			},
			want: 10,
		},
		{
			name: "test3",
			fields: fields{
				cap: 6,
			},
			args: args{
				s: 2,
			},
			want: 19,
		},
		{
			name: "test4",
			fields: fields{
				cap: 6,
			},
			args: args{
				s: 67,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewLRUCache(tt.fields.cap)

			//insert
			c.Insert(tt.args.s, tt.want)

			// Get and verify code
			if got := c.Get(tt.args.s); got != tt.want {
				t.Errorf("LRUCache.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
