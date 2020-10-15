import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { EntUser } from '../../api/models/EntUser';
import { ControllersUser } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
export default function ComponentsRecordUserTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getUsers = async () => {
      const res = await http.listUser({ limit: 10, offset: 0 });
      setLoading(true);
      setUsers(res);
      console.log(res);
    };
    getUsers();
  }, [loading]);
  
  const deleteUsers = async (id: number) => {
    const res = await http.deleteUser({ id: id });
    setLoading(true);
  };
 
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`User information record`} type="Repairing computer systems" >
    <Button variant="contained" color="default" href="/user" startIcon={<PersonAddRoundedIcon />}> New user</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">ID</TableCell>
            <TableCell align="center">Name</TableCell>
            <TableCell align="center">Email</TableCell>
            <TableCell align="center">Department</TableCell>
            <TableCell align="center">Expertise</TableCell>
            <TableCell align="center">Position</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((item:any) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.doctorName}</TableCell>
              <TableCell align="center">{item.doctorEmail}</TableCell>
              <TableCell align="center">{item.edges?.userDepartment?.departmentName}</TableCell>
              <TableCell align="center">{item.edges?.userExpertise?.expertiseName}</TableCell>
              <TableCell align="center">{item.edges?.userPosition?.positionName}</TableCell>
              <TableCell align="center">
                <Button
                  onClick={() => {
                    deleteUsers(item.id);
                  }}
                  style={{ marginLeft: 10 }}
                  variant="contained"
                  color="secondary"
                  href="/table"
                >
                  Delete
                </Button>
              </TableCell>
              
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
   </Content>
  </Page>
);
}