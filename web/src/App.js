import React from "react"
import {
	Edit,
	SimpleForm,
	DisabledInput,
	ReferenceInput,
	SelectInput,
	TextInput,
	LongTextInput,
	EditButton,
	Admin,
	Resource,
	List,
	Create,
	Datagrid,
	TextField
} from "react-admin"
import data from "./data"
import auth from "./auth"
import { List as IList } from "immutable"
import FontAwesomeIcon from "@fortawesome/react-fontawesome"
import faCoffee from "@fortawesome/fontawesome-free-solid/faCoffee"
import fontawesome from "@fortawesome/fontawesome"
fontawesome.library.add(faCoffee)
const ListFactory = (fields: IList<string>) => {
	return props => {
		return <ResourceList fields={fields} {...props} />
	}
}

const UserTitle = ({ record }) => {
	return <span>User {record.first_name ? `"${record.first_name}"` : ""}</span>
}

export const UserEdit = props => {
	return (
		<Edit title={<UserTitle />} {...props}>
			<SimpleForm>
				<TextInput source="first_name" />
				<TextInput source="last_name" />
				<TextInput source="email" />
				<TextInput source="role" />
				<TextInput source="password" />
			</SimpleForm>
		</Edit>
	)
}

export const UserCreate = props => (
	<Create title={<UserTitle />} {...props}>
		<SimpleForm>
			<TextInput source="first_name" />
			<TextInput source="last_name" />
			<TextInput source="email" />
			<TextInput source="role" />
			<TextInput source="password" />
		</SimpleForm>
	</Create>
)

const ResourceList = props => {
	const { fields } = props
	return (
		<List {...props}>
			<Datagrid>
				{fields.map((field, i) => {
					return <TextField key={i} source={field} />
				})}
				<EditButton />
			</Datagrid>
		</List>
	)
}

const dataProvider = data.Provider("http://localhost:8080")
const App = props => {
	const UsersList = ListFactory(IList(["first_name", "last_name", "email", "role"]))

	const SitesList = ListFactory(IList(["id", "name"]))

	const ModelsList = ListFactory(IList(["id", "name"]))

	const ModelGroupsList = ListFactory(IList(["id", "name"]))

	const SnapshotsList = ListFactory(IList(["id", "name"]))
	return (
		<Admin dataProvider={dataProvider} authProvider={auth.Provider}>
			<Resource edit={UserEdit} create={UserCreate} name="users" list={props => <UsersList {...props} />} />
		</Admin>
	)
}

export default App
