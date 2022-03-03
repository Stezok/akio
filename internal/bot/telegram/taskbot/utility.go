package taskbot

func (bot *TaskBot) isTeamMember(id int64) bool {
	for _, _id := range bot.Team {
		if _id == id {
			return true
		}
	}
	return false
}
