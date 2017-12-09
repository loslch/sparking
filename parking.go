package main

import (
	"io"
	"bufio"
	"os/exec"
)

type Location struct {
	Area    string `json:"area,omitempty"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Time    string `json:"time"`
}

type ParkingSpace struct {
	Name      string     `json:"name"`
	Locations []Location `json:"locations"`
}

func parsePage(pagePath string) (io.ReadCloser, error) {
	cmd := exec.Command("./parse.sh", pagePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return stdout, nil
}

func generateParkingSpace(pageFilePath string) ([]ParkingSpace, error) {
	var parkingSpaces []ParkingSpace

	// Transform Parking Page into Text
	stdout, err := parsePage(pageFilePath)
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	scanner := bufio.NewScanner(stdout)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Parse and Generate Parking Space Information
	for scanner.Scan(); scanner.Text() != ""; {
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "=====" {
				break
			}

			lines = append(lines, line)
		}

		var space ParkingSpace

		if len(lines) == 4 {
			space.Name = lines[0]
			space.Locations = []Location{
				{
					Address: lines[1],
					Phone:   lines[2],
					Time:    lines[3],
				},
			}
		} else {
			space.Name = lines[0]
			space.Locations = []Location{
				{
					Area:    lines[1],
					Address: lines[2],
					Phone:   lines[3],
					Time:    lines[4],
				},
			}

			remains := lines[5:]
			switch len(remains) {
			case 0:
				//천안아산,동대구,신경주
				break
			case 1:
				//대전역,울산역
				space.Locations = append(space.Locations, Location{
					Area:    remains[0],
					Address: lines[2],
					Phone:   lines[3],
					Time:    lines[4],
				})
				break
			case 3:
				//광명역
				space.Locations = append(space.Locations, Location{
					Area:    remains[0],
					Address: remains[1],
					Phone:   remains[2],
					Time:    lines[4],
				})
				break
			case 4:
				//서울역
				space.Locations = append(space.Locations, Location{
					Area:    remains[0],
					Address: remains[1],
					Phone:   remains[2],
					Time:    remains[3],
				})
				break
			case 6:
				//부산역
				space.Locations = append(space.Locations, Location{
					Area:    remains[0],
					Address: remains[1],
					Phone:   remains[2],
					Time:    lines[4],
				})
				space.Locations = append(space.Locations, Location{
					Area:    remains[3],
					Address: remains[4],
					Phone:   remains[5],
					Time:    lines[4],
				})
				break
			}
		}

		parkingSpaces = append(parkingSpaces, space)
	}

	return parkingSpaces, nil
}
