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
	SingleFieldList,
	Filter
} from "react-admin"

import { List as IList } from "immutable"
import FontAwesomeIcon from "@fortawesome/react-fontawesome"
import faCoffee from "@fortawesome/fontawesome-free-solid/faCoffee"
import fontawesome from "@fortawesome/fontawesome"

const ListFilter = props => {
	return (
		<Filter {...props}>
			<TextInput label="Search" source="name" alwaysOn />
		</Filter>
	)
}

export const List = props => {
	return (
		<AdminList {...props} filters={<ListFilter />}>
			<Datagrid>
				<TextField source="name" />
				<ReferenceField label="School" source="school_id" reference="schools">
					<TextField source="name" />
				</ReferenceField>
				<ReferenceArrayField label="Teams" reference="teams" source="team_ids">
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
			<ReferenceArrayField label="Teams" reference="teams" source="team_ids">
				<SingleFieldList linkType="show">
					<ChipField source="name" />
				</SingleFieldList>
			</ReferenceArrayField>
		</SimpleShowLayout>
	</AdminShow>
)
export const Title = ({ record }) => {
	return <span>Student {record.name ? `"${record.name}"` : ""}</span>
}

export const Edit = props => {
	return (
		<AdminEdit title={<Title />} {...props}>
			<SimpleForm>
				<ReferenceArrayInput label="Teams" source="team_ids" reference="teams">
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
			<ReferenceArrayInput label="Teams" source="team_ids" reference="teams">
				<SelectArrayInput optionText="name">
					<ChipField source="name" />
				</SelectArrayInput>
			</ReferenceArrayInput>
			<TextInput source="name" />
			<ReferenceInput label="School" source="school_id" reference="schools">
				<SelectInput optionText="name" />
			</ReferenceInput>
		</SimpleForm>
	</AdminCreate>
)
