package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseEvenLengthGroups(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		want  *ListNode
	}{
		{
			name:  "Empty list",
			input: nil,
			want:  nil,
		},
		{
			name:  "Single node",
			input: &ListNode{Val: 1},
			want:  &ListNode{Val: 1},
		},
		{
			name: "Even length list",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: 7,
										Next: &ListNode{
											Val: 8,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 10,
													Next: &ListNode{
														Val: 11,
														Next: &ListNode{
															Val: 12,
															Next: &ListNode{
																Val: 13,
																Next: &ListNode{
																	Val: 14,
																	Next: &ListNode{
																		Val: 15,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: 10,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 8,
												Next: &ListNode{
													Val: 7,
													Next: &ListNode{
														Val: 11,
														Next: &ListNode{
															Val: 12,
															Next: &ListNode{
																Val: 13,
																Next: &ListNode{
																	Val: 14,
																	Next: &ListNode{
																		Val: 15,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseEvenLengthGroups(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseEvenLengthGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
