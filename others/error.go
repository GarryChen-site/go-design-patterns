package others

import (
	"encoding/binary"
	"io"
)

type Point struct {
	Longitude, Latitude, Distance, ElevationGain, ElevationLoss int
}

func parse(r io.Reader) (*Point, error) {

	var p Point
    
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
	    return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
	    return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
	    return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
	    return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
	    return nil, err
	}
	return &p, nil
}

func parse2(r io.Reader) (*Point, error) {
	var p Point
	var err	error

	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)    
	read(&p.Distance)    
	read(&p.ElevationGain)    
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

