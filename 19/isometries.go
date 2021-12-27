package main

//type isometry [3][3]int
//
//func apply(s scanner, iso isometry) scanner {
//	pos := s.pos
//	b := make([]point, len(s.beacons))
//	for i := 0; i < len(s.beacons); i++ {
//		b[i] = matrixMult(iso, s.beacons[i])
//	}
//	return scanner{
//		id:      s.id,
//		pos:     pos,
//		beacons: b,
//	}
//}
//
//func matrixMult(iso isometry, p point) point {
//	in := [3]int{p.x, p.y, p.z}
//	out := [3]int{0, 0, 0}
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			out[i] += iso[i][j] * in[j]
//		}
//	}
//	return point{x: out[0], y: out[1], z: out[2]}
//}
//
//var isometries = []isometry{
//	[3][3]int{
//		{1, 0, 0},
//		{0, 0, -1},
//		{0, 1, 0},
//	},
//	[3][3]int{
//		{1, 0, 0},
//		{0, -1, 0},
//		{0, 0, -1},
//	},
//	[3][3]int{
//		{1, 0, 0},
//		{0, 0, 1},
//		{0, -1, 0},
//	},
//	[3][3]int{
//		{-1, 0, 0},
//		{0, 0, -1},
//		{0, 1, 0},
//	},
//	[3][3]int{
//		{-1, 0, 0},
//		{0, -1, 0},
//		{0, 0, -1},
//	},
//	[3][3]int{
//		{-1, 0, 0},
//		{0, 0, 1},
//		{0, -1, 0},
//	},
//}
