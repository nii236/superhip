import React from 'react';
import { Admin, Resource, List, Datagrid, TextField } from 'react-admin';
import data from './data';
import auth from './auth';
import { List as IList } from "immutable"
import FontAwesomeIcon from '@fortawesome/react-fontawesome'
import faCoffee from '@fortawesome/fontawesome-free-solid/faCoffee'
import fontawesome from '@fortawesome/fontawesome'
fontawesome.library.add(faCoffee)
const ListFactory = (fields: IList<string>) => {
  return (props) => {
    return <ResourceList fields={fields} {...props}/>
  }
}

const ResourceList = (props) => {
  const { fields } = props
  console.log(props)
  return (
    <List {...props}>
      <Datagrid>      
        {fields.map(field => {
          return <TextField key={field} source={field} />
        })}
      </Datagrid>
    </List>
  )
}

const dataProvider = data.Provider("http://localhost:3001")
const App = (props) => {
  const UsersList = ListFactory(IList([
    "id",
    "first_name",
    "last_name",
    "email",
  ]))

  const SitesList = ListFactory(IList([
    "id",
    "name",
  ]))

  const ModelsList = 
  ListFactory(IList([
    "id",
    "name",
  ]))

  const ModelGroupsList = ListFactory(IList([
    "id",
    "name",
  ]))

  const SnapshotsList = ListFactory(IList([
    "id",
    "name",
  ]))
  return (
    <Admin dataProvider={dataProvider} authProvider={auth.Provider}>    
      <Resource icon={() => ( <div>heyo</div>)} name="users" list={(props) => <UsersList {...props}/>} />
      <Resource icon={() => (<FontAwesomeIcon icon={"coffee"} />)} name="sites" list={(props) => <SitesList {...props}/>} />
      <Resource name="models" list={(props) => <ModelsList {...props}/>} />
      <Resource name="modelgroups" list={(props) => <ModelGroupsList {...props}/>} />
      <Resource name="snapshots" list={(props) => <SnapshotsList {...props}/>} />
    </Admin>)
}

export default App;

