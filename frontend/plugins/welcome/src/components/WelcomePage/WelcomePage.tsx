import React, { FC } from 'react';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Icon from '@material-ui/core/Icon';
import CreateIcon from '@material-ui/icons/Create';
import MenuIcon from '@material-ui/icons/Menu';
import HomeIcon from '@material-ui/icons/Home';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import PersonAddIcon from '@material-ui/icons/PersonAdd';

import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,

} from '@backstage/core';
import { Container, TextField } from '@material-ui/core';

const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'to Computer Care' };
 
 
 
 return (
   <Page theme={pageTheme.home}> 
     <Header 
       title={`Welcome ${profile.givenName || 'to Backstage'}`
      
      }
       subtitle="">   
           <Button
                variant="contained"
                color="default"
                startIcon={<ExitToAppRoundedIcon />}
                href="#"
              >Logout
            </Button>
       
         </Header>
     <Content>
       <ContentHeader title="User Information ">
       </ContentHeader>
       <form>
         <Link component={RouterLink} to="/user" color="secondary">
           <Button
                variant="outlined"
                color="secondary"
                startIcon={<PersonAddIcon />}
              >Register
            </Button>
         </Link>
       </form>
         
     </Content>
   </Page>
 );
};
 
export default WelcomePage;

