package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HaveNInCommon(t *testing.T) {
	s1 := scanner{
		id: 0,
		beacons: []point{
			{x: 1, y: 2, z: 3},
			{x: 1, y: 1, z: 3},
			{x: -6, y: 2, z: 30},
			{x: 1, y: 2, z: 1},
		},
	}
	s2 := scanner{
		id: 1,
		beacons: []point{
			{x: 1, y: 2, z: 3},
			{x: 1, y: 1, z: 3},
			{x: -6, y: 2, z: 30},
			{x: 1, y: 2, z: 10},
		},
	}
	assert.True(t, haveNInCommon(s1, s2, 3))
	assert.False(t, haveNInCommon(s1, s2, 4))
}

func Test_Translate(t *testing.T) {
	s1 := scanner{
		beacons: []point{
			{x: 1, y: 2, z: 3},
			{x: 1, y: 1, z: 3},
			{x: -6, y: 2, z: 30},
			{x: 1, y: 2, z: 1},
		},
	}
	s2 := scanner{
		pos: point{x: 1, y: -1, z: 1},
		beacons: []point{
			{x: 2, y: 1, z: 4},
			{x: 2, y: 0, z: 4},
			{x: -5, y: 1, z: 31},
			{x: 2, y: 1, z: 2},
		},
	}
	v := vector{x: 1, y: -1, z: 1}
	moveds1 := translate(s1, v)
	assert.Equal(t, s2, moveds1)
	fmt.Println(s1)
}

func Test_HaveNInCommonWithTranslation(t *testing.T) {
	s1 := scanner{
		id: 0,
		beacons: []point{
			{x: 1, y: 2, z: 3},
			{x: 1, y: 1, z: 3},
			{x: -6, y: 2, z: 30},
			{x: 1, y: 2, z: 1},
		},
	}
	s2 := scanner{
		id:  1,
		pos: point{x: 10, y: 10, z: 10},
		beacons: []point{
			{x: 1, y: 2, z: 3},
			{x: 1, y: 1, z: 3},
			{x: -6, y: 2, z: 30},
			{x: 1, y: 2, z: 10},
		},
	}
	transVector := vector{x: 1, y: 1, z: -1}
	moveds2 := translate(s2, transVector)
	require.Equal(t, point{x: 11, y: 11, z: 9}, moveds2.pos)
	require.NotEqual(t, s2.beacons, moveds2.beacons)
	require.False(t, haveNInCommon(s1, moveds2, 3))
	require.False(t, haveNInCommon(s1, moveds2, 4))
	match, v, newScanner := haveNInCommonWithTranslation(s1, moveds2, 3)
	require.True(t, match)
	require.Equal(t, vector{x: -1, y: -1, z: 1}, v)
	require.Equal(t, s2, newScanner)

	s3 := copy(s2)
	s3.beacons[2].z += 15
	moveds3 := translate(s3, transVector)
	match, v, newScanner = haveNInCommonWithTranslation(s1, moveds3, 3)
	require.False(t, match)
	require.Equal(t, vector{}, v)
	assert.Equal(t, moveds3, newScanner)
}
