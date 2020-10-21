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

import { EntDepartment } from '../../api/models/EntDepartment'; //import interface Department
import { EntExpertise } from '../../api/models/EntExpertise'; //import interface Expertise
import { EntPosition } from '../../api/models/EntPosition'; //import interface Postiion
import { EntUser } from '../../api/models/EntUser'; //import interface User


// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

 
export default function UserInformation() {
 const classes = useStyles();
 const profile = { givenName: 'Hospital Register' };
 const http = new DefaultApi();

 const [users, setUser] = useState<EntUser[]>([]);

 const [departments, setDepartments] = useState<EntDepartment[]>([])
 const [expertises, setExpertises] = useState<EntExpertise[]>([]);
 const [positions, setPositions] = useState<EntPosition[]>([]);


 const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);

 const [doctorname, setDoctorName] = useState(String);
 const [doctoremail, setDoctorEmail] = useState(String);
 const [department, setDepartment] = useState(Number);
 const [expertise, setExpertise] = useState(Number);
 const [position, setPosition] = useState(Number);


 useEffect(() => {
  const getDepartments = async () => {
    const res = await http.listDepartment({ limit: 10, offset: 0 });
    setLoading(false);
    setDepartments(res);
  };
  getDepartments();

  const getExpertises = async () => {
    const res = await http.listExpertise({ limit: 10, offset: 0 });
    setLoading(false);
    setExpertises(res);
    console.log(res);
  };
  getExpertises();

  const getPositions = async () => {
    const res = await http.listPosition({ limit: 10, offset: 0 });
    setLoading(false);
    setPositions(res);
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

  const res: any = await http.createUser({ user:user });
  setStatus(true);
  if (res.id != '') {
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
      title={`User information record`}
      type=".. systems"> 
       <Button variant="contained" color="default" href="/table"> User Record
           </Button>
    </Header>

    <Content>
      <ContentHeader title="User information"> 
            <Button onClick={() => {CreateUser();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new user </Button>
            <Button style={{ marginLeft: 20 }} component={RouterLink} to="/" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>

            {status ? ( 
                    <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
            {alert ? ( 
                    <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
            : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
          ) : null}

      </ContentHeader>
      <div className={classes.root}>
        <form noValidate autoComplete="off">
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
      
              id="doctorEmail"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={doctoremail}
              onChange={handleDoctorEmailChange}
            />

            <div className={classes.paper}><strong>Department</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="department-label"
              id="department"
              value={department}
              onChange={DepartmenthandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">choose department</InputLabel>

              {departments.map((item: EntDepartment) => (
                <MenuItem value={item.id}>{item.departmentName}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>Expertise</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="expertise"
              value={expertise}
              onChange={ExpertisehandleChange}
            >
              <InputLabel className={classes.insideLabel}>choose expertise</InputLabel>

              {expertises.map((item: EntExpertise) => (
                <MenuItem value={item.id}>{item.expertiseName}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>Position</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="position"
              value={position}
              onChange={PositionhandleChange}
            >
              <InputLabel className={classes.insideLabel}>choose position</InputLabel>

              {positions.map((item: EntPosition) => (
                <MenuItem value={item.id}>{item.positionName}</MenuItem>
              ))}
            </Select>
         
          
          </FormControl>

        </form>
      </div>
    </Content>
  </Page>
);
}