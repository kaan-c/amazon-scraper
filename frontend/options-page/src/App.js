import logo from './logo.svg';
import './App.css';
import React from 'react';

class pageComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {address: ""}; 
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this); 
  }
  handleSubmit(event) {
    console.log("address is " + this.state.address);
    fetch("http://localhost:10000/address/" + this.state.address)
      .then((response) => response.json())
      .then((data) => console.log('This is your data', data));
    event.preventDefault();
  }

  handleChange(event) {
    this.setState({address: event.target.value});  
  }

  render() {
    return (<form onSubmit={this.handleSubmit}>
          <label>Address:</label>
          <input type="text" id="url" value={this.state.address} onChange={this.handleChange} />
          <button type="submit">Submit</button>
    </form>);
  }
}

export default pageComponent;
