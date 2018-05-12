import React from "react"
import {
	ReferenceField,
	SimpleForm,
	DisabledInput,
	ReferenceInput,
	SelectInput,
	TextInput,
	LongTextInput,
	EditButton,
	Resource,
	Edit as AdminEdit,
	List as AdminList,
	Create as AdminCreate,
	Show as AdminShow,
	Datagrid,
	TextField,
	SimpleShowLayout
} from "react-admin"
import { List as IList } from "immutable"
import FontAwesomeIcon from "@fortawesome/react-fontawesome"
import faCoffee from "@fortawesome/fontawesome-free-solid/faCoffee"
import fontawesome from "@fortawesome/fontawesome"

export const List = props => {
	return (
		<AdminList {...props}>
			<Datagrid>
				<TextField source="name" />
				<EditButton />
			</Datagrid>
		</AdminList>
	)
}
export const Show = props => (
	<AdminShow title={<Title />} {...props}>
		<SimpleShowLayout>
			<TextField source="name" />
		</SimpleShowLayout>
	</AdminShow>
)
export const Title = ({ record }) => {
	return <span>Role {record.name ? `"${record.name}"` : ""}</span>
}

export const Edit = props => {
	return (
		<AdminEdit title={<Title />} {...props}>
			<SimpleForm>
				<TextInput source="name" />
			</SimpleForm>
		</AdminEdit>
	)
}

export const Create = props => (
	<AdminCreate title={<Title />} {...props}>
		<SimpleForm>
			<TextInput source="name" />
		</SimpleForm>
	</AdminCreate>
)
