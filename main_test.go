package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCharString(t *testing.T) {
	t.Run("Farmer", func(t *testing.T) {
		c := farmer
		got := fmt.Sprintf("%s", c)
		expected := "Farmer"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Goose", func(t *testing.T) {
		c := goose
		got := fmt.Sprintf("%s", c)
		expected := "Goose"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Bean", func(t *testing.T) {
		c := bean
		got := fmt.Sprintf("%s", c)
		expected := "Bean"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Fox", func(t *testing.T) {
		c := fox
		got := fmt.Sprintf("%s", c)
		expected := "Fox"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestCharEats(t *testing.T) {
	t.Run("Fox eats goose when farmer absent", func(t *testing.T) {
		s := state{left, right, right, right}
		got := fox.eats(goose, s)
		expected := true
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Fox doesn't eat goose when farmer present", func(t *testing.T) {
		s := state{right, right, right, right}
		got := fox.eats(goose, s)
		expected := false
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Goose eats bean when farmer absent", func(t *testing.T) {
		s := state{right, left, left, right}
		got := goose.eats(bean, s)
		expected := true
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Goose doesn't eat bean when farmer present", func(t *testing.T) {
		s := state{left, left, left, right}
		got := goose.eats(bean, s)
		expected := false
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})
}

func TestPlaceString(t *testing.T) {
	t.Run("Left", func(t *testing.T) {
		c := left
		got := fmt.Sprintf("%s", c)
		expected := "Left"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Boat", func(t *testing.T) {
		c := boat
		got := fmt.Sprintf("%s", c)
		expected := "Boat"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Right", func(t *testing.T) {
		c := right
		got := fmt.Sprintf("%s", c)
		expected := "Right"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestStateDestination(t *testing.T) {
	t.Run("Left leads to boat", func(t *testing.T) {
		s := state{left, left, left, left}
		got := s.destinations()
		expected := []place{boat}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Right leads to boat", func(t *testing.T) {
		s := state{right, left, left, left}
		got := s.destinations()
		expected := []place{boat}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Boat leads to left and right", func(t *testing.T) {
		s := state{boat, left, left, left}
		got := s.destinations()
		expected := []place{left, right}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestStateTravellers(t *testing.T) {
	t.Run("Travellers are farmer and everyone with farmer", func(t *testing.T) {
		s := state{boat, left, boat, left}
		got := s.travellers()
		expected := []char{farmer, bean}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestStateGenerateMoves(t *testing.T) {
	t.Run("First move: Farmer to boat with and without companion", func(t *testing.T) {
		s := state{left, left, left, left}
		got := s.generateMoves()
		expected := []state{
			{boat, left, left, left},
			{boat, boat, left, left},
			{boat, left, boat, left},
			{boat, left, left, boat},
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Second move: Farmer to left or right with with and without goose", func(t *testing.T) {
		s := state{boat, boat, left, left}
		got := s.generateMoves()
		expected := []state{
			{left, boat, left, left},
			{left, left, left, left},
			{right, boat, left, left},
			{right, right, left, left},
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}
