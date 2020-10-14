import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Select,
  MenuItem,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import InputLabel from '@material-ui/core/InputLabel';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';

import { DefaultApi } from '../../api/apis';
import ComponanceDepartment from '../Department';
import ComponancePosition from '../Position';
import ComponanceExpertise from '../Expertise';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';
import { EntDepartment } from '../../api/models/EntDepartment'; // import interface Department
import { EntExpertise } from '../../api/models/EntExpertise'; // import interface Expertise
import { EntPosition } from '../../api/models/EntPosition'; // import interface Position
import { EntUser } from '../../api/models/EntUser'; // import interface User
//import { EntExpertiseEdges } from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(2),
     justifyContent: 'center',
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '500',
     marginLeft:7,
     marginRight:-7,
   },
   paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  insideLabel: {
    margin: theme.spacing(1),
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
    //marginTop:10,
  },
   formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
 }),
);
 
interface recordUser {
 doctorName: string;
 doctorEmail: string;
 department: number;
 position: number;
 expertise: number;
};
 
export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'Hospital Register' };
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 const [users, setUser] = React.useState<EntUser[]>([]);
 const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
 const [positions, setPositions] = React.useState<EntPosition[]>([]);
 const [expertises,setExpertises] = React.useState<EntExpertise[]>([]);

 const [doctorname, setDoctorName] = useState(String);
 const [doctoremail, setDoctorEmail] = useState(String);
 const [department, setDepartment] = useState(Number);
 const [expertise, setExpertise] = useState(Number);
 const [position, setPosition] = useState(Number);

  useEffect(() => {
  const getDepartments = async () => {
    const res = await api.listDepartment({ limit: 10, offset: 0 });
    setLoading(false);
    setDepartments(res);
    console.log(res);
  };
  getDepartments();

  const getExpertises = async () => {
    const res = await api.listExpertise({ limit: 10, offset: 0 });
    setLoading(false);
    setExpertises(res);
    console.log(res);
  };
  getExpertises();
  const getPositions = async () => {
    const res = await api.listPosition({ limit: 10, offset: 0 });
    setLoading(false);
    setPositions(res);
    console.log(res);
  };
  getPositions();
}, [loading]);
 
const handleDoctorNameChange = (event: any) => {
  setDoctorName(event.target.value as string);
};

const handleDoctorEmailChange = (event: any) => {
  setDoctorEmail(event.target.value as string);
};
const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setDepartment(event.target.value as number);
};
const ExpertisehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setExpertise(event.target.value as number);
};
const PositionhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPosition(event.target.value as number);
};


 const CreateUser = async () => {
  const user = {
    doctorName: doctorname,
    doctorEmail: doctoremail,
    department: department,
    position: position,
    expertise: expertise,
  };
  console.log(user);
   const res:any = await api.createUser({ user:user });
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
   <Page theme={pageTheme.tool}>
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
              //fullWidth
              //className={classes.margin}
              variant="outlined"
             
            >
               <div className={classes.paper}><strong>Name</strong></div>
              <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="doctorName"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={doctorname}
                onChange={handleDoctorNameChange}
              />

            <div className={classes.paper}><strong>Email</strong></div>
              <TextField className={classes.textField}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="doctorEmail"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={doctoremail}
                onChange={handleDoctorEmailChange}
              />

              <div className={classes.paper}><strong>แผนก</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                labelId="department-label"
                id="department"
                value={department}
                onChange={DepartmenthandleChange}
              >
                <InputLabel className={classes.insideLabel} id="department-label">เลือกแผนก(Department)</InputLabel>

                {departments.map((item: EntDepartment) => (
                  <MenuItem value={item.id}>{item.departmentName}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ความชำนาญการ</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="expertise"
                value={expertise}
                onChange={ExpertisehandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกความชำนาญการ(Expertise)</InputLabel>

                {expertises.map((item: EntExpertise) => (
                  <MenuItem value={item.id}>{item.expertiseName}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ตำแหน่ง</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="position"
                value={position}
                onChange={PositionhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกตำแหน่ง(Position)</InputLabel>

                {positions.map((item: EntPosition) => (
                  <MenuItem value={item.id}>{item.positionName}</MenuItem>
                ))}
              </Select>
           
            
            </FormControl>
    
               
             
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
