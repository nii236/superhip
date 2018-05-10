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
	Datagrid,
	TextField
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
				<ReferenceField label="School" source="school_id" reference="schools">
					<TextField source="name" />
				</ReferenceField>
				<EditButton />
			</Datagrid>
		</AdminList>
	)
}

export const Title = ({ record }) => {
	return <span>Student {record.name ? `"${record.name}"` : ""}</span>
}

export const Edit = props => {
	return (
		<AdminEdit title={<Title />} {...props}>
			<SimpleForm>
				<TextInput source="name" />
				<ReferenceInput label="School" source="school_id" reference="schools">
					<SelectInput optionText="name" />
				</ReferenceInput>
			</SimpleForm>
		</AdminEdit>
	)
}

export const Create = props => (
	<AdminCreate title={<Title />} {...props}>
		<SimpleForm>
			<TextInput source="name" />
			<ReferenceInput label="School" source="school_id" reference="schools">
				<SelectInput optionText="name" />
			</ReferenceInput>
		</SimpleForm>
	</AdminCreate>
)
