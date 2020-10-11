/* eslint-disable no-use-before-define */
import React from 'react';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';

export default function ComboBox() {
  return (
    <Autocomplete
      id="combo-box-demo"
      options={Department}
      getOptionLabel={(option) => option.title}
      style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:8}}
      renderInput={(params) => <TextField {...params} label="Department" variant="outlined" />}
    />
  );
}

// Top 100 films as rated by IMDb users. http://www.imdb.com/chart/top
const Department = [
  { title: 'GENERAL'},
  { title: 'MEDICINE'},
  { title: 'OPERATING ROOM '},
  { title: 'LABOR ROOM',},
  { title: 'EMERGENCY ROOM'},
];
