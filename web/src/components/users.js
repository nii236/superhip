import React from "react"
import {
	ReferenceField,
	Edit as AdminEdit,
	SimpleForm,
	DisabledInput,
	ReferenceInput,
	SelectInput,
	TextInput,
	LongTextInput,
	EditButton,
	Admin,
	Resource,
	List as AdminList,
	Create as AdminCreate,
	Show as AdminShow,
	Datagrid,
	TextField,
	ReferenceArrayInput,
	SelectArrayInput,
	ReferenceArrayField,
	SingleFieldList,
	ChipField,
	SimpleShowLayout,
	Field
} from "react-admin"

import { List as IList } from "immutable"


export const List = props => {
	return (
		<AdminList {...props}>
			<Datagrid>
				<TextField source="first_name" />
				<TextField source="last_name" />
				<TextField source="email" />
				<TextField source="role" />
				<ReferenceArrayField label="Roles" reference="roles" source="role_ids">
					<SingleFieldList linkType="show">
						<ChipField source="name" />
					</SingleFieldList>
				</ReferenceArrayField>
				<ReferenceArrayField label="Schools" reference="schools" source="school_ids">
					<SingleFieldList linkType="show">
						<ChipField source="name" />
					</SingleFieldList>
				</ReferenceArrayField>
				<EditButton />
			</Datagrid>
		</AdminList>
	)
}

export const Show = (props) => (
	<AdminShow title={<Title />} {...props}>
		<SimpleShowLayout>
			<TextField source="first_name" />
			<TextField source="last_name" />
			<ReferenceArrayField label="Roles" reference="roles" source="role_ids">
				<SingleFieldList linkType="show">
					<ChipField source="name" />
				</SingleFieldList>
			</ReferenceArrayField>
			<ReferenceArrayField label="Schools" reference="schools" source="school_ids">
				<SingleFieldList linkType="show" >
					<ChipField source="name" />
				</SingleFieldList>
			</ReferenceArrayField>
		</SimpleShowLayout>
	</AdminShow>
)


export const Title = ({ record }) => {
	return <span>User {record.first_name ? `"${record.first_name}"` : ""}</span>
}
const TestInput = (props) => {
	console.log(props.parent)
	// console.log(props.record)
	// console.log(props.input.value)
	return (<SelectArrayInput{...props} />)
}

const WithProps = ({ children, ...props }) => children(props);

export const Edit = props => {
	return (
		<AdminEdit title={<Title />} {...props}>
			{WithProps(
				<SimpleForm>
					<TextInput source="first_name" />
					<TextInput source="last_name" />
					<TextInput source="email" />
					<TextInput source="role" />
					<TextInput source="password" />
					<ReferenceArrayInput label="Role" source="role_ids" reference="roles">
						<TestInput parent={props} />
					</ReferenceArrayInput>
					<ReferenceArrayInput label="School" source="school_ids" reference="schools">
						<SelectArrayInput optionText="name" />
					</ReferenceArrayInput>

				</SimpleForm>)}

		</AdminEdit >
	)
}

export const Create = props => (
	<AdminCreate title={<Title />} {...props}>
		<SimpleForm>
			<TextInput source="first_name" />
			<TextInput source="last_name" />
			<TextInput source="email" />
			<TextInput source="role" />
			<TextInput source="password" />
			<ReferenceArrayInput label="Roles" reference="roles" source="role_ids">
				<SelectArrayInput optionText="name" />
			</ReferenceArrayInput>
			<ReferenceArrayInput label="School" source="school_ids" reference="schools">
				<SelectArrayInput optionText="name" />
			</ReferenceArrayInput>
		</SimpleForm>
	</AdminCreate>
)
