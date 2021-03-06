package systemd

type SystemdService struct {
	manager *SystemdManager
	name    string
}

func NewSystemdService(manager *SystemdManager, name string) *SystemdService {
	return &SystemdService{manager, name}
}

func (ss *SystemdService) Name() string {
	return ss.name
}

func (ss *SystemdService) State() (string, string, string, []string, error) {
	loadState, activeState, subState, err := ss.manager.getUnitStates(ss.name)
	if err != nil {
		return "", "", "", nil, err
	}

	return loadState, activeState, subState, make([]string, 0), nil
}

func (ss *SystemdService) Payload() (string, error) {
	return ss.manager.readUnit(ss.Name())
}
