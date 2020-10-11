import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import ComponanceDepartment from '../Department';
import ComponancePosition from '../Position';
import ComponanceExpertise from '../Expertise';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(1),
     justifyContent: 'center',
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
   formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
 }),
);
 
const initialUserState = {
 name: '',
 doctorid: '',
};
 
export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'Hospital Register' };
 const api = new DefaultApi();


 const [user, setUser] = useState(initialUserState);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 const handleInputChange = (event: any) => {
   const { id, value } = event.target;
   setUser({ ...user, [id]: value });
 };
 useEffect(() => {
  const getUsers = async () => {
    const res = await api.listUser({ limit: 10, offset: 0 });
    setLoading(false);
    setUsers(res);
  };
  getUsers();
}, [loading]);
 
 const CreateUser = async () => {
   const res = await api.createUser({ user });
   setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);

   
 };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
       subtitle="Suranaree Hospital "
     ></Header>
     <Content>
       <ContentHeader title="Register">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 Register Complete!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off" >

        

           <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField 
               id="name"
               label="Name"
               variant="outlined"
               type="string"
               size="medium"
               value={user.name}
               onChange={handleInputChange}
             />
             </FormControl>

            <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            >
              <TextField
               id="doctorid"
               label="DoctorID"
               variant="outlined"
               type="string"
               size="medium"
               value={user.doctorid}
               onChange={handleInputChange}
             />
             </FormControl>
          <ComponanceDepartment></ComponanceDepartment>
          <ComponanceExpertise></ComponanceExpertise>
          <ComponancePosition></ComponancePosition>
    
               
             
           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateUser();
               }}
               variant="contained"
               color="primary"
               startIcon={<CheckCircleIcon/>}
             >
                Create new user
             </Button>

             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
               color="secondary"
               startIcon={<CancelIcon/>}
             >
               Dismiss
             </Button>
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}
