package api

import "fmt"

type App struct {
	OwnerName string
	Id        string
	Platform  string
}

func (a App) verify() error {
	if a.OwnerName == "" {
		return fmt.Errorf("app's owner name must be present")
	}
	if a.Id == "" {
		return fmt.Errorf("application id must be present")
	}
	if a.Platform == "" {
		return fmt.Errorf("platform must be present")
	}

	return nil
}

type UserApp struct {
	OwnerName string
	Id        string
	Platform  string
}

func (a UserApp) verify() error {
	if a.OwnerName == "" {
		return fmt.Errorf("app's owner name must be present")
	}
	if a.Id == "" {
		return fmt.Errorf("application id must be present")
	}
	if a.Platform == "" {
		return fmt.Errorf("platform must be present")
	}

	return nil
}

type OrganizationApp struct {
	OwnerName string
	Id        string
	Platform  string
}

func (a OrganizationApp) verify() error {
	if a.OwnerName == "" {
		return fmt.Errorf("app's owner name must be present")
	}
	if a.Id == "" {
		return fmt.Errorf("application id must be present")
	}
	if a.Platform == "" {
		return fmt.Errorf("platform must be present")
	}

	return nil
}

type Enterprise struct {
	Name string
}

func (a Enterprise) verify() error {
	if a.Name == "" {
		return fmt.Errorf("enterprise name must be present")
	}

	return nil
}
