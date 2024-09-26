package taskmanagerAPI

func SortTasks(list Tasks) Tasks {
	sorted := false
	for !sorted {
		sorted = true

		for i, elem := range list.All {
			if i < len(list.All)-1 && elem.ID > list.All[i+1].ID {
				list.All[i+1], list.All[i] = list.All[i], list.All[i+1]
				sorted = false
			}
		}
	}

	return list
}