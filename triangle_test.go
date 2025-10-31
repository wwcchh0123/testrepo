package main

import (
	"testing"
)

func TestCalculateTriangleArea(t *testing.T) {
	tests := []struct {
		name    string
		base    float64
		height  float64
		want    float64
		wantErr bool
	}{
		{
			name:    "Valid triangle",
			base:    10.0,
			height:  5.0,
			want:    25.0,
			wantErr: false,
		},
		{
			name:    "Another valid triangle",
			base:    6.0,
			height:  8.0,
			want:    24.0,
			wantErr: false,
		},
		{
			name:    "Zero base",
			base:    0,
			height:  5.0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Negative base",
			base:    -10.0,
			height:  5.0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Zero height",
			base:    10.0,
			height:  0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Negative height",
			base:    10.0,
			height:  -5.0,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTriangleArea(tt.base, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTriangleArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
