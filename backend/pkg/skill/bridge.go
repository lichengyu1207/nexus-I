package skill

import (
	"fmt"
)

// Bridge 定义技能桥接接口
type Bridge interface {
	Execute(name string, input map[string]interface{}) (map[string]interface{}, error)
	List() ([]string, error)
}

// SkillManager 技能管理器
type SkillManager struct {
	skills map[string]Bridge
}

func NewSkillManager() *SkillManager {
	return &SkillManager{
		skills: make(map[string]Bridge),
	}
}

func (m *SkillManager) Register(name string, bridge Bridge) {
	m.skills[name] = bridge
}

func (m *SkillManager) Execute(name string, input map[string]interface{}) (map[string]interface{}, error) {
	if skill, ok := m.skills[name]; ok {
		return skill.Execute(name, input)
	}
	return nil, fmt.Errorf("skill not found: %s", name)
}

func (m *SkillManager) List() []string {
	names := make([]string, 0, len(m.skills))
	for name := range m.skills {
		names = append(names, name)
	}
	return names
}
