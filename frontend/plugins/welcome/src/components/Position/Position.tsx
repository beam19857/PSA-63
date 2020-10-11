/* eslint-disable no-use-before-define */
import React from 'react';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';

export default function ComboBox() {
  return (
    <Autocomplete
      id="combo-box-position"
      options={Position}
      getOptionLabel={(option) => option.title}
      style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:8}}
      renderInput={(params) => <TextField {...params} label="Position" variant="outlined" />}
    />
  );
}

// Top 100 films as rated by IMDb users. http://www.imdb.com/chart/top
const Position = [
  { title: 'Head'},
  { title: 'Duty Head'},
];
