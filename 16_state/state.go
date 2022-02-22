package state

import "fmt"

type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type leaderApproveState struct{}
func GetLeaderApproveState() IState {
	return leaderApproveState{}
}
func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}
func (leaderApproveState) Reject(m *Machine) {
	fmt.Println("leader 审批失败")
}
func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

type financeApproveState struct{}
func GetFinanceApproveState() IState {
	return financeApproveState{}
}
func (financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}
func (financeApproveState) Reject(m *Machine) {
	fmt.Println("财务审批失败")
	m.SetState(GetLeaderApproveState())
}
func (financeApproveState) GetName() string {
	return "FinanceApproveState"
}