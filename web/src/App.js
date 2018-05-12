import React from "react"
import * as roles from "./components/roles"
import * as permissions from "./components/permissions"
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
			<Resource show={roles.Show} edit={roles.Edit} create={roles.Create} name="roles" list={roles.List} />
			<Resource show={permissions.Show} edit={permissions.Edit} create={permissions.Create} name="permissions" list={permissions.List} />
			<Resource show={schools.Show} edit={schools.Edit} create={schools.Create} name="schools" list={schools.List} />
			<Resource show={users.Show} edit={users.Edit} create={users.Create} name="users" list={users.List} />
			<Resource show={teams.Show} edit={teams.Edit} create={teams.Create} name="teams" list={teams.List} />
			<Resource show={students.Show} edit={students.Edit} create={students.Create} name="students" list={students.List} />
		</Admin>
	)
}

export default App
