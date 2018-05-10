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
    Datagrid,
    TextField
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
                <ReferenceField label="School" source="school_id" reference="schools">
                    <TextField source="name" />
                </ReferenceField>
                <EditButton />
            </Datagrid>
        </AdminList>
    )
}

export const Title = ({ record }) => {
    return <span>User {record.first_name ? `"${record.first_name}"` : ""}</span>
}

export const Edit = props => {
    return (
        <AdminEdit title={<Title />} {...props}>
            <SimpleForm>
                <TextInput source="first_name" />
                <TextInput source="last_name" />
                <TextInput source="email" />
                <TextInput source="role" />
                <TextInput source="password" />
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
            <TextInput source="first_name" />
            <TextInput source="last_name" />
            <TextInput source="email" />
            <TextInput source="role" />
            <TextInput source="password" />
            <ReferenceInput label="School" source="school_id" reference="schools">
                <SelectInput optionText="name" />
            </ReferenceInput>
        </SimpleForm>
    </AdminCreate>
)
