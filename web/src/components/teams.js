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
	ShowButton,
	Resource,
	Edit as AdminEdit,
	List as AdminList,
	Create as AdminCreate,
	Show as AdminShow,
	Datagrid,
	TextField,
	SimpleShowLayout,
	ReferenceArrayInput,
	SelectArrayInput,
	ChipField,
	ReferenceArrayField,
	SingleFieldList
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
				<ReferenceArrayField label="Users" reference="users" source="user_ids">
					<SingleFieldList linkType="show">
						<ChipField source="first_name" />
					</SingleFieldList>
				</ReferenceArrayField>
				<ReferenceArrayField label="Students" reference="students" source="student_ids">
					<SingleFieldList linkType="show">
						<ChipField source="name" />
					</SingleFieldList>
				</ReferenceArrayField>
				<EditButton />
				<ShowButton />
			</Datagrid>
		</AdminList>
	)
}
export const Show = props => (
	<AdminShow title={<Title />} {...props}>
		<SimpleShowLayout>
			<TextField source="name" />
			<ReferenceField label="School" source="school_id" reference="schools">
				<TextField source="name" />
			</ReferenceField>
			<ReferenceArrayField label="Users" reference="users" source="user_ids">
				<SingleFieldList linkType="show">
					<ChipField source="first_name" />
				</SingleFieldList>
			</ReferenceArrayField>
			<ReferenceArrayField label="Students" reference="students" source="student_ids">
				<SingleFieldList linkType="show">
					<ChipField source="name" />
				</SingleFieldList>
			</ReferenceArrayField>
		</SimpleShowLayout>
	</AdminShow>
)

export const Title = ({ record }) => {
	return <span>Team {record.name ? `"${record.name}"` : ""}</span>
}

export const Edit = props => {
	return (
		<AdminEdit title={<Title />} {...props}>
			<SimpleForm>
				<ReferenceArrayInput label="User" source="user_ids" reference="users">
					<SelectArrayInput optionText="first_name">
						<ChipField source="first_name" />
					</SelectArrayInput>
				</ReferenceArrayInput>
				<ReferenceArrayInput label="Student" source="student_ids" reference="students">
					<SelectArrayInput optionText="name">
						<ChipField source="name" />
					</SelectArrayInput>
				</ReferenceArrayInput>
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
