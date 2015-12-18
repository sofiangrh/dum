func usage() {
	command := os.Args[0]
	var subcommands []string
	for subcommand := range commands.CommandMap {
		subcommands = append(subcommands, subcommand)
	}
	sort.Strings(subcommands)
	fmt.Printf(usageMessageTemplate, command, strings.Join(subcommands, "\n  "), command)
}

func help() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	subcommand, ok := commands.CommandMap[os.Args[2]]
	if !ok {
		fmt.Printf("Unknown command %q\n", os.Args[2])
		usage()
		return
	}
	subcommand.Usage(os.Args[0])
}