import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

ReactDOM.render(
  <div>
    <h3>Country</h3>
    <button>Capital</button>
    <table>
      <tr>
        <th>id</th>
        <th>country</th>
        <th>insert_time</th>
      </tr>
      {/* <tr>
        <td>db</td>
        <td>db</td>
        <td>db</td>
      </tr> */}

    </table>
  </div>,
  document.getElementById('root')
);