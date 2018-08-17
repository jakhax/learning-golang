package main

//@todo: implement a database
// as you guessed it was hell writing this code

type Person struct {
	Name         string
	Job          string
	Summary      string
	PersonalInfo []*PersonalInfo
	Skills       []string
	Experience   []string
	Education    []string
	Posts        []*Employment
}
type PersonalInfo struct {
	Title string
	Info  string
}

type Employment struct {
	Company string
	Period  string
	Title   string
	Duties  []string
}

func getData() Person {
	infoMap := map[string]string{
		"Address":       "4454 Silicon savannah, Kenya",
		"Phone":         "+24700001234",
		"Email":         "jackogina@coder.com",
		"Website":       "jack.solutions.com",
		"Stackoverflow": "stackoverflow/users/12345/jakogina",
		"Github":        "github.com/jakhax",
		"Linked":        "linked.com/users/jakogina",
	}
	pInfo := []*PersonalInfo{}
	for title, info := range infoMap {
		p := PersonalInfo{Title: title, Info: info}
		pInfo = append(pInfo, &p)
	}
	edu := []string{"Kapsabet High School, K.C.S.E",
		"Moringa School, Full Stack Developer",
		"Some University, Mathematics & Computer Science"}

	experince := []string{"Engineering web development, all layers, from database to services to user interfaces",
		"Supporting legacy systems with backups of all cases to/from parallel systems",
		"Analysis and design of databases and user interfaces",
		"Managing requirements",
		"Implementing software development lifecycle policies and procedures",
		"Managing and supporting multiple projects",
		"Highly adaptable in quickly changing technical environments with very strong organizational and analytical skills"}

	summary := "A results-driven, customer-focused, articulate and analytical Senior Software Engineer who can think “out of the box.” Strong in design and integration problem-solving skills. Expert in Java, C#, .NET, and T-SQL with database analysis and design."

	skills := []string{"Databases: MySQL, Oracle, Access, SAP",
		"Software: Microsoft Office, Remedy, Microsoft SQL Server, DB Artisan, Eclipse, Visual Studio.NET, FrontPage",
		"Languages: C#, Java, Visual Basic, ASP, XML, XSL, JWS, SQL, and Python"}

	e1 := Employment{Company: "E*Trade Financial, Silicon Valley.",
		Period: "CA July 2012 – Present",
		Title:  "Software Engineer (Customer Service Systems)",

		Duties: []string{"Re-engineered customer account software systems used by brokerage teams.",
			"Web developer for user interfaces to trading inquiries, support parallel systems.",
			"Developed and implemented new feedback system for users concerns, bugs, and defect tracking regarding use and functionality of new interfaces.",
			"Coded web designed interfaces using Java, XML, XSL, AJAX, and JWS.",
			"Support system for existing intranet for employees, including designing and developing the Advantage@Work system company-wide.",
			"Code and support provided through ASP.NET, T-SQL, Microsoft SQL Server, and Oracle 9i.",
			"Supported existing legacy system to provide newly created cases and ensured they were available in the systems in parallel until legacy systems were retired.",
		}}

	e2 := Employment{Company: "Intel Corporation, Silicon savannah",
		Period: "Jan 2005 – Jul 2012",
		Title:  "Systems Programmer (Remote Servers and SSL Product Analyst)",

		Duties: []string{"Deployed and tested Remote Installation Services(RIS)-Server Installs on Windows XP.",
			"Focused deployment of Server builds and handled some client builds.",
			"Modified Visual Basic applications for use in post-server builds for customizing builds.",
			"Researched RIS and Active Directory for future deployment worldwide.",
			"Wrote bi-monthly progress reports, participated in weekly staff meetings and JDP team meetings designed to develop white paper processing.",
			"Provide technical support to the SSL team, managing inventory."}}

	employment := []*Employment{&e1, &e2}

	user := Person{
		Name:         "JakHax",
		Job:          "Software Developer",
		Summary:      summary,
		PersonalInfo: pInfo,
		Skills:       skills,
		Posts:        employment,
		Experience:   experince,
		Education:    edu,
	}
	return user
}
