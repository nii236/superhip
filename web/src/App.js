import React from "react"
import * as schools from "./components/schools"
import * as users from "./components/users"
import * as teams from "./components/teams"
import * as students from "./components/students"
import FontAwesomeIcon from "@fortawesome/react-fontawesome"
import faCoffee from "@fortawesome/fontawesome-free-solid/faCoffee"
import fontawesome from "@fortawesome/fontawesome"
import data from "./data"
import auth from "./auth"
import {
	Admin,
	Resource,
} from "react-admin"

fontawesome.library.add(faCoffee)


const dataProvider = data.Provider("http://localhost:8080")
const App = props => {
	return (
		<Admin dataProvider={dataProvider} authProvider={auth.Provider}>
			<Resource edit={schools.Edit} create={schools.Create} name="schools" list={schools.List} />
			<Resource edit={users.Edit} create={users.Create} name="users" list={users.List} />
			<Resource edit={teams.Edit} create={teams.Create} name="teams" list={teams.List} />
			<Resource edit={students.Edit} create={students.Create} name="students" list={students.List} />
		</Admin>
	)
}

export default App
