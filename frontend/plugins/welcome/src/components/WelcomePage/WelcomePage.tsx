import React, { FC, useEffect,useState } from 'react';
import Button from '@material-ui/core/Button';
import AddBoxRoundedIcon from '@material-ui/icons/AddBoxRounded';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
 Content,
 Header,
 ItemCard,
 Page,
 pageTheme,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import { EntUser, EntUserFromJSON } from '../../api/models/EntUser'; //import interface User
import { DefaultApi } from '../../api/apis';

interface userInformation {
  doctorName: string;
  doctorEmail: string;  
  department: number; 
  position: number;
  expertise: number;
}


const WelcomePage: FC<{}> = () => {
const [doctoremail, setDoctorEmail] = useState(String);
const [users, setUsers] = useState<EntUser[]>([]);
const http = new DefaultApi();
const email = [];
const [loading, setLoading] = useState(true);

const handleDoctorEmailChange = (event: any) => {
  setDoctorEmail(event.target.value as string);
  const doctorEmail = event.target.doctorEmail as keyof typeof WelcomePage;
};

function login(){
  
  {users.map((item:EntUser) => (
  console.log(item.doctorEmail)

));
}
  };
  
  useEffect(() => {
    const getUsers = async () => {
      const res = await http.listUser({ limit: 10, offset: 0 });
      setLoading(true);
      setUsers(res);
      console.log(res);
    };
    getUsers();
  }, [loading]);

//console.log(email);
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="User Infomation" type="group 02">
     </Header>

    <Content align="center">
    <br></br><br></br><br></br><br></br><br></br><br></br><br></br>
      <TextField label="Email" variant="outlined" ></TextField>
      &nbsp;&nbsp;&nbsp;&nbsp;<br></br><br></br>
      <Button variant="contained" color="primary">Login</Button>
    <br></br>
    <br></br>
  
    <br></br>
     <Button variant="outlined" color="default" href="/user" startIcon={<AddBoxRoundedIcon />}> User information record
           </Button>
     </Content>



   </Page>
 );
};

export default WelcomePage;