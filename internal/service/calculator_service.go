package service

import (
	pc "palettecalculator"
)

type PaletteClient interface {
	CalculatePredominantColor(file string) (*pc.Color, error)
	CalculateComplimentaryColorScheme(dc *pc.Color) []pc.Color
	CalculateSplitComplimentaryColorScheme(dc *pc.Color) []pc.Color
	CalculateTriadicColorScheme(dc *pc.Color) []pc.Color
	CalculateTetradicColorScheme(dc *pc.Color) []pc.Color
}

type Calculator interface {
	CalculatePredominant(file string) (*pc.Color, error)
	CalculateComplimentary(c *pc.Color) []pc.Color
	CalculateSplitComplimentary(c *pc.Color) []pc.Color
	CalculateTriadic(c *pc.Color) []pc.Color
	CalculateTetradic(c *pc.Color) []pc.Color
}

type calculatorService struct {
	pc PaletteClient
}

func NewCalculatorService() (*calculatorService, error) {
	calc, err := pc.NewPaletteCalculator()
	if err != nil {
		return nil, err
	}
	return &calculatorService{
		pc: calc,
	}, nil
}

func (cs *calculatorService) CalculatePredominant(file string) (*pc.Color, error) {
	color, err := cs.pc.CalculatePredominantColor(file)

	if err != nil {
		return nil, err
	}

	return color, nil
}

func (cs *calculatorService) CalculateComplimentary(c *pc.Color) []pc.Color {
	return cs.pc.CalculateComplimentaryColorScheme(c)
}

func (cs *calculatorService) CalculateSplitComplimentary(c *pc.Color) []pc.Color {
	return cs.pc.CalculateSplitComplimentaryColorScheme(c)
}

func (cs *calculatorService) CalculateTriadic(c *pc.Color) []pc.Color {
	return cs.pc.CalculateTriadicColorScheme(c)
}

func (cs *calculatorService) CalculateTetradic(c *pc.Color) []pc.Color {
	return cs.pc.CalculateTetradicColorScheme(c)
}
