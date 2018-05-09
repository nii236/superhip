import React from "react"
import {
	ReferenceField,
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

const UsersList = props => {
	return (
		<List {...props}>
			<Datagrid>
				<TextField source="first_name" />
				<TextField source="last_name" />
				<TextField source="email" />
				<TextField source="role" />
				<EditButton />
			</Datagrid>
		</List>
	)
}

const UserTitle = ({ record }) => {
	return <span>User {record.first_name ? `"${record.first_name}"` : ""}</span>
}

const UserEdit = props => {
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

const UserCreate = props => (
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

const TeamsList = props => {
	return (
		<List {...props}>
			<Datagrid>
				<TextField source="name" />
				<ReferenceField label="User" source="user_id" reference="users">
					<TextField source="first_name" />
				</ReferenceField>
				<EditButton />
			</Datagrid>
		</List>
	)
}

const TeamTitle = ({ record }) => {
	return <span>Team {record.name ? `"${record.name}"` : ""}</span>
}

const TeamEdit = props => {
	return (
		<Edit title={<TeamTitle />} {...props}>
			<SimpleForm>
				<TextInput source="name" />
				<ReferenceInput label="User" source="user_id" reference="users">
					<SelectInput optionText="first_name" />
				</ReferenceInput>
			</SimpleForm>
		</Edit>
	)
}

const TeamCreate = props => (
	<Create title={<TeamTitle />} {...props}>
		<SimpleForm>
			<TextInput source="name" />
			<ReferenceInput label="User" source="user_id" reference="users">
				<SelectInput optionText="first_name" />
			</ReferenceInput>
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
	return (
		<Admin dataProvider={dataProvider} authProvider={auth.Provider}>
			<Resource edit={UserEdit} create={UserCreate} name="users" list={UsersList} />
			<Resource edit={TeamEdit} create={TeamCreate} name="teams" list={TeamsList} />
		</Admin>
	)
}

export default App
